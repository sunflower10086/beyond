package middleware

import (
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest/handler"
)

type AuthMiddleware struct {
	Secret string
}

func NewAuthMiddleware(secret string) *AuthMiddleware {
	return &AuthMiddleware{
		Secret: secret,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		fmt.Println(authHeader)
		if authHeader == "" {

			return
		}

		// reqToken, err := jwt.ParseToken(authHeader)
		// if err != nil {
		// 	return
		// }

		authHandler := handler.Authorize(m.Secret)

		//err := jwt.RefreshToken(authHeader)
		//if err != nil {
		//	return
		//}
		// Passthrough to next handler if need
		authHandler(next).ServeHTTP(w, r)
		//next(w, r)
	}
}
