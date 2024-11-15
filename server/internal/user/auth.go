package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"server/db"
	"server/internal/auth"
	"server/internal/mail"
	"server/types"
	"server/utils"
	"strings"
	"time"

	"github.com/google/uuid"
)

func SignIn(db *db.Database, jwt *auth.JWTConfig) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := struct {
			Email    string
			Password string
		}{}

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			utils.WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		u, _u := db.GetUserByEmail(user.Email)

		if _u != nil {
			utils.WriteResponse(w, http.StatusNotFound, _u.Error())
			return
		}

		if !auth.ComparePassword(u.Password, []byte(user.Password)) {
			utils.WriteResponse(w, http.StatusUnauthorized, "Invalid email or password")
			return
		}

		sId := uuid.New().String()

		_s := db.CreateSession(types.Session{
			Id:        sId,
			UserAgent: r.Header.Get("User-Agent"),
			UserId:    u.Id,
		})

		if _s != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, _s.Error())
			return
		}

		access, _access := jwt.SignToken(u.Id, 15*time.Minute, &sId)
		refresh, _refresh := jwt.SignToken(u.Id, 7*24*time.Hour, &sId)

		if _err := errors.Join(_access, _refresh); _err != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, _err.Error())
			return
		}

		utils.SetAuthCookies(w, access, refresh)

		res := struct {
			Access string
			Refresh string
		}{
			Access: access,
			Refresh: refresh,
		}

		utils.WriteResponse(w, http.StatusOK, res)
	}
}

func SignUp(db *db.Database, jwt *auth.JWTConfig) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := struct {
			Id              string
			Email           string
			Password        string
			ConfirmPassword string
		}{}

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			utils.WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		if !auth.ValidatePassword(user.Password, user.ConfirmPassword) {
			utils.WriteResponse(w, http.StatusBadRequest, "Passwords do not match")
			return
		}

		if _, _u := db.GetUserByEmail(user.Email); _u == nil {
			utils.WriteResponse(w, http.StatusConflict, "Email already in use")
			return
		}

		hash, _hash := auth.HashValue(user.Password)

		if _hash != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, _hash.Error())
			return
		}

		u := types.User{
			Id:       uuid.New().String(),
			Email:    user.Email,
			Password: hash,
		}

		v := types.VerificationCode{
			Id:        uuid.New().String(),
			VCType:    types.EmailVerification,
			UserId:    u.Id,
			ExpiresAt: time.Now().AddDate(1, 0, 0),
		}

		s := types.Session{
			Id:        uuid.New().String(),
			UserAgent: r.Header.Get("User-Agent"),
			UserId:    u.Id,
		}

		if _c := db.CreateAccount(u); _c != nil {
			utils.WriteResponse(w, http.StatusBadRequest, _c.Error())
			return
		}

		if _v := db.CreateVerificationCode(v); _v != nil {
			utils.WriteResponse(w, http.StatusBadRequest, _v.Error())
			return
		}

		if _m := mail.SendMail([]string{"delivered@resend.dev"}, mail.GetEmailVerificationTemplate(fmt.Sprintf("http://localhost:3333/email/verify/%s", v.Id))); _m != nil {
			fmt.Printf("%s", _m.Error())
		}

		if _s := db.CreateSession(s); _s != nil {
			utils.WriteResponse(w, http.StatusBadRequest, _s.Error())
			return
		}

		access, _access := jwt.SignToken(u.Id, 15*time.Minute, &s.Id)
		refresh, _refresh := jwt.SignToken(u.Id, 7*24*time.Hour, &s.Id)

		if err := errors.Join(_access, _refresh); err != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		utils.SetAuthCookies(w, access, refresh)

		res := struct {
			Access string
			Refresh string
		}{
			Access: access,
			Refresh: refresh,
		}

		utils.WriteResponse(w, http.StatusCreated, res)
	}
}

func RefreshAccessToken(db *db.Database, jwt *auth.JWTConfig) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")

		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			utils.WriteResponse(w, http.StatusUnauthorized, "Token inválido")
			return
		}

		token = strings.TrimPrefix(token, "Bearer ")

		id, session, err := jwt.Autorize(token)

		if err != nil {
			utils.WriteResponse(w, http.StatusUnauthorized, err)
			return
		}

		s, _s := db.GetSessionByID(session)

		if _s != nil{
			utils.WriteResponse(w, http.StatusUnauthorized, map[string]string{"message": "Sessão não encontrada ou inválida"})
			return
		}

		if time.Now().After(s.ExpiresAt) {
			utils.WriteResponse(w, http.StatusUnauthorized, map[string]string{"message": "Sessão expirada"})
			return
		}

		access, _access := jwt.SignToken(id, 15*time.Minute, &s.Id)

		if _access != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, map[string]string{"message": fmt.Sprintf("%s", _access)})
			return
		}

		if time.Until(s.ExpiresAt) <= 24*time.Hour {
			if updt := db.UpdateSession(s.Id); updt != nil {
				utils.WriteResponse(w, http.StatusInternalServerError, map[string]string{"message": fmt.Sprintf("%s", updt)})
				return
			}

			refresh , _refresh := jwt.SignToken(id, 7*24*time.Hour, &s.Id)

			if _refresh != nil {
				utils.WriteResponse(w, http.StatusInternalServerError, map[string]string{"message": fmt.Sprintf("%s", _refresh)})
				return
			}

			utils.SetAuthCookies(w, access, refresh)
			utils.WriteResponse(w, http.StatusOK, map[string]string{"message": "Access e Refresh Token atualizados"})
			return
		}

		utils.SetAuthCookies(w, access, "")
		utils.WriteResponse(w, http.StatusOK, map[string]string{"token": access})
	}
}

func Logout(db *db.Database, jwt *auth.JWTConfig) func(http.ResponseWriter, *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		c, _c := r.Cookie("accessToken")

		if _c != nil {
			utils.WriteResponse(w, http.StatusOK, map[string]string{"message": _c.Error()})
			return
		}

		_, session, err := jwt.Autorize(c.Value)

		if err != nil {
			utils.WriteResponse(w, http.StatusOK, map[string]string{"message": err.Error()})
			return
		}


		if _d := db.DeleteSessionById(session); _d != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, _d)
			return
		}

		utils.ClearAuthCookies(w)
		utils.WriteResponse(w, http.StatusOK, map[string]string{"message": "Logout feito com sucesso"})
	}
}