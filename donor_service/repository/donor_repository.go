package repository

import (
	"donor-service/models"
	"errors"
	"time"

	"gorm.io/gorm"
)

type DonorRepository interface {
	Create(donor *models.Donor) error
	GetDonorByID(donorID string) (*models.Donor, error)
	UpdateDonorBalance(donorID string, amount float64) error
	TopUp(topUp *models.TopUp) error
	GetTopUpByID(topUpID string) (*models.TopUp, error)
	GetTopUpHistory(donorID string) ([]models.TopUp, error)
	UpdateTopUpStatus(topUp *models.TopUp, completedAt time.Time, status, paymentMethod string) error
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

func (dr *DonorRepositoryImpl) UpdateTopUpStatus(topUp *models.TopUp, completedAt time.Time, status, paymentMethod string) error {
	return dr.DB.Model(topUp).Select("payment_method", "status", "completed_at").Updates(models.TopUp{PaymentMethod: paymentMethod, Status: status, CompletedAt: &completedAt}).Error
}
