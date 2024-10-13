package study

import (
	"encoding/json"
	"ideas/db"
	"ideas/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func ListUsersStudy(db *db.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		studyID := mux.Vars(r)["id"]

		users, err := db.GetUsersByStudy(studyID)

		if err != nil {
			utils.WriteResponse(w, http.StatusNotFound, err.Error())
			return
		}

		json.NewEncoder(w).Encode(users)
	}
}

func ListUsersThread(db *db.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")

		vars := mux.Vars(r)
		threadId := vars["id"]

		users, err := db.GetUsersByThread(threadId)

		if err != nil {
			utils.WriteResponse(w, http.StatusNotFound, err.Error())
			return
		}

		json.NewEncoder(w).Encode(users)
	}
}
