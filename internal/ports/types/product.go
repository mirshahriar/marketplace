package types

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint    `gorm:"primaryKey"`
	Name        string  `gorm:"type:varchar(100);NOT NULL"`
	Description string  `gorm:"type:varchar(200);NOT NULL"`
	Price       float64 `gorm:"NOT NULL"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Product) TableName() string {
	return "products"
}
