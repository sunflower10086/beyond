package middleware

import (
	"beyond/pkg/jwt"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {

			return
		}

		// reqToken, err := jwt.ParseToken(authHeader)
		// if err != nil {
		// 	return
		// }

		err := jwt.RefreshToken(authHeader)
		if err != nil {
			return
		}
		// Passthrough to next handler if need
		next(w, r)
	}
}
