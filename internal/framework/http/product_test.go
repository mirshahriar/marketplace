package http_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mirshahriar/marketplace/config"
	"github.com/mirshahriar/marketplace/helper/errors"
	app_adapter "github.com/mirshahriar/marketplace/internal/application"
	http_adapter "github.com/mirshahriar/marketplace/internal/framework/http"
	db_adapter "github.com/mirshahriar/marketplace/internal/framework/mysql"
	"github.com/mirshahriar/marketplace/internal/ports/types"
	. "github.com/stretchr/testify/require"
)

func TestAdapter_CreateProduct(t *testing.T) {
	e := echo.New()

	dbAdapter, err := db_adapter.NewAdapterForTest(t)
	Nil(t, err)

	appAdapter := app_adapter.NewApplication(config.AppConfig{}, dbAdapter)

	httpAdapter := http_adapter.NewAdapter(config.AppConfig{}, appAdapter)

	type args struct {
		body types.ProductBody
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     errors.Error
	}{
		{
			name: "create product with no name",
			args: args{
				body: types.ProductBody{
					Description: "no name provided",
					Price:       100,
				},
			},
			wantErr: true,
			err:     errors.ValidationError("You must provide the name"),
		},
		{
			name: "create product with long name",
			args: args{
				body: types.ProductBody{
					Name:        "this is a very long name that is not allowed. so this product should not be created. hence return an validation error",
					Description: "test description",
					Price:       100,
				},
			},
			wantErr: true,
			err:     errors.ValidationError("You must provide a valid name with maximum length of 100"),
		},
		{
			name: "create product with negative price",
			args: args{
				body: types.ProductBody{
					Name:        "item 1",
					Description: "test description",
					Price:       -10,
				},
			},
			wantErr: true,
			err:     errors.ValidationError("You must provide a valid price greater than or equal to 0"),
		},
		{
			name: "create product successfully",
			args: args{
				body: types.ProductBody{
					Name:        "item 1",
					Description: "test description",
					Price:       10.5,
				},
			},
			wantErr: false,
		},
		{
			name: "create product successfully with zero price",
			args: args{
				body: types.ProductBody{
					Name:        "item 2",
					Description: "test description",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// make args body into io.Reader
			var body []byte
			body, err = json.Marshal(tt.args.body)
			Nil(t, err)

			requestReader := bytes.NewReader(body)
			req := httptest.NewRequest(http.MethodPost, "/products", requestReader)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err = httpAdapter.CreateProduct(c)
			Nil(t, err)

			if tt.wantErr {
				Equal(t, tt.err.Status(), rec.Code)

				var respBody errors.CustomError
				err = json.NewDecoder(rec.Result().Body).Decode(&respBody)
				Nil(t, err)

				EqualValues(t, tt.err.Print(), respBody.Print())
			} else {
				Equal(t, http.StatusOK, rec.Code)
			}
		})
	}
}

func TestAdapter_UpdateProduct(t *testing.T) {
	e := echo.New()

	dbAdapter, err := db_adapter.NewAdapterForTest(t)
	Nil(t, err)

	appAdapter := app_adapter.NewApplication(config.AppConfig{}, dbAdapter)
	httpAdapter := http_adapter.NewAdapter(config.AppConfig{}, appAdapter)

	testProduct := types.ProductBody{
		Name:        "item 1",
		Description: "test description",
		Price:       10.5,
	}

	// create a product
	product, err := appAdapter.CreateProduct(testProduct)
	Nil(t, err)
	NotZero(t, product.ID)

	type args struct {
		productID uint
		body      types.ProductBody
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     errors.Error
	}{
		{
			name: "update product with incorrect id",
			args: args{
				productID: 0,
				body: types.ProductBody{
					Name:        "item 1",
					Description: "test description",
					Price:       100,
				},
			},
			wantErr: true,
			err:     errors.NoEntityError("product"),
		},
		{
			name: "update product with no name",
			args: args{
				productID: product.ID,
				body: types.ProductBody{
					Description: "no name provided",
					Price:       100,
				},
			},
			wantErr: true,
			err:     errors.ValidationError("You must provide the name"),
		},
		{
			name: "update product with negative price",
			args: args{
				productID: product.ID,
				body: types.ProductBody{
					Name:        "item 1",
					Description: "test description",
					Price:       -10,
				},
			},
			wantErr: true,
			err:     errors.ValidationError("You must provide a valid price greater than or equal to 0"),
		},
		{
			name: "update product successfully",
			args: args{
				productID: product.ID,
				body: types.ProductBody{
					Name:        "item 1.2",
					Description: "test description updated",
					Price:       11.5,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// make args body into io.Reader
			var body []byte
			body, err = json.Marshal(tt.args.body)
			Nil(t, err)

			requestReader := bytes.NewReader(body)
			req := httptest.NewRequest(http.MethodPut, "/products/:product", requestReader)

			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("product")
			c.SetParamValues(fmt.Sprintf("%d", tt.args.productID))

			err = httpAdapter.UpdateProduct(c)
			Nil(t, err)

			if tt.wantErr {
				Equal(t, tt.err.Status(), rec.Code)

				var respBody errors.CustomError
				err = json.NewDecoder(rec.Result().Body).Decode(&respBody)
				Nil(t, err)

				EqualValues(t, tt.err.Print(), respBody.Print())
			} else {
				Equal(t, http.StatusOK, rec.Code)

				updatedProduct, cErr := appAdapter.GetProductByID(tt.args.productID)
				Nil(t, cErr)

				Equal(t, tt.args.body.Name, updatedProduct.Name)
				Equal(t, tt.args.body.Description, updatedProduct.Description)
				Equal(t, tt.args.body.Price, updatedProduct.Price)
			}
		})
	}
}

func TestAdapter_ListProduct(t *testing.T) {
	e := echo.New()

	dbAdapter, err := db_adapter.NewAdapterForTest(t)
	Nil(t, err)

	appAdapter := app_adapter.NewApplication(config.AppConfig{}, dbAdapter)
	httpAdapter := http_adapter.NewAdapter(config.AppConfig{PaginationSize: 10}, appAdapter)

	type args struct {
		size int
		body *types.ProductBody
	}
	tests := []struct {
		name  string
		args  args
		want  int
		total int
	}{
		{
			name:  "list products with no products",
			want:  0,
			total: 0,
		},
		{
			name: "add 1st product",
			args: args{
				size: 5,
				body: &types.ProductBody{
					Name:        "item 1",
					Description: "test description",
					Price:       100,
				},
			},
			want:  1,
			total: 1,
		},
		{
			name: "add 2nd product",
			args: args{
				size: 10,
				body: &types.ProductBody{
					Name:        "item 2",
					Description: "test description",
					Price:       100,
				},
			},
			want:  2,
			total: 2,
		},
		{
			name: "add 3rd product",
			args: args{
				size: 10,
				body: &types.ProductBody{
					Name:        "item 3",
					Description: "test description",
					Price:       100,
				},
			},
			want:  3,
			total: 3,
		},
		{
			name: "list with page size 2",
			args: args{
				size: 2,
			},
			want:  2,
			total: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if tt.args.body != nil {
				_, err = appAdapter.CreateProduct(*tt.args.body)
				Nil(t, err)
			}

			req := httptest.NewRequest(http.MethodGet, "/products", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if tt.args.size > 0 {
				qParams := c.QueryParams()
				qParams.Set("size", fmt.Sprintf("%d", tt.args.size))
			}

			err = httpAdapter.ListProduct(c)
			Nil(t, err)

			Equal(t, http.StatusOK, rec.Code)

			var respBody types.Page[types.ProductResponse]
			err = json.NewDecoder(rec.Result().Body).Decode(&respBody)
			Nil(t, err)

			Equal(t, tt.want, len(respBody.Data))
			Equal(t, tt.total, respBody.Total)
		})
	}
}

func TestAdapter_GetProduct(t *testing.T) {
	e := echo.New()

	dbAdapter, err := db_adapter.NewAdapterForTest(t)
	Nil(t, err)

	appAdapter := app_adapter.NewApplication(config.AppConfig{}, dbAdapter)
	httpAdapter := http_adapter.NewAdapter(config.AppConfig{}, appAdapter)

	testProduct := types.ProductBody{
		Name:        "item 1",
		Description: "test description",
		Price:       10.5,
	}

	// create a product
	product, err := appAdapter.CreateProduct(testProduct)
	Nil(t, err)
	NotZero(t, product.ID)

	type args struct {
		productID uint
	}
	tests := []struct {
		name    string
		args    args
		resp    types.ProductResponse
		wantErr bool
		err     errors.Error
	}{
		{
			name:    "get product with incorrect id",
			args:    args{productID: 0},
			wantErr: true,
			err:     errors.NoEntityError("product"),
		},
		{
			name: "get product successfully",
			args: args{productID: product.ID},
			resp: product,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodGet, "/products/:product", nil)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("product")
			c.SetParamValues(fmt.Sprintf("%d", tt.args.productID))

			err = httpAdapter.GetProduct(c)
			Nil(t, err)

			if tt.wantErr {
				Equal(t, tt.err.Status(), rec.Code)

				var respBody errors.CustomError
				err = json.NewDecoder(rec.Result().Body).Decode(&respBody)
				Nil(t, err)

				EqualValues(t, tt.err.Print(), respBody.Print())
			} else {
				Equal(t, http.StatusOK, rec.Code)

				var respBody types.ProductResponse
				err = json.NewDecoder(rec.Result().Body).Decode(&respBody)
				Nil(t, err)

				Equal(t, tt.resp.ID, respBody.ID)
				Equal(t, tt.resp.Name, respBody.Name)
				Equal(t, tt.resp.Description, respBody.Description)
				Equal(t, tt.resp.Price, respBody.Price)

			}
		})
	}
}

func TestAdapter_DeleteProduct(t *testing.T) {
	e := echo.New()

	dbAdapter, err := db_adapter.NewAdapterForTest(t)
	Nil(t, err)

	appAdapter := app_adapter.NewApplication(config.AppConfig{}, dbAdapter)
	httpAdapter := http_adapter.NewAdapter(config.AppConfig{}, appAdapter)

	testProduct := types.ProductBody{
		Name:        "item 1",
		Description: "test description",
		Price:       10.5,
	}

	// create a product
	product, err := appAdapter.CreateProduct(testProduct)
	Nil(t, err)
	NotZero(t, product.ID)

	type args struct {
		productID uint
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
		err     errors.Error
	}{
		{
			name:    "delete product with incorrect id",
			args:    args{productID: 0},
			wantErr: true,
			err:     errors.NoEntityError("product"),
		},
		{
			name:    "delete product successfully",
			args:    args{productID: product.ID},
			wantErr: false,
		},
		{
			name:    "delete product second time",
			args:    args{productID: product.ID},
			wantErr: true,
			err:     errors.NoEntityError("product"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodGet, "/products/:product", nil)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetParamNames("product")
			c.SetParamValues(fmt.Sprintf("%d", tt.args.productID))

			err = httpAdapter.DeleteProduct(c)
			Nil(t, err)

			if tt.wantErr {
				Equal(t, tt.err.Status(), rec.Code)

				var respBody errors.CustomError
				err = json.NewDecoder(rec.Result().Body).Decode(&respBody)
				Nil(t, err)

				EqualValues(t, tt.err.Print(), respBody.Print())
			} else {
				Equal(t, http.StatusOK, rec.Code)
			}
		})
	}
}
