package customer

import (
	"context"
	"example-backend/internal/customer/dtos"
	"example-backend/internal/repository"
	"example-backend/internal/utils"

	"github.com/google/uuid"
)

type CustomerService struct {
	repository *repository.Queries
}

func NewCustomerService(repository *repository.Queries) *CustomerService {
	return &CustomerService{repository}
}

func (c *CustomerService) Create(
	ctx context.Context,
	customerDTO dtos.CreateCustomerRequestDTO,
) (uuid.UUID, error) {

	passwordHash, err := utils.HashPassword(customerDTO.Password)
	if err != nil {
		return uuid.UUID{}, err
	}

	id, err := c.repository.AddCustomer(ctx, repository.AddCustomerParams{
		FirstName: customerDTO.FirstName,
		LastName:  customerDTO.LastName,
		Address:   customerDTO.Address,
		Username:  customerDTO.Username,
		Password:  passwordHash,
	})

	return id, err
}

func (c *CustomerService) Get(ctx context.Context, id uuid.UUID) (dtos.GetCustomerResponseDTO, error) {
	user, err := c.repository.GetCustomer(ctx, id)
	if err != nil {
		return dtos.GetCustomerResponseDTO{}, err
	}

	response := dtos.GetCustomerResponseDTO{
		UserID:    user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Address:   user.Address,
		Username:  user.Username,
	}

	return response, err
}
