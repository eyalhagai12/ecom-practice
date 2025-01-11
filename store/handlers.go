package store

import (
	"ecom/db"
	"ecom/server"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type GetStoreByUUIDRequest struct {
	UUID uuid.UUID `param:"storeUuid"`
}

func GetStore(c echo.Context, env server.Env, request GetStoreByUUIDRequest) (db.Store, error) {
	store, err := env.Queries.GetStoreByUUID(c.Request().Context(), request.UUID)
	if err != nil {
		return db.Store{}, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("failed to get store with uuid %s: %s", request.UUID.String(), err.Error()))
	}

	return store, nil
}

type NewStoreRequest struct {
	Name string `json:"name"`
}

func NewStore(c echo.Context, env server.Env, request NewStoreRequest) (db.Store, error) {
	store, err := env.Queries.InsertNewStore(c.Request().Context(), uuid.New(), request.Name)
	if err != nil {
		return db.Store{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to create store: %s", err.Error()))
	}

	return store, nil
}

type UpdateStoreRequest struct {
	UUID uuid.UUID `param:"storeUuid"`
	Name string    `json:"name"`
}

func UpdateStore(c echo.Context, env server.Env, request UpdateStoreRequest) (db.Store, error) {
	store, err := env.Queries.UpdateStore(c.Request().Context(), request.UUID, request.Name)
	if err != nil {
		return db.Store{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to update store: %s", err.Error()))
	}

	return store, nil
}

func DeleteStore(c echo.Context, env server.Env, request GetStoreByUUIDRequest) (db.Store, error) {
	store, err := env.Queries.DeleteStore(c.Request().Context(), request.UUID)
	if err != nil {
		return db.Store{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to delete store: %s", err.Error()))
	}

	return store, nil
}

func RegisterHandlers(app *echo.Group, env server.Env) {
	app.GET("/stores/:storeUuid", server.HandlerFromFunc(env, GetStore, http.StatusOK))
	app.POST("/stores", server.HandlerFromFunc(env, NewStore, http.StatusCreated))
	app.PUT("/stores/:storeUuid", server.HandlerFromFunc(env, UpdateStore, http.StatusOK))
	app.DELETE("/stores/:storeUuid", server.HandlerFromFunc(env, DeleteStore, http.StatusOK))
}
