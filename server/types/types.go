package types

import (
	"time"
)

type User struct {
	Id        string    `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`
	Verified  bool      `json:"_verified" db:"_verified"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type VerificationCodeType string

const (
	EmailVerification VerificationCodeType = "email_verification"
	PasswordReset     VerificationCodeType = "password_reset"
)

type VerificationCode struct {
	Id        string               `json:"id" db:"id"`
	VCType    VerificationCodeType `json:"type" db:"type"`
	UserId    string               `json:"user_id" db:"user_id"`
	CreatedAt time.Time            `json:"created_at" db:"created_at"`
	ExpiresAt time.Time            `json:"expires_at" db:"expires_at"`
}

type Session struct {
	Id        string    `json:"id" db:"id"`
	UserAgent string    `json:"user_agent,omitempty" db:"user_agent"`
	UserId    string    `json:"user_id" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
}

type Study struct {
	Id                string  `json:"id" db:"id"`
	Title             string  `json:"title" db:"title"`
	Objective         string  `json:"objective" db:"objective"`
	Methodology       *string `json:"methodology" db:"methodology"`
	MaxParticipants   *int    `json:"max_participants" db:"max_participants"`
	NumParticipants   *int    `json:"num_participants" db:"num_participants"`
	ParticipationType *string `json:"participation_type" db:"participation_type"`
	Private           bool    `json:"_private" db:"_private"`
	UserID            string  `json:"user_id" db:"user_id"`
	CreatedAt         string  `json:"created_at" db:"created_at"`
	UpdatedAt         string  `json:"updated_at" db:"updated_at"`
}
