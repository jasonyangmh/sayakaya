package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email      string    `gorm:"type:varchar;unique;not null"`
	Birthday   time.Time `gorm:"type:date;not null"`
	IsVerified bool      `gorm:"type:boolean;default:false"`
}
