package shipping

import (
	"context"
	"ecom/db"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ShippingProcess struct {
	Shipping      db.Shipping
	StatusUpdates <-chan db.ShippingStatus
	Errs          <-chan error
}

func CreateShippingProcess(ctx context.Context, queries *db.Queries, txQueries *db.Queries, orderID uuid.UUID) (ShippingProcess, error) {
	shipping, err := txQueries.CreateShippingProcess(ctx, uuid.New(), orderID)
	if err != nil {
		return ShippingProcess{}, err
	}

	statusUpdates := make(chan db.ShippingStatus)
	errs := make(chan error)

	go func(ctx context.Context) {
		<-time.After(30 * time.Second)
		fmt.Println("shipping is delivered")

		if _, err := queries.UpdateShippingStatus(ctx, shipping.ID, db.ShippingStatusDelivered); err != nil {
			errs <- err
			return
		}

		fmt.Println("shipping status updated to delivered")

		statusUpdates <- db.ShippingStatusDelivered
	}(context.Background())

	return ShippingProcess{
		Shipping:      shipping,
		StatusUpdates: statusUpdates,
		Errs:          errs,
	}, nil
}
