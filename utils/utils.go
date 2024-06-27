package utils

import (
	"encoding/json"
	"errors"
	"net/http"
)

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