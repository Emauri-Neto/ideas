package auth

import (
	"fmt"
	"ideas/utils"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HashValue(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CompareValue(hash string, value []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), value)

	return err == nil
}

func SignToken(secret []byte, userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(15 * time.Minute).Unix(),
	})

	t, err := token.SignedString(secret)

	if err != nil {
		return "", err
	}

	return t, nil
}

func GetToken(tokenStr string) (*jwt.Token, error) {
	secret, _secret := utils.GetEnv("JWT_SECRET")

	if _secret != nil {
		return nil, fmt.Errorf("%s", _secret)
	}

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return token, nil
}

func Authorize(tokenStr string) (bool, error) {

	var token, err = GetToken(tokenStr)

	if err != nil {
		return false, err
	}

	if token.Valid {
		return true, nil
	}

	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return false, fmt.Errorf("token malformado")
		}
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return false, fmt.Errorf("token expirado")
		}
	}

	return false, fmt.Errorf("nao foi possivel processar o token")
}

func GetUserFromToken(tokenStr string) (string, error) {

	if _, err := Authorize(tokenStr); err != nil {
		return "", err
	}

	var token *jwt.Token
	token, err := GetToken(tokenStr)

	if err != nil {
		return "", err
	}

	// Acessa diretamente as claims
	claims := token.Claims.(jwt.MapClaims)
	value := claims["sub"]

	if subValue, ok := value.(string); ok {
		return subValue, nil // Retorna o valor como string
	}

	return "", fmt.Errorf("o valor da chave 'sub' não é uma string")
}

func GetTokenFromRequest(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", fmt.Errorf("cabeçalho Authorization não encontrado")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", fmt.Errorf("cabeçalho Authorization inválido")
	}

	return parts[1], nil
}
