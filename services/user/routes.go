package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	
}

func NewHandler() *Handler {
	return &Handler{}
}

// router.HandleFunc("/path/to/be/reg", handlerFunction(http.ResponseWriter, *http.Request))

func (h *Handler) RegisterRoutes(router *mux.Router) {
	/* 
	router.HandleFunc( "/login", func(w http.ResponseWriter, r *http.Request) {} ).Methods("POST")

	not a very clear way to write
	*/

	router.HandleFunc( "/login", h.handleLogin ).Methods("POST")
	router.HandleFunc( "/register", h.handleRegister ).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

}

