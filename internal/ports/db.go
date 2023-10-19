package ports

import (
	"github.com/mirshahriar/marketplace/helper/errors"
	"github.com/mirshahriar/marketplace/internal/ports/types"
)

// DBPort is the interface for all DB ports
type DBPort interface {
	ListProduct(page types.PageReq, sort types.SortReq) ([]types.Product, errors.Error)
	InsertProduct(product types.Product) (types.Product, errors.Error)
	GetProductByID(productID uint) (types.Product, errors.Error)
	UpdateProduct(productID uint, update map[string]interface{}) errors.Error
	DeleteProduct(productID uint) errors.Error

	CountProduct() (int, errors.Error)
}
