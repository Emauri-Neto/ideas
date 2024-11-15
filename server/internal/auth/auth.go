package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type JWTConfig struct {
	secret []byte
}

func DefaultOpts(secret []byte) *JWTConfig {
	return &JWTConfig{
		secret: secret,
	}
}

func ValidatePassword(password, confirmPassword string) bool {
	return password == confirmPassword
}

func HashValue(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("could not hash value -> %w", err)
	}

	return string(hash), nil
}

func ComparePassword(hash string, value []byte) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), value)

	return err == nil
}

func (j JWTConfig) SignToken(id string, duration time.Duration, session *string) (string, error) {
	claims := jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(duration).Unix(),
	}

	if session != nil {
		claims["session"] = *session
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString(j.secret)

	if err != nil {
		return "", fmt.Errorf("could not sign token -> %w", err)
	}

	return t, nil
}

func (j JWTConfig) Autorize(token string) (string, string, error) {
	tkn, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})

	if tkn != nil && tkn.Valid {
		if claims, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {
			sub := claims["sub"].(string)
			session := claims["session"].(string)

			return sub, session, nil
		}
	}

	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return "", "",fmt.Errorf("malformed token")
		}
		if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
			return "", "",fmt.Errorf("token expired or not valid yet")
		}
	}

	return "", "", fmt.Errorf("token error: %w", err)
}