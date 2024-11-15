package study

import (
	"encoding/json"
	"net/http"
	"server/db"
	"server/types"
	"server/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateStudy(db *db.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request){
		var s types.Study

		if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
			utils.WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		userId := r.Context().Value("UserID").(string)

		st, _st := db.CreateStudy(types.Study{
			Id: uuid.New().String(),
			Title: s.Title,
			Objective: s.Objective,
			Methodology: s.Methodology,
			MaxParticipants: s.MaxParticipants,
			Private: s.Private,
			UserID: userId,
		})

		if _st != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, _st.Error())
			return
		}

		utils.WriteResponse(w, http.StatusCreated, st)
	}
}

func ListStudies(db *db.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request){
		userId := r.Context().Value("UserID").(string)

		st, _st := db.ListStudies(userId)

		if _st != nil {
			utils.WriteResponse(w, http.StatusBadRequest, _st.Error())
			return
		}

		utils.WriteResponse(w, http.StatusOK, st)
	}
}

func GetStudy(db *db.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request){
		id := mux.Vars(r)["stID"]

		st, _st := db.GetStudy(id)

		if _st != nil {
			utils.WriteResponse(w, http.StatusOK, _st)
			return
		}

		utils.WriteResponse(w, http.StatusOK, st)
	}
}