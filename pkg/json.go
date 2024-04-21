package pkg

import (
	"encoding/json"
	"net/http"
)

func GetJsonBody(r *http.Request, v any) error {
	err := json.NewDecoder(r.Body).Decode(v)
	defer r.Body.Close()
	return err
}

func WriteJsonError(w http.ResponseWriter, v any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	m := map[string]interface{}{
		"message": v,
	}

	json.NewEncoder(w).Encode(m)
}

func WriteJson(w http.ResponseWriter, code int, v any) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(v)
}
