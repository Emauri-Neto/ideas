package types

import (
	"time"
)

type User struct {
	Id        string    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type LoginCredentials struct {
	Email    string
	Password string
}

type RegisterCredentials struct {
	Id              string
	Email           string
	Password        string
	ConfirmPassword string
	Name            string
}

type UpdateUser struct {
	Id   string
	Name string
}

type Study struct {
	Id                 string    `json:"id" db:"id"`
	Name               string    `json:"name" db:"name"`
	Objective          string    `json:"objective" db:"objective"`
	Methodology        string    `json:"methodology" db:"methodology"`
	Max_participants   *uint     `json:"max_participants,omitempty" db:"max_participants"`
	Participation_type *string   `json:"participation_type,omitempty" db:"participation_type"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	Responsible_id     string    `json:"responsible_id" db:"responsible_id"`
}

type Thread struct {
	Id                  string    `json:"id" db:"id"`
	Name                string    `json:"name" db:"name"`
	Max_participants    *uint     `json:"max_participantes,omitempty" db:"max_participants"`
	Discussion_deadline time.Time `json:"discussion_deadline" db:"discussion_deadline"`
	Status              string    `json:"status" db:"status"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
	Responsible_id      string    `json:"responsible_id" db:"responsible_id"`
	Study_id            string    `json:"study_id" db:"study_id"`
}

type Invitation struct {
	Id              string    `json:"id" db:"id"`
	Type_invitation string    `json:"type" db:"type"`
	Text            string    `json:"text" db:"text"`
	Accept          *bool     `json:"accept" db:"accept"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	Study_id        string    `json:"study_id" db:"study_id"`
	Thread_id       string    `json:"thread_id" db:"thread_id"`
}

type RequestInvitation struct {
	Type_invitation string `json:"type"`
	Text            string `json:"text"`
	Receiver_id     string `json:"receiver"`
}

type UserInvitation struct {
	Id            string
	Invitation_id string
	Sender_id     string
	Receiver_id   string
}

type Responsibles struct {
	Study_id           string `db:"study_id"`
	Study_responsible  string `db:"study_responsible"`
	Thread_responsible string `db:"thread_responsible"`
}

type UserResponse struct {
	Id    string `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Email string `json:"email" db:"email"`
}

type UserKey string
