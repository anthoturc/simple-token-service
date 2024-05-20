package main

import (
	"net/http"
	"strings"
	"time"
)

type AuthMiddleware struct {
	db *DbService
}

func (amw *AuthMiddleware) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerVal := r.Header.Get("Authorization")
		if headerVal == "" {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		authComponents := strings.Split(headerVal, " ")
		if len(authComponents) != 2 || authComponents[0] != "Bearer" {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		token := authComponents[1]

		maybeApiToken, err := amw.db.GetApiToken(Hash(token))
		if err != nil {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		if !maybeApiToken.IsEnabled || maybeApiToken.ExpiresAt.Before(time.Now()) {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
