package study

import (
	"encoding/json"
	"fmt"
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

func GetAllStudies(db *db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		studies, err := db.GetAllStudy()
		if err != nil {

			fmt.Println("Erro ao recuperar estudos: ", err)
			utils.WriteResponse(w, http.StatusInternalServerError, "Erro ao recuperar estudos")
			return
		}

		studiesJSON, err := json.Marshal(studies)
		if err != nil {

			fmt.Println("Erro ao converter estudos para JSON: ", err)
			utils.WriteResponse(w, http.StatusInternalServerError, "Erro ao converter estudos para JSON")
			return
		}

		utils.WriteResponse(w, http.StatusOK, string(studiesJSON))
	}
}

func GetById(db *db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		studyId := vars["id"]

		study, err := db.GetStudyById(studyId)
		if err != nil {
			utils.WriteResponse(w, http.StatusNotFound, "Estudo não encontrado")
			return
		}

		studyJSON, err := json.Marshal(study)
		if err != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, "Erro ao converter o estudo para JSON")
			return
		}

		utils.WriteResponse(w, http.StatusOK, string(studyJSON))
	}
}

func DeleteStudy(db *db.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		studyId := vars["id"]

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
		vars := mux.Vars(r)
		studyId := vars["id"]

		var study types.Study
		if err := json.NewDecoder(r.Body).Decode(&study); err != nil {
			fmt.Printf("Erro ao decodificar JSON: %v", err)
			utils.WriteResponse(w, http.StatusBadRequest, "Dados inválidos")
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
