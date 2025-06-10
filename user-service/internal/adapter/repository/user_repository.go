package repository

import (
	"context"
	"errors"
	"user-service/internal/core/domain/entity"
	"user-service/internal/core/domain/model"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryInterface {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) GetUserByEmail(ctx context.Context, email string) (*entity.UserEntity, error) {
	var user model.User
	if err := u.db.Where("email = ? AND is_verified = ?", email, true).Preload("Roles").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = errors.New("404")
			log.Errorf("[UserRepository-1] GetUserByEmail: User not found")
			return nil, err
		}
		log.Errorf("[UserRepository-1] GetUserByEmail: %v", err)
		return nil, err
	}

	return &entity.UserEntity{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Password:   user.Password,
		RoleName:   user.Roles[0].Name,
		Address:    user.Address,
		Lat:        user.Lat,
		Lng:        user.Lng,
		Phone:      user.Phone,
		Photo:      user.Photo,
		IsVerified: user.IsVerified,
	}, nil
}
