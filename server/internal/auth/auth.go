package auth

import (
	"context"
	"fmt"
	"ideas/types"
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

func Authorize(token *jwt.Token) (bool, error) {

	if token.Valid {
		return true, nil
	}

	var err error
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

	var token *jwt.Token
	token, err := GetToken(tokenStr)

	if err != nil {
		return "", err
	}

	if _, err := Authorize(token); err != nil {
		return "", err
	}

	claims := token.Claims.(jwt.MapClaims)
	value := claims["sub"]

	if subValue, ok := value.(string); ok {
		return subValue, nil
	}

	return "", fmt.Errorf("o valor da chave 'sub' não é uma string")
}

func GetUserFromTokenMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := GetTokenFromRequest(r)
		if err != nil {
			utils.WriteResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		var userId string
		userId, err = GetUserFromToken(token)
		if err != nil {
			utils.WriteResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		var idKey types.UserKey = "userId"

		ctx := context.WithValue(r.Context(), idKey, userId)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
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
