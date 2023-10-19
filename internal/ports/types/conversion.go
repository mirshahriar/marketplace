package types

import "github.com/fatih/structs"

func (p ProductBody) ToProduct() Product {
	return Product{
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
	}
}

func (p Product) ToProductResponse() ProductResponse {
	return ProductResponse{
		ID: p.ID,
		ProductBody: ProductBody{
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
		},
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	}
}

func (p ProductBody) ToMap() map[string]interface{} {
	return structs.Map(p)
}
