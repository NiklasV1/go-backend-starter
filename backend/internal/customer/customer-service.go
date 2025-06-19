package customer

import (
	"context"
	"example-backend/internal/repository"
)

type CustomerService struct {
	repository *repository.Queries
}

func NewCustomerService(repository *repository.Queries) *CustomerService {
	return &CustomerService{repository}
}

func (c *CustomerService) Create(
	ctx context.Context,
	firstName,
	lastName,
	address,
	username string,
	password []byte,
) error {
	err := c.repository.AddCustomer(ctx, repository.AddCustomerParams{
		FirstName: firstName,
		LastName:  lastName,
		Address:   address,
		Username:  username,
		Password:  password,
	})

	return err
}
