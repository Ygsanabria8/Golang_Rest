package middleWares

import (
	"net/http"

	jwt "modules/src/jwt"
)

func JwtValidation(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		valid, err := jwt.ProccessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
		}
		if !valid {
			http.Error(w, "Expired token", http.StatusUnauthorized)
		}
		next.ServeHTTP(w, r)
	}
}
