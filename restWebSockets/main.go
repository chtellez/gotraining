package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/chtellez/gotraining/restWebSockets/handlers"
	"github.com/chtellez/gotraining/restWebSockets/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JwtSecret := os.Getenv("JWT_SECRET")
	DatabaseUrl := os.Getenv("DATABASE_URL")

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:        PORT,
		JWTSecret:   JwtSecret,
		DatabaseUrl: DatabaseUrl,
	})
	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)
}

func BindRoutes(s server.Server, r * mux.Router){
	r.HandleFunc("/", handlers.HomeHandler(s)).Methods(http.MethodGet)
}