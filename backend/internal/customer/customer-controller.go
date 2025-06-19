package customer

import (
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
		r.Get("/", c.getCustomer)
		r.Post("/", c.createCustomer)
	})

	return r
}

func (c *CustomerController) getCustomer(w http.ResponseWriter, r *http.Request) {

}

func (c *CustomerController) createCustomer(w http.ResponseWriter, r *http.Request) {
	err := c.customerService.Create(r.Context(), "Test", "Person", "Street 3", "testman", []byte("password"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to create customer", http.StatusInternalServerError)
	}
}
