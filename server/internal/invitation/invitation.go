package invitation

import (
	"encoding/json"
	"ideas/db"
	"ideas/types"
	"ideas/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func CreateInvitation(db *db.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var invitationRequest types.RequestInvitation

		if err := json.NewDecoder(r.Body).Decode(&invitationRequest); err != nil {
			utils.WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		_, err := db.GetUsersById(invitationRequest.Receiver_id)

		if err != nil {
			utils.WriteResponse(w, http.StatusBadRequest, "receiver não existe")
			return
		}

		vars := mux.Vars(r)
		threadId := vars["id"]

		thread, err := db.GetThreadById(threadId)

		if err != nil {
			utils.WriteResponse(w, http.StatusNotFound, err.Error())
			return
		}

		userId := r.Context().Value("UserID").(string)

		if err := db.IsThreadResponsible(thread, userId); err != nil {
			utils.WriteResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		receiverId := invitationRequest.Receiver_id

		if err := db.ExistInvitation(thread.Id, receiverId); err != nil {
			utils.WriteResponse(w, http.StatusConflict, err.Error())
			return
		}

		invitation := types.Invitation{
			Id:              uuid.New().String(),
			Type_invitation: invitationRequest.Type_invitation,
			Text:            invitationRequest.Text,
			Thread_id:       threadId,
			Study_id:        thread.Study_id,
		}

		userInvitation := types.UserInvitation{
			Id:            uuid.New().String(),
			Invitation_id: invitation.Id,
			Sender_id:     userId,
			Receiver_id:   receiverId,
		}

		if _c := db.CreateInvitation(invitation, userInvitation); _c != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, _c.Error())
			return
		}

		utils.WriteResponse(w, http.StatusOK, "Invitation criada com sucesso")
	}
}
