package api

import (
	"fmt"
	"ideas/db"
	"ideas/internal/auth"
	"ideas/internal/invitation"
	"ideas/internal/study"
	"ideas/internal/thread"
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
	publicrouter := router.PathPrefix("/s").Subrouter()
	subrouter := router.PathPrefix("/api").Subrouter()
	authrouter := router.PathPrefix("/auth").Subrouter()

	authrouter.HandleFunc("/register", user.SignUp(s.db)).Methods("POST")
	authrouter.HandleFunc("/login", user.SignIn(s.db)).Methods("POST")

	subrouter.HandleFunc("/user", user.GetUser(s.db)).Methods("GET")
	subrouter.HandleFunc("/user", user.UpdateUser(s.db)).Methods("PATCH")
	subrouter.HandleFunc("/user", user.DeleteUser(s.db)).Methods("DELETE")
	subrouter.HandleFunc("/user/invitations", invitation.ListInvitations(s.db)).Methods("GET")

	subrouter.HandleFunc("/invitation/{id}/accept", invitation.AcceptInvite(s.db)).Methods("GET")
	subrouter.HandleFunc("/invitation/{id}/refuse", invitation.RefuseInvite(s.db)).Methods("GET")

	subrouter.HandleFunc("/study/create", study.CreateStudy(s.db)).Methods("POST")
	publicrouter.HandleFunc("/study", study.GetAllStudies(s.db)).Methods("GET")
	subrouter.HandleFunc("/study/{id}", study.GetStudyById(s.db)).Methods("GET")
	subrouter.HandleFunc("/study/{id}", study.DeleteStudy(s.db)).Methods("DELETE")
	subrouter.HandleFunc("/study/{id}", study.UpdateStudy(s.db)).Methods("PUT")
	subrouter.HandleFunc("/study/{id}/thread", thread.CreateThread(s.db)).Methods("POST")
	subrouter.HandleFunc("/study/{id}/users", study.ListUsersStudy(s.db)).Methods("GET")
	subrouter.HandleFunc("/study/{id}/thread", thread.CreateThread(s.db)).Methods("POST")

	subrouter.HandleFunc("/thread/{id}/invite", invitation.CreateInvitation(s.db)).Methods("POST")
	subrouter.HandleFunc("/thread/{id}/users", study.ListUsersThread(s.db)).Methods("GET")

	subrouter.Use(auth.IsAuthenticated)

	fmt.Printf("Servidor aqui -> %s 🔥", s.addr)
	return http.ListenAndServe(s.addr, router)
}
