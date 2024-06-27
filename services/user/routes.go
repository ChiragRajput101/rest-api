package user

import (
	// "encoding/json"
	"net/http"

	"github.com/ChiragRajput101/rest-api/types"
	"github.com/ChiragRajput101/rest-api/utils"
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
	// takes in the JSON payload

	/*
	common functionality: (hence modularising the code, also follows DRY)
	
	var payload types.RegisterUserPayload
	err := json.NewDecoder(r.Body).Decode(payload)
	*/

	var payload types.RegisterUserPayload

	// Parse -> Unmarshalling
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// check if the user already exixts
	// if not then create a new account
}

