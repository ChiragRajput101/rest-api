package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ChiragRajput101/rest-api/services/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db *sql.DB
}

func InitServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// User Service Init
	userService := user.NewHandler()
	userService.RegisterRoutes(subrouter) // passing the subrouter to match the prefix: /api/v1



	log.Println("listening on", s.addr)

	// mux.Router implements the http.Handler interface hence ListenAndServe(addr, http.Handler) works with mux.Router as the param passed
	return http.ListenAndServe(s.addr, router) 
}