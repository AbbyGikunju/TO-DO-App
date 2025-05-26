package middleware

import (
	"net/http"
	"todo-app/utils"
)

// AuthMiddleware ensures that a user is logged in before accessing protected routes
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sessionToken, err := r.Cookie("session_token")
		if err != nil || !utils.IsSessionValid(sessionToken.Value) {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
