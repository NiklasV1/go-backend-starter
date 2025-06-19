package main

import (
	"context"
	"example-backend/internal/customer"
	"example-backend/internal/repository"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Load env file
	err := godotenv.Load("./backend/config.env")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Get DB connection
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Println("Unable to establish database connection")
		panic(err)
	}
	defer conn.Close(ctx)

	// Create repository instance
	repository := repository.New(conn)

	// Setup router
	r := chi.NewRouter()

	// Setup middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(15 * time.Second))

	// Add health check endpoint
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running."))
	})

	// Create services
	customerService := customer.NewCustomerService(repository)

	// Setup controllers
	customerController := customer.NewCustomerController(customerService)
	r.Mount("/", customerController.Routes())

	// Start application
	http.ListenAndServe(":"+os.Getenv("BACKEND_PORT"), r)
}
