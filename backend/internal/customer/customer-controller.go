package customer

import (
	"example-backend/internal/customer/dtos"
	"example-backend/internal/middleware"
	"example-backend/internal/utils"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type CustomerController struct {
	customerService *CustomerService
}

func NewCustomerController(customerService *CustomerService) *CustomerController {
	return &CustomerController{customerService}
}

func (c *CustomerController) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/customer", func(r chi.Router) {
		r.With(middleware.ValidateBody[dtos.GetCustomerRequestDTO]()).Get("/", c.getCustomer)

		r.With(middleware.ValidateBody[dtos.CreateCustomerRequestDTO]()).Post("/", c.createCustomer)
	})

	return r
}

func (c *CustomerController) getCustomer(w http.ResponseWriter, r *http.Request) {
	body, ok := middleware.BodyFromContext[dtos.GetCustomerRequestDTO](r.Context(), w)
	if !ok {
		return
	}

	user, err := c.customerService.Get(r.Context(), body.UserID)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to get requested customer", http.StatusInternalServerError)
	}

	utils.EncodeDTO(w, user)
}

func (c *CustomerController) createCustomer(w http.ResponseWriter, r *http.Request) {
	body, ok := middleware.BodyFromContext[dtos.CreateCustomerRequestDTO](r.Context(), w)
	if !ok {
		return
	}

	id, err := c.customerService.Create(r.Context(), body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to create customer", http.StatusInternalServerError)
	}

	response := dtos.CreateCustomerResponseDTO{UserID: id}

	utils.EncodeDTO(w, response)
}
