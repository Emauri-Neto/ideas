package study

import (
	"encoding/json"
	"ideas/db"
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
			utils.WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		userId := r.Context().Value("UserID").(string)

		_c := db.CreateStudy(types.Study{
			Id:             uuid.New().String(),
			Name:           study.Name,
			Responsible_id: userId,
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
			utils.WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		vars := mux.Vars(r)
		studyId := vars["id"]

		userId := r.Context().Value("UserID").(string)

		err := db.IsStudyOwner(studyId, userId)
		if err != nil {
			utils.WriteResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		_c := db.CreateThread(types.Thread{
			Id:             uuid.New().String(),
			Name:           thread.Name,
			Responsible_id: userId,
			Study_id:       studyId,
		})

		if _c != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, _c.Error())
			return
		}

		utils.WriteResponse(w, http.StatusCreated, "Thread Criada com sucesso")
	}
}
