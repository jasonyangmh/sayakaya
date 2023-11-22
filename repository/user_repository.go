package repository

import (
	"context"

	"github.com/jasonyangmh/sayakaya/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	Find(ctx context.Context) ([]model.User, error)
	FindByID(ctx context.Context, user *model.User) (*model.User, error)
	FindByEmail(ctx context.Context, user *model.User) (*model.User, error)
	Create(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Find(ctx context.Context) ([]model.User, error) {
	users := []model.User{}

	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) FindByID(ctx context.Context, user *model.User) (*model.User, error) {
	if err := r.db.WithContext(ctx).First(&user, user.Model.ID).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) FindByEmail(ctx context.Context, user *model.User) (*model.User, error) {
	if err := r.db.WithContext(ctx).Where("email = ?", user.Email).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Create(ctx context.Context, user *model.User) (*model.User, error) {
	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	if err := r.db.WithContext(ctx).Model(&user).Update("is_verified", user.IsVerified).Error; err != nil {
		return nil, err
	}

	return user, nil
}
