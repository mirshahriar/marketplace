package http

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mirshahriar/marketplace/helper/errors"
	"github.com/mirshahriar/marketplace/internal/ports/types"
)

func (a Adapter) CreateProduct(c echo.Context) error {

	var body types.ProductBody

	if err := c.Bind(&body); err != nil {
		cErr := errors.InvalidRequestParsingError(err)
		return c.JSON(cErr.Status(), cErr)
	}

	if cErr := types.Validate(&body, c.Request().Method); cErr != nil {
		return c.JSON(cErr.Status(), cErr)
	}

	resp, cErr := a.api.CreateProduct(body)
	if cErr != nil {
		return c.JSON(cErr.Status(), cErr)
	}

	return c.JSON(200, resp)

}

func (a Adapter) ListProduct(c echo.Context) error {
	params := struct {
		types.PageReq
		types.SortReq
	}{
		PageReq: types.NewPageReq(a.config.PaginationSize),
		SortReq: types.NewSortReq(),
	}

	if err := a.binder.BindQueryParams(c, &params); err != nil {
		cErr := errors.InvalidRequestParsingError(err)
		return c.JSON(cErr.Status(), cErr)
	}

	resp, cErr := a.api.ListProduct(params.PageReq, params.SortReq)
	if cErr != nil {
		return c.JSON(cErr.Status(), cErr)
	}

	return c.JSON(200, resp)
}

func (a Adapter) GetProduct(c echo.Context) error {
	param := struct {
		ProductID uint `param:"product"`
	}{}

	if err := a.binder.BindPathParams(c, &param); err != nil {
		cErr := errors.InvalidRequestParsingError(err)
		return c.JSON(cErr.Status(), cErr)
	}

	resp, cErr := a.api.GetProductByID(param.ProductID)
	if cErr != nil {
		return c.JSON(cErr.Status(), cErr)
	}

	return c.JSON(200, resp)
}

func (a Adapter) UpdateProduct(c echo.Context) error {
	params := struct {
		ProductID uint `param:"product"`
		types.ProductBody
	}{}

	if err := c.Bind(&params); err != nil {
		cErr := errors.InvalidRequestParsingError(err)
		return c.JSON(cErr.Status(), cErr)
	}

	if cErr := types.Validate(&params.ProductBody, c.Request().Method); cErr != nil {
		return c.JSON(cErr.Status(), cErr)
	}

	cErr := a.api.UpdateProduct(params.ProductID, params.ProductBody)
	if cErr != nil {
		return c.JSON(cErr.Status(), cErr)
	}

	return c.NoContent(200)
}

func (a Adapter) DeleteProduct(c echo.Context) error {
	param := struct {
		ProductID uint `param:"product"`
	}{}

	if err := a.binder.BindPathParams(c, &param); err != nil {
		cErr := errors.InvalidRequestParsingError(err)
		return c.JSON(cErr.Status(), cErr)
	}

	cErr := a.api.DeleteProduct(param.ProductID)
	if cErr != nil {
		fmt.Println("------ failed", cErr.Status())
		return c.JSON(cErr.Status(), cErr)
	}

	fmt.Println("------ success", 200)

	return c.NoContent(200)
}
