package middleware

import "net/http"

func ContentTypeMiddleware(next http.Handler) http.Handler {
	/*
		Middleware para settar todas as Content-Types que forem
		utilizadas como application/json
	*/
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-type", "application/json")
		next.ServeHTTP(w, r)
	})
}