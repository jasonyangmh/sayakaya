package usecase

import (
	"context"
	"time"

	"github.com/jasonyangmh/sayakaya/config"
	"github.com/jasonyangmh/sayakaya/model"
	"github.com/jasonyangmh/sayakaya/repository"
	"github.com/jasonyangmh/sayakaya/shared"
)

type promoUsecase struct {
	config          *config.Config
	promoRepository repository.PromoRepository
}

type PromoUsecase interface {
	FindPromos(ctx context.Context) ([]model.Promo, error)
	FindLatestPromo(ctx context.Context) (*model.Promo, error)
	FindPromoByCode(ctx context.Context, userPromo *model.UserPromo) (*model.UserPromo, error)
	CreatePromo(ctx context.Context, user *model.Promo) (*model.Promo, error)
	CreatePromoForUser(ctx context.Context, userPromo *model.UserPromo) (*model.UserPromo, error)
	CheckPromoEndDate(ctx context.Context, promo *model.Promo) bool
	GeneratePromo(ctx context.Context, userPromo *model.UserPromo) (*model.UserPromo, error)
	RedeemPromo(ctx context.Context, userPromo *model.UserPromo) (*model.UserPromo, error)
}

func NewPromoUsecase(cfg *config.Config, pr repository.PromoRepository) PromoUsecase {
	return &promoUsecase{
		config:          cfg,
		promoRepository: pr,
	}
}

func (u *promoUsecase) FindPromos(ctx context.Context) ([]model.Promo, error) {
	return u.promoRepository.Find(ctx)
}

func (u *promoUsecase) FindLatestPromo(ctx context.Context) (*model.Promo, error) {
	return u.promoRepository.FindLatest(ctx)
}

func (u *promoUsecase) FindPromoByCode(ctx context.Context, userPromo *model.UserPromo) (*model.UserPromo, error) {
	return u.promoRepository.FindByCode(ctx, userPromo)
}

func (u *promoUsecase) CreatePromo(ctx context.Context, promo *model.Promo) (*model.Promo, error) {
	return u.promoRepository.Create(ctx, promo)
}

func (u *promoUsecase) CreatePromoForUser(ctx context.Context, userPromo *model.UserPromo) (*model.UserPromo, error) {
	return u.promoRepository.CreateForUser(ctx, userPromo)
}

func (u *promoUsecase) CheckPromoEndDate(ctx context.Context, promo *model.Promo) bool {
	today := time.Now()

	return promo.EndDate.After(today)
}

func (u *promoUsecase) GeneratePromo(ctx context.Context, userPromo *model.UserPromo) (*model.UserPromo, error) {
	userPromo.Code = shared.GenerateCode(u.config.CodeLen)
	return u.promoRepository.CreateForUser(ctx, userPromo)
}

func (u *promoUsecase) RedeemPromo(ctx context.Context, userPromo *model.UserPromo) (*model.UserPromo, error) {
	userID := userPromo.UserID

	userPromo, err := u.promoRepository.FindByCode(ctx, userPromo)
	if err != nil || userPromo.UserID != userID {
		return nil, shared.ErrInvalidPromo
	}

	if userPromo.IsUsed {
		return nil, shared.ErrPromoIsUsed
	}

	if !userPromo.User.IsVerified {
		return nil, shared.ErrUserIsNotVerified
	}

	today := time.Now()

	if userPromo.Promo.EndDate.After(today) {
		return nil, shared.ErrPromoHasEnded
	}

	userPromo.IsUsed = true

	userPromo, err = u.promoRepository.Update(ctx, userPromo)
	if err != nil {
		return nil, err
	}

	return userPromo, nil
}
