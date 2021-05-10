package auth

import (
	"net/http"
)

const (
	authCookieKey   = "auth"
	authCookieValue = "shepherd"
)

func Wrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authCookie, err := r.Cookie(authCookieKey)
		if err != nil {
			http.Error(w, "missing auth cookie!", http.StatusUnauthorized)
			return
		}
		if authCookie.Value != authCookieValue {
			http.Error(w, "invalid auth cookie!", http.StatusForbidden)
			return
		}
		h.ServeHTTP(w, r)
	})
}
