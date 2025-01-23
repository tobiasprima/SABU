package handlers

import (
	"context"
	"restaurant-service-grpc/models"
	"restaurant-service-grpc/proto/pb"
	"restaurant-service-grpc/repository"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RestaurantGRpcHandler struct {
	pb.UnimplementedRestaurantServiceServer
	RestaurantRepo      *repository.RestaurantRepository
	preparedRestaurants map[string]*models.Restaurant
	preparedMeals       map[string]*models.Meal
	mu                  sync.Mutex
}

func NewRestaurantGRpcHandler() *RestaurantGRpcHandler {
	return &RestaurantGRpcHandler{
		RestaurantRepo:      repository.NewRestaurantRepository(),
		preparedRestaurants: make(map[string]*models.Restaurant),
		preparedMeals:       make(map[string]*models.Meal),
	}
}

func (h *RestaurantGRpcHandler) PrepareRestaurant(ctx context.Context, req *pb.PrepareRestaurantRequest) (*pb.PrepareRestaurantResponse, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, exists := h.preparedRestaurants[req.UserId]; exists {
		return &pb.PrepareRestaurantResponse{
			Success: false,
			Message: "Restaurant preparation already exists for this user",
		}, nil
	}

	h.preparedRestaurants[req.UserId] = &models.Restaurant{
		UserID:  req.UserId,
		Email:   req.Email,
		Name:    req.Name,
		Address: req.Address,
	}

	return &pb.PrepareRestaurantResponse{
		Success: true,
		Message: "Restaurant prepared successfully",
	}, nil
}

func (h *RestaurantGRpcHandler) CommitRestaurant(ctx context.Context, req *pb.CommitRestaurantRequest) (*pb.CommitRestaurantResponse, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	restaurant, exists := h.preparedRestaurants[req.UserId]
	if !exists {
		return &pb.CommitRestaurantResponse{
			Success: false,
			Message: "No Prepared restaurant found for this user",
		}, nil
	}

	if err := h.RestaurantRepo.CreateRestaurant(restaurant); err != nil {
		return &pb.CommitRestaurantResponse{
			Success: false,
			Message: "Failed to commit restaurant: " + err.Error(),
		}, nil
	}

	delete(h.preparedRestaurants, req.UserId)

	return &pb.CommitRestaurantResponse{
		Success: true,
		Message: "Restaurant committed successfully",
	}, nil
}

func (h *RestaurantGRpcHandler) RollbackRestaurant(ctx context.Context, req *pb.RollbackRestaurantRequest) (*pb.RollbackRestaurantResponse, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, exists := h.preparedRestaurants[req.UserId]; exists {
		delete(h.preparedRestaurants, req.UserId)
		return &pb.RollbackRestaurantResponse{
			Success: true,
			Message: "Restaurant rollback successfully",
		}, nil
	}

	return &pb.RollbackRestaurantResponse{
		Success: false,
		Message: "No prepared restaurant found for this user",
	}, nil
}

func (h *RestaurantGRpcHandler) GetMealByID(ctx context.Context, req *pb.MealID) (*pb.GetMealByIDResponse, error) {
	mealTmp, err := h.RestaurantRepo.GetMealByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	if mealTmp == nil {
		return nil, status.Error(codes.NotFound, "Meals not found")
	}

	meal := &pb.GetMealByIDResponse{
		Id:           mealTmp.ID,
		RestaurantId: mealTmp.RestaurantID,
		Name:         mealTmp.Name,
		Description:  mealTmp.Description,
		Price:        float32(mealTmp.Price),
		Stock:        uint32(mealTmp.Stock),
	}

	return meal, nil
}

func (h *RestaurantGRpcHandler) PrepareDeductMealStock(ctx context.Context, req *pb.PrepareDeductMealStockRequest) (*pb.PrepareDeductMealStockResponse, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, exists := h.preparedMeals[req.DonationId]; exists {
		return &pb.PrepareDeductMealStockResponse{
			Success: false,
			Message: "Meal preparation already exists for this donation",
		}, nil
	}

	h.preparedMeals[req.DonationId] = &models.Meal{
		ID:    req.MealId,
		Stock: int(req.Quantity),
	}

	return &pb.PrepareDeductMealStockResponse{
		Success: true,
		Message: "Meal prepared successfully",
	}, nil
}

func (h *RestaurantGRpcHandler) CommitDeductMealStock(ctx context.Context, req *pb.CommitDeductMealStockRequest) (*pb.CommitDeductMealStockResponse, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	meal, exists := h.preparedMeals[req.DonationId]
	if !exists {
		return &pb.CommitDeductMealStockResponse{
			Success: false,
			Message: "No Prepared meal found for this donation",
		}, nil
	}

	if err := h.RestaurantRepo.DeductMealStock(ctx, meal); err != nil {
		return &pb.CommitDeductMealStockResponse{
			Success: false,
			Message: "Failed to commit meal: " + err.Error(),
		}, nil
	}

	delete(h.preparedMeals, req.DonationId)

	return &pb.CommitDeductMealStockResponse{
		Success: true,
		Message: "Meal committed successfully",
	}, nil
}

func (h *RestaurantGRpcHandler) RollbackDeductMealStock(ctx context.Context, req *pb.RollbackDeductMealStockRequest) (*pb.RollbackDeductMealStockResponse, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, exists := h.preparedMeals[req.DonationId]; exists {
		delete(h.preparedMeals, req.DonationId)
		return &pb.RollbackDeductMealStockResponse{
			Success: true,
			Message: "Meal rollback successfully",
		}, nil
	}

	return &pb.RollbackDeductMealStockResponse{
		Success: false,
		Message: "No prepared meal found for this donation",
	}, nil
}
