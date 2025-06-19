package dtos

import "github.com/google/uuid"

type GetCustomerRequestDTO struct {
	UserID uuid.UUID `json:"userId" validate:"required"`
}

type GetCustomerResponseDTO struct {
	UserID    uuid.UUID `json:"userId"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Address   string    `json:"address"`
	Username  string    `json:"username"`
}
