package handlers

import (
	"context"
	"donor-service-grpc/models"
	"donor-service-grpc/proto/pb"
	"donor-service-grpc/repository"
	"sync"
)

type DonorGrpcHandler struct {
	pb.UnimplementedDonorServiceServer
	DonorRepository repository.DonorRepository
	preparedDonors  map[string]*models.Donor
	mu              sync.Mutex
}

func NewDonorGrpcHandlerImpl(donorRepository repository.DonorRepository) *DonorGrpcHandler {
	return &DonorGrpcHandler{
		DonorRepository: donorRepository,
		preparedDonors:  make(map[string]*models.Donor),
	}
}

func (dh *DonorGrpcHandler) PrepareDonor(ctx context.Context, req *pb.PrepareDonorRequest) (*pb.PrepareDonorResponse, error) {
	dh.mu.Lock()
	defer dh.mu.Unlock()

	if _, exists := dh.preparedDonors[req.UserId]; exists {
		return &pb.PrepareDonorResponse{
			Success: false,
			Message: "Donor preparation already exists for this user",
		}, nil
	}

	dh.preparedDonors[req.UserId] = &models.Donor{
		UserID: req.UserId,
		Name:   req.Name,
	}

	return &pb.PrepareDonorResponse{
		Success: true,
		Message: "Donor prepared successfully",
	}, nil
}

func (dh *DonorGrpcHandler) CommitDonor(ctx context.Context, req *pb.CommitDonorRequest) (*pb.CommitDonorResponse, error) {
	dh.mu.Lock()
	defer dh.mu.Unlock()

	donor, exists := dh.preparedDonors[req.UserId]
	if !exists {
		return &pb.CommitDonorResponse{
			Success: false,
			Message: "No prepared donor found for this user",
		}, nil
	}

	if err := dh.DonorRepository.Create(donor); err != nil {
		return &pb.CommitDonorResponse{
			Success: false,
			Message: "Failed to commit donor: " + err.Error(),
		}, nil
	}

	delete(dh.preparedDonors, req.UserId)

	return &pb.CommitDonorResponse{
		Success: true,
		Message: "Donor committed successfully",
	}, nil
}

func (dh *DonorGrpcHandler) RollbackDonor(ctx context.Context, req *pb.RollbackDonorRequest) (*pb.RollbackDonorResponse, error) {
	dh.mu.Lock()
	defer dh.mu.Unlock()

	if _, exists := dh.preparedDonors[req.UserId]; exists {
		delete(dh.preparedDonors, req.UserId)
		return &pb.RollbackDonorResponse{
			Success: true,
			Message: "Donor rollback successfully",
		}, nil
	}

	return &pb.RollbackDonorResponse{
		Success: false,
		Message: "No prepared donor found for this use",
	}, nil
}
