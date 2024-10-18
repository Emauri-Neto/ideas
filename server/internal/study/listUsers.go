package study

import (
	"ideas/db"
	"ideas/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func ListUsersStudy(db *db.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		studyID := mux.Vars(r)["id"]

		users, err := db.GetUsersByStudy(studyID)

		if err != nil {
			utils.WriteResponse(w, http.StatusNotFound, err.Error())
			return
		}

		utils.WriteResponse(w, http.StatusNotFound, users)
	}
}

func ListUsersThread(db *db.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		threadId := mux.Vars(r)["id"]

		users, err := db.GetUsersByThread(threadId)

		if err != nil {
			utils.WriteResponse(w, http.StatusNotFound, err.Error())
			return
		}

		utils.WriteResponse(w, http.StatusNotFound, users)
	}
}
