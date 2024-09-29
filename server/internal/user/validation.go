package user

import (
	"errors"
	"ideas/types"
	"net/mail"
	"strings"
	"unicode"
)

func ValidUser(user *types.RegisterCredentials) error {

	name, err := ValidName(user.Name)
	if err != nil {
		return err
	}
	user.Name = name

	if err := ValidEmail(user.Email); err != nil {
		return err
	}

	if err := ValidPassword(user.Password, user.ConfirmPassword); err != nil {
		return err
	}

	return nil
}

func ValidEmail(email string) error {
	if _, err := mail.ParseAddress(email); err != nil {
		return err
	}

	return nil
}

func ValidPassword(password, confirmPassword string) error {

	if password != confirmPassword {
		return errors.New("as senhas não coincidem")
	}

	if len(password) < 6 {
		return errors.New("senha curta")
	}

	var hasNumber, hasSpecial, hasUpper, hasLower bool

	for _, code := range password {
		if !hasNumber && unicode.IsNumber(code) {
			hasNumber = true
			continue
		}
		if !hasUpper && unicode.IsUpper(code) {
			hasUpper = true
			continue
		}
		if !hasLower && unicode.IsLower(code) {
			hasLower = true
			continue
		}
		if !hasSpecial && (unicode.IsPunct(code) || unicode.IsSymbol(code)) {
			hasSpecial = true
			continue
		}
	}

	validations := map[bool]string{
		!hasUpper:   "a senha deve conter pelo menos uma letra maiúscula",
		!hasLower:   "a senha deve conter pelo menos uma letra minúscula",
		!hasNumber:  "a senha deve conter pelo menos um número",
		!hasSpecial: "a senha deve conter pelo menos um caractere especial",
	}

	for condition, message := range validations {
		if condition {
			return errors.New(message)
		}
	}

	return nil
}

func ValidName(name string) (string, error) {
	bitName := strings.Fields(name)

	if len(bitName) == 0 {
		return "", errors.New("nome não informado")
	}

	for _, subName := range bitName {
		if len(subName) < 2 {
			return "", errors.New("nome inválido")
		}
	}

	name = strings.Join(bitName, " ")

	return name, nil
}
