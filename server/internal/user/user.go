package user

import (
	"encoding/json"
	"ideas/db"
	secure "ideas/internal/auth"
	"ideas/types"
	"ideas/utils"
	"net/http"

	"github.com/google/uuid"
)

func GetUser(db *db.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		userId := r.Context().Value("UserID").(string)

		u, _u := db.GetUsersById(userId)

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
			utils.WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if err := ValidEmail(user.Email); err != nil {
			utils.WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		u, _u := db.GetUserByEmail(user.Email)

		if _u != nil {
			utils.WriteResponse(w, http.StatusNotFound, _u.Error())
			return
		}

		if !secure.CompareValue(u.Password, []byte(user.Password)) {
			utils.WriteResponse(w, http.StatusUnauthorized, "Email ou Senha inválidos.")
			return
		}

		secret, _ := utils.GetEnv("JWT_SECRET")

		sec := []byte(secret)

		token, err := secure.SignToken(sec, u.Id)

		if err != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.WriteResponse(w, http.StatusOK, map[string]string{"token": token})
	}
}

func SignUp(db *db.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user types.RegisterCredentials

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			utils.WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if err := ValidUser(&user); err != nil {
			utils.WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		_, _u := db.GetUserByEmail(user.Email)

		if _u == nil {
			utils.WriteResponse(w, http.StatusConflict, "Email já está em uso.")
			return
		}

		hash, _hash := secure.HashValue(user.Password)

		if _hash != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, _hash.Error())
			return
		}

		_c := db.CreateAccount(types.User{
			Id:       uuid.New().String(),
			Name:     user.Name,
			Email:    user.Email,
			Password: hash,
		})

		if _c != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, _c.Error())
			return
		}

		utils.WriteResponse(w, http.StatusCreated, "Usuário criado com sucesso")
	}
}

func UpdateUser(db *db.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var user types.UpdateUser

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			utils.WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if err := ValidUserUpdate(&user); err != nil {
			utils.WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		userId := r.Context().Value("UserID").(string)

		_c := db.UpdateUser(types.User{
			Id:   userId,
			Name: user.Name,
		})

		if _c != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, _c.Error())
			return
		}

		utils.WriteResponse(w, http.StatusCreated, "Usuário atualizado com sucesso")
	}
}

func DeleteUser(db *db.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		userId := r.Context().Value("UserID").(string)

		_c := db.DeleteUser(userId)

		if _c != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, _c.Error())
			return
		}

		utils.WriteResponse(w, http.StatusCreated, "Usuário deletado com sucesso")
	}
}