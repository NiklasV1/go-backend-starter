package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello world")

	// Load env file
	err := godotenv.Load("./backend/config.env")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Get DB connection
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close(ctx)

}
