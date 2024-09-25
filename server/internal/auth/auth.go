package auth

import (
	"fmt"
	"ideas/utils"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func HashValue(pw string) (string, error){
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

func SignToken(secret []byte, userId string) (string, error){
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

func Authorize(tokenStr string) (bool, error){
	secret, _secret := utils.GetEnv("JWT_SECRET")

	if _secret != nil {
		return false, fmt.Errorf("%s", _secret)
	}

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("método de assinatura inesperado: %v", t.Header["alg"])
        }
        return []byte(secret), nil
    })

    if err != nil {
        return false, fmt.Errorf("%w", err)
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