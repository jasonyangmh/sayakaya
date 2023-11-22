package usecase

import (
	"context"
	"time"

	"github.com/jasonyangmh/sayakaya/model"
	"github.com/jasonyangmh/sayakaya/repository"
	"github.com/jasonyangmh/sayakaya/shared"
)

type userUsecase struct {
	userRepository repository.UserRepository
}

type UserUsecase interface {
	FindUsers(ctx context.Context) ([]model.User, error)
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	CheckUserBirthday(ctx context.Context, user *model.User) bool
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: ur,
	}
}

func (u *userUsecase) FindUsers(ctx context.Context) ([]model.User, error) {
	return u.userRepository.Find(ctx)
}

func (u *userUsecase) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	if _, err := u.userRepository.FindByEmail(ctx, user); err == nil {
		return nil, shared.ErrEmailAlreadyRegistered
	}

	return u.userRepository.Create(ctx, user)
}

func (u *userUsecase) CheckUserBirthday(ctx context.Context, user *model.User) bool {
	today := time.Now()

	if user.Birthday.Month() == today.Month() && user.Birthday.Day() == today.Day() {
		return true
	}

	return false
}
