package auth

import "golang.org/x/crypto/bcrypt"

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