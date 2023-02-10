package middleware

import (
	"go-mindoc/helper"
	"net/http"
)

type AuthMiddleware struct {
	JwtKey string
}

func NewAuthMiddleware(jwtKey string) *AuthMiddleware {
	return &AuthMiddleware{JwtKey: jwtKey}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		uc, err := helper.AnalyzeToken(auth, m.JwtKey)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}
		r.Header.Set("UserId", string(rune(uc.Id)))
		r.Header.Set("UserName", uc.Username)
		next(w, r)
	}
}
