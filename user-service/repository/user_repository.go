package repository

import (
	"sabu-user-service/config"
	"sabu-user-service/models"
	"sabu-user-service/utils"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{DB: config.Database.User}
}

func (r *UserRepository) BeginTransaction() *gorm.DB {
	return r.DB.Begin()
}

func (r *UserRepository) CommitTransaction(tx *gorm.DB) error {
	return tx.Commit().Error
}

func (r *UserRepository) RollbackTransaction(tx *gorm.DB) error {
	return tx.Rollback().Error
}

func (r *UserRepository) CreateUser(tx *gorm.DB, user *models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return tx.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	if err := r.DB.Session(&gorm.Session{}).Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
