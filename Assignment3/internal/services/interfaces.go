package services

import (
	"Project/internal/models"
	"context"
)

type Customer interface {
	GetCustomer(ctx context.Context, customerId int) ([]models.Customer, error)
	PostNewCustomer(ctx context.Context, o models.Customer) (*models.Customer, error)
}
