package repository

import (
	"context"

	"github.com/jasonyangmh/sayakaya/model"
	"gorm.io/gorm"
)

type promoRepository struct {
	db *gorm.DB
}

type PromoRepository interface {
	Find(ctx context.Context) ([]model.Promo, error)
	FindLatest(ctx context.Context) (*model.Promo, error)
	FindByCode(ctx context.Context, userPromo *model.UserPromo) (*model.UserPromo, error)
	Create(ctx context.Context, promo *model.Promo) (*model.Promo, error)
	CreateForUser(ctx context.Context, userPromo *model.UserPromo) (*model.UserPromo, error)
	Update(ctx context.Context, userPromo *model.UserPromo) (*model.UserPromo, error)
}

func NewPromoRepository(db *gorm.DB) PromoRepository {
	return &promoRepository{
		db: db,
	}
}

func (r *promoRepository) Find(ctx context.Context) ([]model.Promo, error) {
	promos := []model.Promo{}

	if err := r.db.WithContext(ctx).Find(&promos).Error; err != nil {
		return nil, err
	}

	return promos, nil
}

func (r *promoRepository) FindLatest(ctx context.Context) (*model.Promo, error) {
	promo := &model.Promo{}

	if err := r.db.WithContext(ctx).Last(&promo).Error; err != nil {
		return nil, err
	}

	return promo, nil
}

func (r *promoRepository) FindByCode(ctx context.Context, userPromo *model.UserPromo) (*model.UserPromo, error) {
	if err := r.db.WithContext(ctx).Preload("User").Preload("Promo").Where("code = ?", userPromo.Code).First(&userPromo).Error; err != nil {
		return nil, err
	}

	return userPromo, nil
}

func (r *promoRepository) Create(ctx context.Context, promo *model.Promo) (*model.Promo, error) {
	if err := r.db.WithContext(ctx).Create(&promo).Error; err != nil {
		return nil, err
	}

	return promo, nil
}

func (r *promoRepository) CreateForUser(ctx context.Context, userPromo *model.UserPromo) (*model.UserPromo, error) {
	if err := r.db.WithContext(ctx).Create(&userPromo).Error; err != nil {
		return nil, err
	}

	return userPromo, nil
}

func (r *promoRepository) Update(ctx context.Context, userPromo *model.UserPromo) (*model.UserPromo, error) {
	if err := r.db.WithContext(ctx).Model(&userPromo).Update("is_used", userPromo.IsUsed).Error; err != nil {
		return nil, err
	}

	return userPromo, nil
}
