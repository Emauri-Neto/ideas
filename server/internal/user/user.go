package user

import (
	"encoding/json"
	"ideas/db"
	"ideas/internal/auth"
	"ideas/types"
	"ideas/utils"
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
			utils.WriteResponse(w, http.StatusNotFound, _u.Error())
		}

		if !auth.CompareValue(u.Password, []byte(user.Password)) {
			utils.WriteResponse(w, http.StatusUnauthorized, "Email ou Senha inválidos.")
		}

		utils.WriteResponse(w, http.StatusOK, "Login realizado com sucesso.")
	}
}

func SignUp(db *db.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user types.RegisterCredentials

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			utils.WriteResponse(w, http.StatusBadRequest, err.Error())
		}

		_, _u := db.GetUserByEmail(user.Email)

		if _u == nil {
			utils.WriteResponse(w, http.StatusConflict, "Email já está em uso.")
		}

		hash, _hash := auth.HashValue(user.Password)

		if _hash != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, _hash.Error())
		}

		_c := db.CreateAccount(types.User{
			Id:       uuid.New().String(),
			Name:     user.Name,
			Email:    user.Email,
			Password: hash,
		})

		if _c != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, _c.Error())
		}

		w.WriteHeader(http.StatusCreated)
	}
}