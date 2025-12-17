package middleware

import (
	"context"
	"liveAt/utils"
	"net/http"
	"os"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == ""{
			utils.Error(w,http.StatusUnauthorized,"Unauthorized, token missing")
			return 
		}

		token := strings.TrimSpace(authHeader)

		claims, err := utils.VerifyJWT(token, os.Getenv("JWT_SECTRET"))
		if err!=nil {
			utils.Error(w,http.StatusUnauthorized,"Unauthorised, token invalid")
			return 
		}

		ctx := context.WithValue(r.Context(),UserContextKey,claims)

		next.ServeHTTP(w, r.WithContext(ctx))

	})
}