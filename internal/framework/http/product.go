// Package http implements HTTP server
// nolint: wrapcheck
package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mirshahriar/marketplace/helper/errors"
	"github.com/mirshahriar/marketplace/internal/ports/types"
)

// CreateProduct creates a product with valid data
// @Summary	Creates a product
// @Description	Creates a product with valid data
// @Tags	product
// @Accept	json
// @Produce	json
// @param Authorization header string true "Authorization"
// @Param	product	body	types.ProductBody	true	"Request body of Product"
// @Success	200	{object}	types.ProductResponse
// @Failure	400	{object}	errors.CustomError
// @Failure	403	{object}	errors.CustomError
// @Router	/products [POST]
func (a Adapter) CreateProduct(c echo.Context) error {

	var body types.ProductBody

	if err := c.Bind(&body); err != nil {
		cErr := errors.InvalidRequestParsingError(err)
		return c.JSONPretty(cErr.Status(), cErr, " ")
	}

	// Validating the request body.
	if cErr := types.Validate(&body, c.Request().Method); cErr != nil {
		return c.JSONPretty(cErr.Status(), cErr, " ")
	}

	// Creating the product and returning the response.
	resp, cErr := a.api.CreateProduct(body)
	if cErr != nil {
		return c.JSONPretty(cErr.Status(), cErr, " ")
	}

	return c.JSONPretty(http.StatusOK, resp, " ")
}

// ListProduct lists all products
// @Summary	Lists all products
// @Description	Lists all products
// @Tags	product
// @Produce	json
// @Param	page	query	int	false	"Page number"
// @Param	size	query	int	false	"Page size"
// @Param	sort_by	query	string	false	"Sort by"
// @Param	sort_direction	query	string	false	"Order by"
// @Success	200	{object}	types.Page[types.ProductResponse]
// @Failure	400	{object}	errors.CustomError
// @Router	/products [GET]
func (a Adapter) ListProduct(c echo.Context) error {
	params := struct {
		types.PageReq
		types.SortReq
	}{
		// Setting the default page size.
		PageReq: types.NewPageReq(a.config.PaginationSize),
		SortReq: types.NewSortReq(),
	}

	if err := a.binder.BindQueryParams(c, &params); err != nil {
		cErr := errors.InvalidRequestParsingError(err)
		return c.JSONPretty(cErr.Status(), cErr, " ")
	}

	// Listing the products and returning the response.
	resp, cErr := a.api.ListProduct(params.PageReq, params.SortReq)
	if cErr != nil {
		return c.JSONPretty(cErr.Status(), cErr, " ")
	}

	return c.JSONPretty(http.StatusOK, resp, " ")
}

// GetProduct gets a product by ID
// @Summary	Gets a product by ID
// @Description	Gets a product by ID
// @Tags	product
// @Produce	json
// @Param	product	path	int	true	"Product ID"
// @Success	200	{object}	types.ProductResponse
// @Failure	400	{object}	errors.CustomError
// @Router	/products/{product} [GET]
func (a Adapter) GetProduct(c echo.Context) error {
	param := struct {
		ProductID uint `param:"product"`
	}{}

	if err := a.binder.BindPathParams(c, &param); err != nil {
		cErr := errors.InvalidRequestParsingError(err)
		return c.JSONPretty(cErr.Status(), cErr, " ")
	}

	// Getting the product by its ID and returning the response.
	resp, cErr := a.api.GetProduct(param.ProductID)
	if cErr != nil {
		return c.JSONPretty(cErr.Status(), cErr, " ")
	}

	return c.JSONPretty(http.StatusOK, resp, " ")
}

// UpdateProduct updates a product by ID
// @Summary	Updates a product by ID
// @Description	Updates a product by ID
// @Tags	product
// @Accept	json
// @Produce	json
// @Param	product	path	int	true	"Product ID"
// @Param	product	body	types.ProductBody	true	"Request body of Product"
// @param Authorization header string true "Authorization"
// @Success	200
// @Failure	400	{object}	errors.CustomError
// @Failure	403	{object}	errors.CustomError
// @Failure	404	{object}	errors.CustomError
// @Router	/products/{product} [PUT]
func (a Adapter) UpdateProduct(c echo.Context) error {
	params := struct {
		ProductID uint `param:"product"`
		types.ProductBody
	}{}

	if err := c.Bind(&params); err != nil {
		cErr := errors.InvalidRequestParsingError(err)
		return c.JSONPretty(cErr.Status(), cErr, " ")
	}

	if cErr := types.Validate(&params.ProductBody, c.Request().Method); cErr != nil {
		return c.JSON(cErr.Status(), cErr)
	}

	cErr := a.api.UpdateProduct(params.ProductID, params.ProductBody)
	if cErr != nil {
		return c.JSONPretty(cErr.Status(), cErr, " ")
	}

	return c.NoContent(http.StatusOK)
}

// DeleteProduct deletes a product by ID
// @Summary	Deletes a product by ID
// @Description	Deletes a product by ID
// @Tags	product
// @Produce	json
// @Param	product	path	int	true	"Product ID"
// @param Authorization header string true "Authorization"
// @Success	200
// @Failure	400	{object}	errors.CustomError
// @Failure	403	{object}	errors.CustomError
// @Failure	404	{object}	errors.CustomError
// @Router	/products/{product} [DELETE]
func (a Adapter) DeleteProduct(c echo.Context) error {
	param := struct {
		ProductID uint `param:"product"`
	}{}

	if err := a.binder.BindPathParams(c, &param); err != nil {
		cErr := errors.InvalidRequestParsingError(err)
		return c.JSONPretty(cErr.Status(), cErr, " ")
	}

	// Deleting the product by its ID.
	cErr := a.api.DeleteProduct(param.ProductID)
	if cErr != nil {
		return c.JSONPretty(cErr.Status(), cErr, " ")
	}

	return c.NoContent(http.StatusOK)
}
