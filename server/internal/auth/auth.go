package auth

import (
	"context"
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

func Authorize(token string) (string, error) {
	secret, _secret := utils.GetEnv("JWT_SECRET")

	if _secret != nil {
		return "", fmt.Errorf("%s", _secret)
	}

	tkn, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if tkn != nil && tkn.Valid {
		if claims, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {
			sub := claims["sub"].(string)
			
			return sub, nil
		}
	}

	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return "", fmt.Errorf("token malformado")
		}
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return "", fmt.Errorf("token expirado")
		}
	}

	return "", fmt.Errorf("nao foi possivel processar o token")
}

func IsAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			utils.WriteResponse(w, http.StatusUnauthorized, "Token inválido")
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		id, err := Authorize(token)

		if err != nil {
			utils.WriteResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		r = r.WithContext(context.WithValue(r.Context(), "UserID", id))

		next.ServeHTTP(w, r)
	})
}