package app

import (
	"github.com/mirshahriar/marketplace/helper/errors"
	"github.com/mirshahriar/marketplace/internal/ports/types"
)

func (a Adapter) CreateProduct(body types.ProductBody) (types.ProductResponse, errors.Error) {
	product, cErr := a.db.InsertProduct(body.ToProduct())
	if cErr != nil {
		return types.ProductResponse{}, cErr
	}

	return product.ToProductResponse(), nil
}

func (a Adapter) ListProduct(page types.PageReq, sort types.SortReq) (types.Page[types.ProductResponse], errors.Error) {
	products, cErr := a.db.ListProduct(page, sort)
	if cErr != nil {
		return types.Page[types.ProductResponse]{}, cErr
	}

	var resp []types.ProductResponse
	for _, product := range products {
		resp = append(resp, product.ToProductResponse())
	}

	total, err := a.db.CountProduct()
	if err != nil {
		return types.Page[types.ProductResponse]{}, err
	}

	return types.Page[types.ProductResponse]{
		Data:  resp,
		Page:  page.Page,
		Size:  len(resp),
		Total: total,
	}, nil
}

func (a Adapter) GetProductByID(productID uint) (types.ProductResponse, errors.Error) {
	product, cErr := a.db.GetProductByID(productID)
	if cErr != nil {
		return types.ProductResponse{}, cErr
	}

	return product.ToProductResponse(), nil
}

func (a Adapter) UpdateProduct(productID uint, body types.ProductBody) errors.Error {
	_, cErr := a.db.GetProductByID(productID)
	if cErr != nil {
		return cErr
	}

	cErr = a.db.UpdateProduct(productID, body.ToMap())
	if cErr != nil {
		return cErr
	}

	return nil
}

func (a Adapter) DeleteProduct(productID uint) errors.Error {
	_, cErr := a.db.GetProductByID(productID)
	if cErr != nil {
		return cErr
	}

	if cErr = a.db.DeleteProduct(productID); cErr != nil {
		return cErr
	}

	return nil
}
