package api

import (
	"fmt"
	"ideas/db"
	secure "ideas/internal/auth"
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

	subrouter.HandleFunc("/users", user.List(s.db)).Methods("GET")
	subrouter.HandleFunc("/users/{id}", user.GetById(s.db)).Methods("GET")
	subrouter.HandleFunc("/user/update", user.UpdateUser(s.db)).Methods("POST")
	subrouter.HandleFunc("/user/delete", user.DeleteUser(s.db)).Methods("DELETE")

	subrouter.HandleFunc("/study", study.CreateStudy(s.db)).Methods("POST")
	subrouter.HandleFunc("/study/{id}/thread", study.CreateThread(s.db)).Methods("POST")

	subrouter.Use(secure.IsAuthenticated)

	fmt.Printf("Servidor aqui -> %s 🔥", s.addr)
	return http.ListenAndServe(s.addr, router)
}
