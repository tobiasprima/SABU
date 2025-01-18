package handlers

import (
	"context"
	"foundation-service/models"
	"foundation-service/proto/pb"
	"foundation-service/repository"
	"sync"
)

type FoundationGrpcHandler struct {
	pb.UnimplementedFoundationServiceServer
	FoundationRepository repository.FoundationRepository
	preparedFoundations  map[string]*models.Foundation
	mu                   sync.Mutex
}

func NewFoundationGrpcHandlerImpl(foundationRepository repository.FoundationRepository) *FoundationGrpcHandler {
	return &FoundationGrpcHandler{
		FoundationRepository: foundationRepository,
		preparedFoundations:  make(map[string]*models.Foundation),
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

	if err := fh.FoundationRepository.Create(foundation); err != nil {
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
