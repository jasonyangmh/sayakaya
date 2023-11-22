package model

import "gorm.io/gorm"

type UserPromo struct {
	gorm.Model
	Code    string `gorm:"type:varchar;not null"`
	IsUsed  bool   `gorm:"type:boolean;default:false"`
	UserID  uint   `gorm:"type:bigint;not null"`
	PromoID uint   `gorm:"type:bigint;not null"`
	User    User   `gorm:"foreignKey:UserID"`
	Promo   Promo  `gorm:"foreignKey:PromoID"`
}
