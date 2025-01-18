package handlers

import (
	"context"
	"net/http"
	"sabu-user-service/dtos"
	"sabu-user-service/models"
	"sabu-user-service/proto/pb"
	"sabu-user-service/repository"
	"sabu-user-service/utils"
	"time"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserRepo       *repository.UserRepository
	RestaurantGRPC  pb.RestaurantServiceClient
	DonorGRPC  		pb.DonorServiceClient
}

func NewUserHandler(restaurantClient pb.RestaurantServiceClient, donorClient pb.DonorServiceClient) *UserHandler {
	return &UserHandler{
		UserRepo:       repository.NewUserRepository(),
		RestaurantGRPC: restaurantClient,
		DonorGRPC: 		donorClient,
	}
}

func (h *UserHandler) RegisterUser(c echo.Context) error {
	req := new(dtos.RegisterUserRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if req.Email == "" || req.Password == "" || req.UserType == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Missing required fields"})
	}

	// Transaction for user creation
	tx := h.UserRepo.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Save user to the database
	user := &models.User{
		Email:    req.Email,
		Password: req.Password,
		UserType: req.UserType,
	}
	if err := h.UserRepo.CreateUser(tx, user); err != nil {
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to register user"})
	}

	// Create Restaurant
	if req.UserType == "restaurant" {
		if req.Name == "" || req.Address == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Name and address are required for user_type 'restaurant'",
			})
		}

		restaurantRequest := &pb.PrepareRestaurantRequest{
			UserId:  user.ID,
			Email:   user.Email,
			Name:    req.Name,
			Address: req.Address,
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		response, err := h.RestaurantGRPC.PrepareRestaurant(ctx, restaurantRequest)
		if err != nil || !response.Success {
			tx.Rollback()
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to create restaurant via gRPC: " + response.GetMessage(),
			})
		}

		if err := h.UserRepo.CommitTransaction(tx); err != nil {
			_, _ = h.RestaurantGRPC.RollbackRestaurant(ctx, &pb.RollbackRestaurantRequest{UserId: user.ID})
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to commit user transaction"})
		}

		_, err = h.RestaurantGRPC.CommitRestaurant(ctx, &pb.CommitRestaurantRequest{UserId: user.ID})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to commit restaurant via gRPC"})
		}
	} else if req.UserType == "donor" {
		if req.Name == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Name are required for user_type 'donor'",
			})
		}

		donorRequest := &pb.PrepareDonorRequest{
			UserId:  user.ID,
			Name:    req.Name,
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		response, err := h.DonorGRPC.PrepareDonor(ctx, donorRequest)
		if err != nil || !response.Success {
			tx.Rollback()
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to create donor via gRPC: " + response.GetMessage(),
			})
		}

		if err := h.UserRepo.CommitTransaction(tx); err != nil {
			_, _ = h.DonorGRPC.RollbackDonor(ctx, &pb.RollbackDonorRequest{UserId: user.ID})
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to commit user transaction"})
		}

		_, err = h.DonorGRPC.CommitDonor(ctx, &pb.CommitDonorRequest{UserId: user.ID})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to commit donor via gRPC"})
		}
	}
	return c.JSON(http.StatusCreated, map[string]string{"message": "User registered successfully"})
}

func (h *UserHandler) LoginUser(c echo.Context) error {
	req := new(models.User)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	user, err := h.UserRepo.FindByEmail(req.Email)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid email or password"})
	}

	token, err := utils.GenerateJWT(user.ID, user.UserType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to generate token"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Login successful",
		"token":   token,
	})

}