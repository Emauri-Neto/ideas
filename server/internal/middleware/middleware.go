package middleware

import (
	"context"
	"net/http"
	"server/internal/auth"
	"server/utils"
	"strings"
)

func IsAuthenticated(jwt *auth.JWTConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, _token := r.Cookie("accessToken")

			if _token == nil {
				id, _, err := jwt.Autorize(token.Value)

				if err != nil {
					utils.WriteResponse(w, http.StatusForbidden, err.Error())
					return
				}

				next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "UserID", id)))
				return
			}

			tkn := r.Header.Get("Authorization")

			if tkn == "" || !strings.HasPrefix(tkn, "Bearer ") {
				utils.WriteResponse(w, http.StatusForbidden, "Token inválido")
				return
			}

			tkn = strings.TrimPrefix(tkn, "Bearer ")

			id, _, err := jwt.Autorize(tkn)

				if err != nil {
					utils.WriteResponse(w, http.StatusForbidden, err.Error())
					return
				}

				next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), "UserID", id)))
		})
	}
}
