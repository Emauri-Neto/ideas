package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func GetEnv(key string) (string, error) {
	val, ok := os.LookupEnv(key)

	if ok {
		return val, nil
	}

	return "", fmt.Errorf("variavel ambiente %s nao encontrada", key)
}

func WriteResponse(w http.ResponseWriter, status uint, message string) error {
	var r struct {
		Status  uint `json:"status"`
		Message string `json:"message"`
	}

	r.Status = status
	r.Message = message

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(r)
}