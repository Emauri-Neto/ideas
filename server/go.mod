module server

go 1.23

toolchain go1.23.2

require (
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/google/uuid v1.6.0
	github.com/gorilla/mux v1.8.1
	github.com/jmoiron/sqlx v1.4.0
	github.com/lib/pq v1.10.9
	golang.org/x/crypto v0.28.0
)

require github.com/felixge/httpsnoop v1.0.3 // indirect

require (
	github.com/gorilla/handlers v1.5.2
	github.com/resend/resend-go/v2 v2.13.0 // indirect
)
