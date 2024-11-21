package thread

import (
	"encoding/json"
	"net/http"
	"server/db"
	"server/types"
	"server/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateThread(db *db.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var t types.Thread

		stID := mux.Vars(r)["stID"]

		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			utils.WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		userId := r.Context().Value("UserID").(string)

		_th := db.CreateThread(types.Thread{
			Id: uuid.New().String(),
			Name: t.Name,
			DeadLine: t.DeadLine,
			ResponsibleID: userId,
			StudyID: stID,
		})

		if _th != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, _th.Error())
			return
		}

		utils.WriteResponse(w, http.StatusCreated, "Thread Criada")
	}
}
