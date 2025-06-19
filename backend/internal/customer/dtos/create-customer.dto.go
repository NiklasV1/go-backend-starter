package dtos

import "github.com/google/uuid"

type CreateCustomerRequestDTO struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Address   string `json:"address" validate:"required"`
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type CreateCustomerResponseDTO struct {
	UserID uuid.UUID `json:"userId"`
}
