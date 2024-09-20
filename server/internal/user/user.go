package user

import (
	"encoding/json"
	"ideas/db"
	"ideas/internal/auth"
	"ideas/types"
	"net/http"

	"github.com/google/uuid"
)



func List(db *db.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		u, _u := db.GetUsers()

		if _u != nil {
			http.Error(w, _u.Error(), http.StatusInternalServerError)
			return
		}

		if err := json.NewEncoder(w).Encode(u); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func SignIn(db *db.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user types.LoginCredentials

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		u, _u := db.GetUserByEmail(user.Email)

		if _u != nil {
			http.Error(w, _u.Error(), http.StatusNotFound)
			return
		}

		if !auth.CompareValue(u.Password, []byte(user.Password)) {
			http.Error(w, "invalid email or password", http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func SignUp(db *db.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user types.RegisterCredentials

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, _u := db.GetUserByEmail(user.Email)

		if _u == nil {
			http.Error(w, "email already in use", http.StatusConflict)
			return
		}

		hash, _hash := auth.HashValue(user.Password)

		if _hash != nil {
			http.Error(w, _hash.Error(), http.StatusInternalServerError)
			return
		}

		_c := db.CreateAccount(types.User{
			Id:       uuid.New().String(),
			Name:     user.Name,
			Email:    user.Email,
			Password: hash,
		})

		if _c != nil {
			http.Error(w, _c.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
