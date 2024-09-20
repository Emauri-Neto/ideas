package types

import "time"

type User struct {
	Id        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"password" db:"password"`
	Expertise string    `json:"expertise" db:"expertise"`
	Interests []string  `json:"interests" db:"interests"`
	Roles     []string  `json:"roles" db:"roles"`
	Status    string    `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time   `json:"updated_at" db:"updated_at"`
}

type LoginCredentials struct {
	Email    string
	Password string
}

type RegisterCredentials struct {
	Id       string
	Email    string
	Password string
	Name     string
}