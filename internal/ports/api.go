// Package ports declares all PORT interface
package ports

import (
	"github.com/mirshahriar/marketplace/helper/errors"
	"github.com/mirshahriar/marketplace/internal/ports/types"
)

// APIPort is the interface for all API ports
type APIPort interface {
	ListProduct(page types.PageReq, sort types.SortReq) (types.Page[types.ProductResponse], errors.Error)
	CreateProduct(body types.ProductBody) (types.ProductResponse, errors.Error)
	GetProductByID(productID uint) (types.ProductResponse, errors.Error)
	UpdateProduct(productID uint, body types.ProductBody) errors.Error
	DeleteProduct(productID uint) errors.Error
}
