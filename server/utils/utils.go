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

func WriteResponse(w http.ResponseWriter, status uint, payload any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(status))
	return json.NewEncoder(w).Encode(payload)
}