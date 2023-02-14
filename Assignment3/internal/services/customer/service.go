package customer

import (
	"Project/internal/filters"
	"Project/internal/models"
	"Project/internal/stores"
	"context"
	"errors"
	"fmt"
)

type service struct {
	store stores.CustomerStorer
}

func New(store stores.CustomerStorer) *service {
	return &service{store: store}
}

type InvalidParam struct {
	Param string
}

func (i InvalidParam) Error() string {
	return fmt.Sprintf("param '%s' is invalid", i.Param)
}

func (s *service) GetCustomer(ctx context.Context, customerId int) ([]models.Customer, error) {
	if customerId <= 0 {
		return nil, InvalidParam{Param: "customerId"}
	}
	fmt.Println(customerId)

	res, err := s.store.GetCustomers(ctx, filters.Customer{Id: customerId})

	if err != nil {
		return nil, errors.New("unable to get from database")
	}
	return res, nil
}

func (s *service) PostNewCustomer(ctx context.Context, o models.Customer) (*models.Customer, error) {

	fmt.Println("Reached service layer ")
	res, err := s.store.PostNewCustomer(ctx, o)

	if err != nil {
		fmt.Println("Error in the store ", err)
		return nil, err
	}

	return res, nil
}
