package api

import (
	"fmt"
	"ideas/db"
	secure "ideas/internal/auth"
	"ideas/internal/invitation"
	"ideas/internal/study"
	"ideas/internal/user"
	"net/http"

	"github.com/gorilla/mux"
)

type WebServer struct {
	addr string
	db   *db.Database
}

func Serve(addr string, db *db.Database) *WebServer {
	return &WebServer{
		addr: addr,
		db:   db,
	}
}

func (s *WebServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api").Subrouter()
	auth := router.PathPrefix("/auth").Subrouter()

	auth.HandleFunc("/register", user.SignUp(s.db)).Methods("POST")
	auth.HandleFunc("/login", user.SignIn(s.db)).Methods("POST")

	subrouter.HandleFunc("/user", user.GetUser(s.db)).Methods("GET")
	subrouter.HandleFunc("/user/update", user.UpdateUser(s.db)).Methods("PUT")
	subrouter.HandleFunc("/user/delete", user.DeleteUser(s.db)).Methods("DELETE")
	subrouter.HandleFunc("/user/invitations", invitation.ListInvitations(s.db)).Methods("GET")

	subrouter.HandleFunc("/invitation/{id}/accept", invitation.AcceptInvite(s.db)).Methods("GET")
	subrouter.HandleFunc("/invitation/{id}/refuse", invitation.RefuseInvite(s.db)).Methods("GET")

	subrouter.HandleFunc("/study", study.CreateStudy(s.db)).Methods("POST")
	subrouter.HandleFunc("/study/{id}/thread", study.CreateThread(s.db)).Methods("POST")
	subrouter.HandleFunc("/study/{id}/users", study.ListUsersStudy(s.db)).Methods("GET")

	subrouter.HandleFunc("/thread/{id}/invite", invitation.CreateInvitation(s.db)).Methods("POST")
	subrouter.HandleFunc("/thread/{id}/users", study.ListUsersThread(s.db)).Methods("GET")

	subrouter.Use(secure.IsAuthenticated)

	fmt.Printf("Servidor aqui -> %s 🔥", s.addr)
	return http.ListenAndServe(s.addr, router)
}
