package middleware

import (
	"log"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/go-chi/jwtauth"
)

// Authorizer authorize middleware
func Authorizer(e *casbin.Enforcer) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())

			role, _ := claims["role"].(string)
			log.Println("Role: ", role)
			log.Println("Path: ", r.URL.Path)
			log.Println("Medthod: ", r.Method)
			res, err := e.Enforce(role, r.URL.Path, r.Method)

			if err != nil {
				w.WriteHeader(http.StatusForbidden)
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}

			if res {
				next.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusForbidden)
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}
		}
		return http.HandlerFunc(fn)
	}
}
