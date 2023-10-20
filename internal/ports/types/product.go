package types

import (
	"time"

	"gorm.io/gorm"
)

// Product is the model for products table
type Product struct {
	// ID is the primary key
	ID uint `gorm:"primaryKey"`
	// Name is the name of the product
	Name string `gorm:"type:varchar(100);NOT NULL"`
	// Description is the description of the product
	Description string `gorm:"type:varchar(200);NOT NULL"`
	// Price is the price of the product
	Price float64 `gorm:"NOT NULL"`
	// CreatedAt is the time when the product was created
	CreatedAt time.Time
	// UpdatedAt is the time when the product was updated
	UpdatedAt time.Time
	// DeletedAt is the time when the product was deleted (soft delete)
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TableName returns the table name for the Product model
func (Product) TableName() string {
	return "products"
}
