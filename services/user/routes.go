package user

import (
	"fmt"
	"net/http"

	"github.com/ChiragRajput101/rest-api/config"
	"github.com/ChiragRajput101/rest-api/services/auth"
	"github.com/ChiragRajput101/rest-api/types"
	"github.com/ChiragRajput101/rest-api/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{
		store: store,
	}
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
	var user types.LoginUserPayload

	if err := utils.ParseJSON(r, &user); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(user); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload: %v", errors))
		return
	}

	u, err := h.store.GetUserByEmail(user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("not found, invalid email"))
		return
	}

	if !auth.ComparePasswords(u.Password, []byte(user.Password)) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid email or password"))
		return
	}

	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.ID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"token": token})

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
	// Unmarshalled form in stored in v 
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// validate the payload

	if err := utils.Validate.Struct(payload); err != nil {
		errs := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errs))
		return
	}

	// check if the user already exixts

	_, err := h.store.GetUserByEmail(payload.Email)

	// if err == nil, then we got a user with the specified email
	if err == nil {
		utils.WriteError( w, http.StatusBadRequest, fmt.Errorf("user already exists") )
		return
	}

	// create the user if non-existant

	hashedPassword, _ := auth.HashPassword(payload.Password)
	
	err1 := h.store.CreateUser(types.User{
		FirstName: payload.FirstName,
		LastName: payload.LastName,
		Email: payload.Email,
		Password: hashedPassword,
	})

	if err1 != nil {
		utils.WriteError(w, http.StatusInternalServerError, err1)
		return
	}

}

