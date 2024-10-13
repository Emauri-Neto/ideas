package thread

import (
	"encoding/json"
	"ideas/db"
	"ideas/types"
	"ideas/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var thread types.Thread

func CreateThread(db *db.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
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
