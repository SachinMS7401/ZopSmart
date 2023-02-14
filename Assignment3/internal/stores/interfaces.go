package stores

import (
	"Project/internal/filters"
	"Project/internal/models"
	"context"
)

type CustomerStorer interface {
	GetCustomers(ctx context.Context, f filters.Customer) ([]models.Customer, error)
	PostNewCustomer(ctx context.Context, o models.Customer) (*models.Customer, error)
}
