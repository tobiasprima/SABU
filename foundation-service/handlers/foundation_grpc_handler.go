package handlers

import (
	"context"
	"foundation-service/models"
	"foundation-service/proto/pb"
	"foundation-service/repository"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type FoundationGrpcHandler struct {
	pb.UnimplementedFoundationServiceServer
	FoundationRepository repository.FoundationRepository
	preparedFoundations  map[string]*models.Foundation
	preparedOrders       map[string]*models.Order
	mu                   sync.Mutex
}

func NewFoundationGrpcHandlerImpl(foundationRepository repository.FoundationRepository) *FoundationGrpcHandler {
	return &FoundationGrpcHandler{
		FoundationRepository: foundationRepository,
		preparedFoundations:  make(map[string]*models.Foundation),
		preparedOrders:       make(map[string]*models.Order),
	}
}

func (fh *FoundationGrpcHandler) PrepareFoundation(ctx context.Context, req *pb.PrepareFoundationRequest) (*pb.PrepareFoundationResponse, error) {
	fh.mu.Lock()
	defer fh.mu.Unlock()

	if _, exists := fh.preparedFoundations[req.UserId]; exists {
		return &pb.PrepareFoundationResponse{
			Success: false,
			Message: "Foundation preparation already exists for this user",
		}, nil
	}

	fh.preparedFoundations[req.UserId] = &models.Foundation{
		UserID: req.UserId,
		Name:   req.Name,
		Address: req.Address,
	}

	return &pb.PrepareFoundationResponse{
		Success: true,
		Message: "Foundation prepared successfully",
	}, nil
}

func (fh *FoundationGrpcHandler) CommitFoundation(ctx context.Context, req *pb.CommitFoundationRequest) (*pb.CommitFoundationResponse, error) {
	fh.mu.Lock()
	defer fh.mu.Unlock()

	foundation, exists := fh.preparedFoundations[req.UserId]
	if !exists {
		return &pb.CommitFoundationResponse{
			Success: false,
			Message: "No prepared foundation found for this user",
		}, nil
	}

	if err := fh.FoundationRepository.CreateFoundation(foundation); err != nil {
		return &pb.CommitFoundationResponse{
			Success: false,
			Message: "Failed to commit donor: " + err.Error(),
		}, nil
	}

	delete(fh.preparedFoundations, req.UserId)

	return &pb.CommitFoundationResponse{
		Success: true,
		Message: "Foundation comitted successfully",
	}, nil
}

func (fh *FoundationGrpcHandler) RollbackFoundation(ctx context.Context, req *pb.RollbackFoundationRequest) (*pb.RollbackFoundationResponse, error) {
	fh.mu.Lock()
	defer fh.mu.Unlock()

	if _, exists := fh.preparedFoundations[req.UserId]; exists {
		delete(fh.preparedFoundations, req.UserId)
		return &pb.RollbackFoundationResponse{
			Success: true,
			Message: "Foundation rollback successfully",
		}, nil
	}

	return &pb.RollbackFoundationResponse{
		Success: false,
		Message: "No prepared foundation found for this user",
	}, nil
}

func (fh *FoundationGrpcHandler) GetOrderByID(ctx context.Context, req *pb.OrderID) (*pb.GetOrderByIDResponse, error) {
	orderTmp, err := fh.FoundationRepository.GetOrderByID(req.Id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, status.Error(codes.NotFound, "Order not found")
		}
		return nil, err
	}

	order := &pb.GetOrderByIDResponse{
		Id:              orderTmp.ID,
		OrderListId:     orderTmp.OrderListID,
		MealsId:         orderTmp.MealsID,
		Quantity:        uint32(orderTmp.Quantity),
		DesiredQuantity: uint32(orderTmp.DesiredQuantity),
	}

	return order, nil
}

func (fh *FoundationGrpcHandler) PrepareAddOrderQuantity(ctx context.Context, req *pb.PrepareAddOrderQuantityRequest) (*pb.PrepareAddOrderQuantityResponse, error) {
	fh.mu.Lock()
	defer fh.mu.Unlock()

	if _, exists := fh.preparedOrders[req.DonationId]; exists {
		return &pb.PrepareAddOrderQuantityResponse{
			Success: false,
			Message: "Order preparation already exists for this donation",
		}, nil
	}

	fh.preparedOrders[req.DonationId] = &models.Order{
		ID:       req.OrderId,
		Quantity: int(req.Quantity),
	}

	return &pb.PrepareAddOrderQuantityResponse{
		Success: true,
		Message: "Order prepared successfully",
	}, nil
}

func (fh *FoundationGrpcHandler) CommitAddOrderQuantity(ctx context.Context, req *pb.CommitAddOrderQuantityRequest) (*pb.CommitAddOrderQuantityResponse, error) {
	fh.mu.Lock()
	defer fh.mu.Unlock()

	order, exists := fh.preparedOrders[req.DonationId]
	if !exists {
		return &pb.CommitAddOrderQuantityResponse{
			Success: false,
			Message: "No prepared order found for this donation",
		}, nil
	}

	if err := fh.FoundationRepository.AddOrderQuantity(order.ID, order.Quantity); err != nil {
		return &pb.CommitAddOrderQuantityResponse{
			Success: false,
			Message: "Failed to commit order: " + err.Error(),
		}, nil
	}

	delete(fh.preparedOrders, req.DonationId)

	return &pb.CommitAddOrderQuantityResponse{
		Success: true,
		Message: "Order comitted successfully",
	}, nil
}

func (fh *FoundationGrpcHandler) RollbackAddOrderQuantity(ctx context.Context, req *pb.RollbackAddOrderQuantityRequest) (*pb.RollbackAddOrderQuantityResponse, error) {
	fh.mu.Lock()
	defer fh.mu.Unlock()

	if _, exists := fh.preparedOrders[req.DonationId]; exists {
		delete(fh.preparedOrders, req.DonationId)
		return &pb.RollbackAddOrderQuantityResponse{
			Success: true,
			Message: "Order rollback successfully",
		}, nil
	}

	return &pb.RollbackAddOrderQuantityResponse{
		Success: false,
		Message: "No prepared order found for this donation",
	}, nil
}
