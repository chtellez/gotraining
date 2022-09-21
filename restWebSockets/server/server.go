package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Struct with parameters needed to start server
type Config struct {
	Port 		string
	JWTSecret 	string
	DatabaseUrl string
}

//Interface needed to start a new server
type Server interface {
	Config() *Config
}

//Broker that handle the conexion to the server
type Broker struct {
	config *Config
	router *mux.Router
}

//Broker method for setting up a new Broker
func (b *Broker) Config() *Config {
	return b.config
}

//func that validates the config needed to create a new Broker object
func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	//Port, JWTSecret and DatabaseUrl must be specified
	if config.Port == "" {
		return nil, errors.New("port is required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("secret is required")
	}
	if config.DatabaseUrl == "" {
		return nil, errors.New("database url is required")
	}

	//Returns a new broker
	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
	}
	return broker, nil
}

//broker method that starts a new server
func (b *Broker) Start(binder func (s Server, r *mux.Router)){
	b.router = mux.NewRouter()
	binder(b, b.router)
	log.Println("Startting server on port", b.Config().Port)
	if err := http.ListenAndServe(b.config.Port, b.router); err != nil {
		log.Fatal("ListenAndServe error: ", err)
	}
}