package repository

import "donor-service/models"

type DonationRepository interface {
	// CreateDonation()
	GetDonationHistory(donorID string) ([]models.Donation, error)
}
