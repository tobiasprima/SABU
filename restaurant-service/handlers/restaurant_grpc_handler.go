package handlers

import (
	"context"
	"sabu-restaurant-service/models"
	"sabu-restaurant-service/proto/pb"
	"sabu-restaurant-service/repository"
)

type RestaurantGRprcHandler struct {
	RestaurantRepo		*repository.RestaurantRepository
	pb.UnimplementedRestaurantServiceServer
}

func NewRestaurantGRpcHandler() *RestaurantGRprcHandler {
	return &RestaurantGRprcHandler{RestaurantRepo: repository.NewRestaurantRepository()}
}

func (h *RestaurantGRprcHandler) CreateRestaurant(ctx context.Context, req *pb.CreateRestaurantRequest) (*pb.CreateRestaurantResponse, error) {
	restaurant := &models.Restaurant{
		UserID:  req.UserId,
		Email:   req.Email,
		Name:    req.Name,
		Address: req.Address,
	}

	if err := h.RestaurantRepo.CreateRestaurant(restaurant); err != nil {
		return &pb.CreateRestaurantResponse{
			Success: false,
			Message: "Failed to create restaurant: " + err.Error(),
		}, nil
	}

	return &pb.CreateRestaurantResponse{
		Success: true,
		Message: "Restaurant created successfully",
	}, nil
}