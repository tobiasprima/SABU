package repository

import (
	"donor-service/models"
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DonorRepository interface {
	Create(donor *models.Donor) error
	GetDonorByID(donorID string) (*models.Donor, error)
	UpdateDonorBalance(donorID string, amount float64) error
	TopUp(topUp *models.TopUp) error
	GetTopUpByID(topUpID string) (*models.TopUp, error)
	GetTopUpHistory(donorID string) ([]models.TopUp, error)
	UpdateTopUpStatus(topUpID, status, paymentMethod string, completedAt time.Time) (*models.TopUp, error)
}

type DonorRepositoryImpl struct {
	DB *gorm.DB
}

func NewDonorRepositoryImpl(db *gorm.DB) DonorRepository {
	return &DonorRepositoryImpl{db}
}

func (dr *DonorRepositoryImpl) Create(donor *models.Donor) error {
	return dr.DB.Create(donor).Error
}

func (dr *DonorRepositoryImpl) GetDonorByID(donorID string) (*models.Donor, error) {
	donor := new(models.Donor)

	if err := dr.DB.Where("id = ?", donorID).Take(donor).Error; err != nil {
		return nil, err
	}

	return donor, nil
}

func (dr *DonorRepositoryImpl) UpdateDonorBalance(donorID string, amount float64) error {
	donor := models.Donor{ID: donorID}
	res := dr.DB.Model(&donor).Update("balance", gorm.Expr("balance + ?", amount))
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("not found")
	}

	return nil
}

func (dr *DonorRepositoryImpl) TopUp(topUp *models.TopUp) error {
	return dr.DB.Create(topUp).Error
}

func (dr *DonorRepositoryImpl) GetTopUpByID(topUpID string) (*models.TopUp, error) {
	topUp := new(models.TopUp)

	if err := dr.DB.Where("id = ?", topUpID).Take(topUp).Error; err != nil {
		return nil, err
	}

	return topUp, nil
}

func (dr *DonorRepositoryImpl) GetTopUpHistory(donorID string) ([]models.TopUp, error) {
	var topUps []models.TopUp

	if err := dr.DB.Where("donor_id = ?", donorID).Find(&topUps).Error; err != nil {
		return nil, err
	}

	return topUps, nil
}

func (dr *DonorRepositoryImpl) UpdateTopUpStatus(topUpID, status, paymentMethod string, completedAt time.Time) (*models.TopUp, error) {
	topUp := models.TopUp{ID: topUpID}
	res := dr.DB.Model(&topUp).Clauses(clause.Returning{Columns: []clause.Column{{Name: "donor_id"}, {Name: "amount"}}}).
		Select("payment_method", "status", "completed_at").
		Updates(models.TopUp{PaymentMethod: paymentMethod, Status: status, CompletedAt: &completedAt})
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	return &topUp, nil
}
