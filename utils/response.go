package utils

import (
	"encoding/json"
	"net/http"
)

type U struct{}

func (u U) HttpSendJSON(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "Application/json")
	json.NewEncoder(w).Encode(body)
}
