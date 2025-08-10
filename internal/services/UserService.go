package services

import (
	"main/internal/models"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

func (s *UserService) CreateUserWithProfile(userData *models.User, profileData *models.Profile) (*models.User, error) {
	// Start a transaction
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create the user
	if err := tx.Create(userData).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Associate the profile with the user
	profileData.UserID = userData.ID
	if err := tx.Create(profileData).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	// Reload the user to get the associated profile
	var createdUser models.User
	if err := s.db.Preload("Profile").First(&createdUser, userData.ID).Error; err != nil {
		return nil, err
	}

	return &createdUser, nil
}
