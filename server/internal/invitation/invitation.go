package invitation

import (
	"encoding/json"
	"errors"
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

		threadId := mux.Vars(r)["id"]

		responsibles, err := db.GetResponsibleAndStudyId(threadId)

		if err != nil {
			utils.WriteResponse(w, http.StatusNotFound, err.Error())
			return
		}

		userId := r.Context().Value("UserID").(string)

		if responsibles.Study_responsible != userId && responsibles.Thread_responsible != userId {
			utils.WriteResponse(w, http.StatusUnauthorized, "usuario não autorizado")
			return
		}

		receiverId := invitationRequest.Receiver_id

		if err := db.ExistInvitationAndUser(threadId, receiverId); err != nil {
			utils.WriteResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		invitation := types.Invitation{
			Id:              uuid.New().String(),
			Type_invitation: invitationRequest.Type_invitation,
			Text:            invitationRequest.Text,
			Thread_id:       threadId,
			Study_id:        responsibles.Study_id,
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

func ListInvitations(db *db.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("UserID").(string)

		u, _u := db.GetInvitationsByReceiver(userId)

		if _u != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, _u.Error())
			return
		}

		if err := json.NewEncoder(w).Encode(u); err != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
}

func AcceptRefuseInvite(db *db.Database) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		status := r.URL.Query().Get("status")

		if status != "accept" && status != "refuse" {
			utils.WriteResponse(w, http.StatusUnauthorized, errors.New("opção inválida"))
			return
		}

		invID := mux.Vars(r)["id"]

		userId := r.Context().Value("UserID").(string)

		invitation, err := db.GetInvitationOwner(invID, userId)

		if err != nil {
			utils.WriteResponse(w, http.StatusUnauthorized, err.Error())
			return
		}

		if err := db.AcceptRefuseInvite(invID, status); err != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if err := db.CreateMiddleTableUser(userId, invitation); err != nil {
			utils.WriteResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		msg := "Convite " + status
		utils.WriteResponse(w, http.StatusOK, msg)
	}
}
