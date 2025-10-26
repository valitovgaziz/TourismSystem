// middleware/auth.go
package middleware

import (
	"context"
	"net/http"
	"api_tp/internal/utils"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            utils.WriteError(w, http.StatusUnauthorized, "Authorization header required")
            return
        }
        
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            utils.WriteError(w, http.StatusUnauthorized, "Invalid authorization format")
            return
        }
        
        claims, err := utils.ValidateJWT(parts[1])
        if err != nil {
            utils.WriteError(w, http.StatusUnauthorized, "Invalid token")
            return
        }
        
        // Добавляем claims в контекст
        ctx := context.WithValue(r.Context(), "userClaims", claims)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}