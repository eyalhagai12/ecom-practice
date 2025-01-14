package inventory

import (
	"ecom/db"
	"ecom/server"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type GetInventoryByProductUUIDRequest struct {
	ProductUUID string `param:"productUuid"`
}

func GetProductInventories(c echo.Context, env server.Env, request GetInventoryByProductUUIDRequest) ([]db.Inventory, error) {
	inventory, err := env.Queries.GetProductInventories(c.Request().Context(), uuid.MustParse(request.ProductUUID))
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("failed to get inventories of product with uuid %s: %s", request.ProductUUID, err.Error()))
	}

	return inventory, nil
}

type GetInventoryByUUIDRequest struct {
	InventoryUUID string `param:"inventoryUuid"`
}

func GetInventoryByUUID(c echo.Context, env server.Env, request GetInventoryByUUIDRequest) (db.Inventory, error) {
	inventory, err := env.Queries.GetInventoryByID(c.Request().Context(), uuid.MustParse(request.InventoryUUID))
	if err != nil {
		return db.Inventory{}, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("failed to get inventory with uuid %s: %s", request.InventoryUUID, err.Error()))
	}

	return inventory, nil
}

type LocationRequestData struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type NewInventoryReqeust struct {
	ProductID uuid.UUID           `json:"productUuid"`
	Quantity  int32               `json:"quantity"`
	Location  LocationRequestData `json:"location"`
}

func NewInventory(c echo.Context, env server.Env, request NewInventoryReqeust) (db.Inventory, error) {
	// todo: Use transaction here!!!
	tx, err := env.DB.Begin(c.Request().Context())
	if err != nil {
		return db.Inventory{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to start transaction: %s", err.Error()))
	}

	txQueries := env.Queries.WithTx(tx)
	defer tx.Rollback(c.Request().Context())

	location, err := txQueries.CreateLocation(c.Request().Context(), request.Location.Name, request.Location.Address)
	if err != nil {
		return db.Inventory{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to create location: %s", err.Error()))
	}

	inventory, err := txQueries.CreateInventory(c.Request().Context(), uuid.New(), request.ProductID, request.Quantity, location.ID)
	if err != nil {
		return db.Inventory{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to create inventory: %s", err.Error()))
	}

	if err := txQueries.IncreaseProductQuantity(c.Request().Context(), request.ProductID, &request.Quantity); err != nil {
		return db.Inventory{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to increase product quantity: %s", err.Error()))
	}

	tx.Commit(c.Request().Context())
	return inventory, nil
}

type UpdateInventoryRequest struct {
	InventoryUUID string `param:"inventoryUuid"`
	Quantity      int32  `json:"quantity"`
}

func UpdateInventory(c echo.Context, env server.Env, request UpdateInventoryRequest) (db.Inventory, error) {
	inventory, err := env.Queries.UpdateInventory(c.Request().Context(), uuid.MustParse(request.InventoryUUID), request.Quantity)
	if err != nil {
		return db.Inventory{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to update inventory: %s", err.Error()))
	}

	return inventory, nil
}

type DeleteInventoryRequest struct {
	InventoryUUID string `param:"inventoryUuid"`
}

func DeleteInventory(c echo.Context, env server.Env, request DeleteInventoryRequest) (db.Inventory, error) {
	inventory, err := env.Queries.DeleteInventory(c.Request().Context(), uuid.MustParse(request.InventoryUUID))
	if err != nil {
		return db.Inventory{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to delete inventory: %s", err.Error()))
	}

	return inventory, nil
}

func RegisterHandlers(app *echo.Group, env server.Env) {
	app.GET("/products/:productUuid/inventories", server.HandlerFromFunc(env, GetProductInventories, http.StatusOK))
	app.GET("/inventories/:inventoryUuid", server.HandlerFromFunc(env, GetInventoryByUUID, http.StatusOK))
	app.POST("/inventories", server.HandlerFromFunc(env, NewInventory, http.StatusCreated))
	app.PUT("/inventories/:inventoryUuid", server.HandlerFromFunc(env, UpdateInventory, http.StatusOK))
	app.DELETE("/inventories/:inventoryUuid", server.HandlerFromFunc(env, DeleteInventory, http.StatusOK))
}
