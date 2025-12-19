package middleware

import(
	"net/http"
	"liveAt/utils"
)

func TeacherOnly(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		claims := r.Context().Value(UserContextKey).(*utils.JwtClaims)

		if claims.Role != "teacher" {
			utils.Error(w,http.StatusForbidden,"forbideen, teacher access required")
			return 
		}

		next.ServeHTTP(w,r)
	})
}

func StudentOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		claims := r.Context().Value(UserContextKey).(*utils.JwtClaims)

		if claims.Role != "student" {
			utils.Error(w, http.StatusForbidden, "Forbidden, student access required")
			return
		}

		next.ServeHTTP(w, r)
	})
}
