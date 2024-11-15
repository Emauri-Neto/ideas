package api

import (
	"fmt"
	"net/http"
	"server/db"
	"server/internal/auth"
	"server/internal/middleware"
	"server/internal/study"
	"server/internal/user"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type WebServer struct {
	addr string
	db *db.Database
	secret *auth.JWTConfig
}

func Serve(addr string, db *db.Database, secret *auth.JWTConfig) *WebServer {
	return &WebServer{
		addr: addr,
		db: db,
		secret: secret,
	}
}

func (s *WebServer) Run() error {
	router := mux.NewRouter()
	authrouter := router.PathPrefix("/auth").Subrouter()
	protectedrouter := router.PathPrefix("/api").Subrouter()

	authrouter.HandleFunc("/register", user.SignUp(s.db, s.secret)).Methods("POST")
	authrouter.HandleFunc("/login", user.SignIn(s.db, s.secret)).Methods("POST")
	authrouter.HandleFunc("/refresh", user.RefreshAccessToken(s.db, s.secret)).Methods("GET")
	authrouter.HandleFunc("/logout", user.Logout(s.db, s.secret)).Methods("GET")

	protectedrouter.HandleFunc("/study", study.CreateStudy(s.db)).Methods("POST")
	protectedrouter.HandleFunc("/study", study.ListStudies(s.db)).Methods("GET")
	protectedrouter.HandleFunc("/study/{stID}", study.GetStudy(s.db)).Methods("GET")

	protectedrouter.HandleFunc("/user", user.GetUser(s.db)).Methods("GET")

	protectedrouter.Use(middleware.IsAuthenticated(s.secret))

	cors := handlers.CORS(
		handlers.AllowCredentials(),
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	fmt.Printf("Server running on port%s 🔥", s.addr)
	return http.ListenAndServe(s.addr, cors(router))
}
