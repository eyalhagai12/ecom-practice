package product

import (
	"ecom/db"
	"ecom/server"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type GetStoreProductsRequest struct {
	StoreUUID string `param:"storeUuid"`
}

type GetStoreProductsResponse struct {
	Products []db.Product `json:"products"`
}

func GetStoreProducts(c echo.Context, env server.Env, request GetStoreProductsRequest) (GetStoreProductsResponse, error) {
	products, err := env.Queries.GetStoreProducts(c.Request().Context(), uuid.MustParse(request.StoreUUID))
	if err != nil {
		return GetStoreProductsResponse{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to get products for store with uuid %s: %s", request.StoreUUID, err.Error()))
	}

	return GetStoreProductsResponse{Products: products}, nil
}

type GetProductByUUIDRequest struct {
	ProductUUID string `param:"productUuid"`
}

func GetProductByUUID(c echo.Context, env server.Env, request GetProductByUUIDRequest) (db.Product, error) {
	product, err := env.Queries.GetProductByID(c.Request().Context(), uuid.MustParse(request.ProductUUID))
	if err != nil {
		return db.Product{}, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("failed to get product with uuid %s: %s", request.ProductUUID, err.Error()))
	}

	return product, nil
}

type NewProductRequest struct {
	Name     string    `json:"name"`
	Price    float64   `json:"price"`
	Quantity int32     `json:"quantity"`
	StoreID  uuid.UUID `json:"storeUuid"`
}

func NewProduct(c echo.Context, env server.Env, request NewProductRequest) (db.Product, error) {
	product, err := env.Queries.CreateProduct(c.Request().Context(), uuid.New(), request.Name, request.Price, request.Quantity, request.StoreID)
	if err != nil {
		return db.Product{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to create product: %s", err.Error()))
	}

	return product, nil
}

type UpdateProductRequest struct {
	ID       string  `param:"productUuid"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int32   `json:"quantity"`
}

func UpdateProduct(c echo.Context, env server.Env, request UpdateProductRequest) (db.Product, error) {
	product, err := env.Queries.UpdateProduct(c.Request().Context(), uuid.MustParse(request.ID), request.Name, request.Price, request.Quantity)
	if err != nil {
		return db.Product{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to update product: %s", err.Error()))
	}

	return product, nil
}

func DeleteProduct(c echo.Context, env server.Env, request GetProductByUUIDRequest) (db.Product, error) {
	product, err := env.Queries.DeleteProduct(c.Request().Context(), uuid.MustParse(request.ProductUUID))
	if err != nil {
		return db.Product{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to delete product: %s", err.Error()))
	}

	return product, nil
}

func RegisterHandlers(app *echo.Group, env server.Env) {
	app.GET("/stores/:storeUuid/products", server.HandlerFromFunc(env, GetStoreProducts, http.StatusOK))
	app.GET("/products/:productUuid", server.HandlerFromFunc(env, GetProductByUUID, http.StatusOK))
	app.POST("/products", server.HandlerFromFunc(env, NewProduct, http.StatusCreated))
	app.PUT("/products/:productUuid", server.HandlerFromFunc(env, UpdateProduct, http.StatusOK))
	app.DELETE("/products/:productUuid", server.HandlerFromFunc(env, DeleteProduct, http.StatusOK))
}
