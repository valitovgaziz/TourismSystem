package handlers

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func CommonMiddleware() []func(http.Handler) http.Handler {
    return []func(http.Handler) http.Handler{
        middleware.Logger,
        middleware.Recoverer,
        middleware.Timeout(60 * time.Second),
        cors.Handler(cors.Options{
            AllowedOrigins:   []string{"https://*", "http://*"},
            AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
            AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
            ExposedHeaders:   []string{"Link"},
            AllowCredentials: false,
            MaxAge:           300,
        }),
    }
}