package api

import (
	"fmt"
	"ideas/db"
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
	
	fmt.Printf("Servidor aqui -> %s 🔥", s.addr)
	return http.ListenAndServe(s.addr, router)
}