package utils

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func ParseJSON(r *http.Request, payload any) error {
	if r.Body == nil {
		return errors.New("no request body found")
	} 

	// dec := json.NewDecoder(input stream to read using io.Reader)
	// de.Decode(v any) -> Unmarshals JSON-encoded data and stores the result in the value pointed to by v. 
	// If v is nil or not a pointer, Unmarshal returns an InvalidUnmarshalError.
	return json.NewDecoder(r.Body).Decode(payload)
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json") // declaring => the data below is of json format
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, err error) {
	WriteJSON(w, status, map[string]string{"error":err.Error()})
}

func GetTokenFromRequest(r *http.Request) string {
	tokenAuth := r.Header.Get("Authorization")
	tokenQuery := r.URL.Query().Get("token")
	
	if tokenAuth != "" {
		return tokenAuth
	}

	if tokenQuery != "" {
		return tokenQuery
	}

	return ""
}