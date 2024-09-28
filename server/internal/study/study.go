package study

import (
	"encoding/json"
	"ideas/db"
	"ideas/internal/auth"
	"ideas/types"
	"ideas/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateStudy(db *db.Database) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		var study types.Study

		if err := json.NewDecoder(r.Body).Decode(&study); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var token, err = auth.GetTokenFromRequest(r)
		if err != nil {
			utils.WriteResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		var userId string
		userId, err = auth.GetUserFromToken(token)

		if err != nil {
			utils.WriteResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		var data_creation = types.Time{Id: uuid.New().String()}

		_c := db.CreateTime(data_creation)

		if _c != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, _c.Error())
			return
		}

		_c = db.CreateStudy(types.Study{
			Id:             uuid.New().String(),
			Name:           study.Name,
			Responsible_id: userId,
			Time_id:        data_creation.Id,
		})

		if _c != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, _c.Error())
			return
		}

		utils.WriteResponse(w, http.StatusCreated, "Estudo criado com sucesso")
	}

}

func CreateThread(db *db.Database) func(http.ResponseWriter, *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		var thread types.Thread

		if err := json.NewDecoder(r.Body).Decode(&thread); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		vars := mux.Vars(r)
		studyId := vars["id"]

		var tokenStr, err = auth.GetTokenFromRequest(r)

		if err != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		var userId string
		userId, err = auth.GetUserFromToken(tokenStr)

		if err != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		err = db.IsStudyOwner(studyId, userId)
		if err != nil {
			utils.WriteResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		var data_creation = types.Time{Id: uuid.New().String()}
		_c := db.CreateTime(data_creation)

		if _c != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, _c.Error())
			return
		}

		_c = db.CreateThread(types.Thread{
			Id:             uuid.New().String(),
			Name:           thread.Name,
			Responsible_id: userId,
			Time_id:        data_creation.Id,
			Study_id:       studyId,
		})

		if _c != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, _c.Error())
			return
		}

		utils.WriteResponse(w, http.StatusCreated, "Thread Criada com sucesso")
	}
}
