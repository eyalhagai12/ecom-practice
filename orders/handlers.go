package orders

import (
	"ecom/db"
	"ecom/server"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type GetOrderByUUIDRequest struct {
	OrderUUID string `param:"orderUuid"`
}

type OrderAction struct {
	Name   string `json:"name"`
	URI    string `json:"uri"`
	Method string `json:"method"`
}

type OrderResponse struct {
	db.Order
	Items   []db.OrderItem `json:"items"`
	Actions []OrderAction  `json:"_actions,omitempty"`
}

func GetOrderByUUID(c echo.Context, env server.Env, request GetOrderByUUIDRequest) (OrderResponse, error) {
	orderItemPairs, err := env.Queries.GetOrderByUUID(c.Request().Context(), uuid.MustParse(request.OrderUUID))
	if err != nil {
		return OrderResponse{}, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("failed to get order with uuid %s: %s", request.OrderUUID, err.Error()))
	}

	items := []db.OrderItem{}
	for _, item := range orderItemPairs {
		items = append(items, item.OrderItem)
	}

	if len(orderItemPairs) == 0 {
		return OrderResponse{}, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("order with uuid %s not found", request.OrderUUID))
	}

	order := orderItemPairs[0].Order

	var actions []OrderAction

	if order.Status == db.OrderStatusPending {
		actions = append(actions, OrderAction{
			Name:   "cancel",
			URI:    fmt.Sprintf("/orders/%s", order.ID),
			Method: http.MethodDelete,
		})
	}

	return OrderResponse{
		Order:   order,
		Items:   items,
		Actions: actions,
	}, nil
}

func GetOrders(c echo.Context, env server.Env) ([]db.Order, error) {
	orders, err := env.Queries.GetOrders(c.Request().Context())
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to get orders: %s", err.Error()))
	}

	return orders, nil
}

type NewOrderRequest struct {
	Items []struct {
		ProductID uuid.UUID `json:"productUuid"`
		Quantity  int32     `json:"quantity"`
		Price     float64   `json:"price"`
	} `json:"items" validate:"required,min=1"`
}

func NewOrder(c echo.Context, env server.Env, request NewOrderRequest) (db.Order, error) {
	tx, err := env.DB.Begin(c.Request().Context())
	if err != nil {
		return db.Order{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to start transaction: %s", err.Error()))
	}

	txQueries := env.Queries.WithTx(tx)
	defer tx.Rollback(c.Request().Context())

	var totalPrice float64
	for _, item := range request.Items {
		totalPrice += item.Price * float64(item.Quantity)
	}

	order, err := txQueries.CreateOrder(c.Request().Context(), uuid.New(), totalPrice)
	if err != nil {
		return db.Order{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to create order: %s", err.Error()))
	}

	for _, item := range request.Items {
		_, err := txQueries.CreateOrderItem(c.Request().Context(), order.ID, item.ProductID, item.Quantity)
		if err != nil {
			return db.Order{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to create order item: %s", err.Error()))
		}
	}

	if err := tx.Commit(c.Request().Context()); err != nil {
		return db.Order{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to commit transaction: %s", err.Error()))
	}

	return order, nil
}

func CancelOrder(c echo.Context, env server.Env, request GetOrderByUUIDRequest) (db.Order, error) {
	tx, err := env.DB.Begin(c.Request().Context())
	if err != nil {
		return db.Order{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to start transaction: %s", err.Error()))
	}

	txQueries := env.Queries.WithTx(tx)
	defer tx.Rollback(c.Request().Context())

	orderItemPairs, err := txQueries.GetOrderByUUID(c.Request().Context(), uuid.MustParse(request.OrderUUID))
	if err != nil {
		return db.Order{}, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("failed to get order with uuid %s: %s", request.OrderUUID, err.Error()))
	}

	if len(orderItemPairs) == 0 {
		return db.Order{}, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("order with uuid %s not found", request.OrderUUID))
	}

	order := orderItemPairs[0].Order
	if order.Status == db.OrderStatusCancelled {
		return db.Order{}, echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("order with uuid %s is already cancelled", request.OrderUUID))
	}

	if order.Status == db.OrderStatusDelivered {
		return db.Order{}, echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("order with uuid %s is already delivered", request.OrderUUID))
	}

	if order.Status == db.OrderStatusShipped {
		return db.Order{}, echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("order with uuid %s is already shipped", request.OrderUUID))
	}

	order, err = txQueries.UpdateOrderStatus(c.Request().Context(), order.ID, db.OrderStatusCancelled)
	if err != nil {
		return db.Order{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to cancel order: %s", err.Error()))
	}

	if err := tx.Commit(c.Request().Context()); err != nil {
		return db.Order{}, echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("failed to commit transaction: %s", err.Error()))
	}

	return order, nil
}

func RegisterHandlers(app *echo.Group, env server.Env) {
	app.GET("/orders/:orderUuid", server.HandlerFromFunc(env, GetOrderByUUID, http.StatusOK))
	app.GET("/orders", server.HandlerNoRequestFromFunc(env, GetOrders, http.StatusOK))
	app.POST("/orders", server.HandlerFromFunc(env, NewOrder, http.StatusCreated))
	app.DELETE("/orders/:orderUuid", server.HandlerFromFunc(env, CancelOrder, http.StatusOK))
}
