package model

import (
	"time"

	"gorm.io/gorm"
)

type Promo struct {
	gorm.Model
	Name      string    `gorm:"type:varchar;not null"`
	StartDate time.Time `gorm:"type:date;not null"`
	EndDate   time.Time `gorm:"type:date;not null"`
	Amount    int64     `gorm:"type:bigint;not null"`
}
