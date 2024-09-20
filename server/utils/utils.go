package utils

import (
	"fmt"
	"os"
)

func GetEnv(key string) (string, error) {
	val, ok := os.LookupEnv(key)

	if ok {
		return val, nil
	}

	return "", fmt.Errorf("variavel ambiente %s nao encontrada", key)
}