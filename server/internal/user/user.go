package user

import (
	"net/http"
	"server/db"
	"server/utils"
)

func GetUser(db *db.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("UserID").(string)

		u, _u := db.GetUserById(user)

		if _u != nil {
			utils.WriteResponse(w, http.StatusBadRequest, _u.Error())
			return
		}

		utils.WriteResponse(w, http.StatusOK, u)
	}
}