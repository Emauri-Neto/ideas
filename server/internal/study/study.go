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

		_s := db.CreateStudy(types.Study{
			Id:             uuid.New().String(),
			Name:           study.Name,
			Objective:      study.Objective,
			Methodology:    study.Methodology,
			Responsible_id: userId,
		})

		if _s != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, _s.Error())
			return
		}

		utils.WriteResponse(w, http.StatusCreated, "Estudo criado com sucesso")
	}

}

func GetAllStudies(db *db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		studies, err := db.GetAllStudy()

		if err != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		json.NewEncoder(w).Encode(studies)
	}
}

func GetStudyById(db *db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		studyId := mux.Vars(r)["id"]

		study, err := db.GetStudyById(studyId)

		if err != nil {
			utils.WriteResponse(w, http.StatusNotFound, "Estudo não encontrado")
			return
		}

		json.NewEncoder(w).Encode(study)
	}
}

func DeleteStudy(db *db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		studyId := mux.Vars(r)["id"]

		userId := r.Context().Value("UserID").(string)

		if err := db.IsStudyOwner(studyId, userId); err != nil {
			utils.WriteResponse(w, http.StatusUnauthorized, "Você não tem permissão para remover este estudo")
			return
		}

		err := db.DeleteStudy(studyId)
		if err != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, "Erro ao remover estudo")
			return
		}

		utils.WriteResponse(w, http.StatusOK, "Estudo removido com sucesso")
	}
}

func UpdateStudy(db *db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		studyId := mux.Vars(r)["id"]

		var study types.Study
		
		if err := json.NewDecoder(r.Body).Decode(&study); err != nil {
			utils.WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		userId := r.Context().Value("UserID").(string)

		if err := db.IsStudyOwner(studyId, userId); err != nil {
			utils.WriteResponse(w, http.StatusUnauthorized, "Você não tem permissão para alterar este estudo")
			return
		}

		study.Id = studyId
		err := db.UpdateStudy(study)
		if err != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, "Erro ao atualizar estudo")
			return
		}

		utils.WriteResponse(w, http.StatusOK, "Estudo atualizado com sucesso")
	}
}
