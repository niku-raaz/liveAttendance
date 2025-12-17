package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Incoming requests: %s, %s",r.Method,r.URL.Path)

		next.ServeHTTP(w,r)
	})
}
func LogRequestTime(next http.Handler) http.Handler{

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start :=time.Now()
		next.ServeHTTP(w,r)
		log.Println(time.Since(start))
	})
}