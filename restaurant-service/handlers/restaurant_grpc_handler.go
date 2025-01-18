package handlers

import (
	"context"
	"sabu-restaurant-service/models"
	"sabu-restaurant-service/proto/pb"
	"sabu-restaurant-service/repository"
	"sync"
)

type RestaurantGRpcHandler struct {
	pb.UnimplementedRestaurantServiceServer
	RestaurantRepo		*repository.RestaurantRepository
	preparedRestaurants	map[string]*models.Restaurant
	mu					sync.Mutex
}

func NewRestaurantGRpcHandler() *RestaurantGRpcHandler {
	return &RestaurantGRpcHandler{
		RestaurantRepo: repository.NewRestaurantRepository(),
		preparedRestaurants: make(map[string]*models.Restaurant),
	}
}

func (h *RestaurantGRpcHandler) PrepareRestaurant(ctx context.Context, req *pb.PrepareRestaurantRequest) (*pb.PrepareRestaurantResponse, error){
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

func (h *RestaurantGRpcHandler) CommitRestaurant(ctx context.Context, req *pb.CommitRestaurantRequest) (*pb.CommitRestaurantResponse, error){
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