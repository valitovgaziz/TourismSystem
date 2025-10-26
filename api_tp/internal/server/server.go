package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"api_tp/internal/handlers"
	"api_tp/internal/middleware"
	"api_tp/internal/repository"
	"api_tp/internal/service"
	"time"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type Server struct {
	router *chi.Mux
	db     *gorm.DB
}

func New(db *gorm.DB) *Server {
	s := &Server{
		router: chi.NewRouter(),
		db:     db,
	}
	s.configureRouter(db)
	return s
}

func (s *Server) configureRouter(db *gorm.DB) {
	// Общие middleware
	for _, middleware := range handlers.CommonMiddleware() {
		s.router.Use(middleware)
	}

	// Health check
	s.router.Get("/health", s.healthCheck)

	// API routes
	s.router.Route("/v1", func(r chi.Router) {
		r.Get("/check", s.healthCheck)
		s.setupUserRoutes(r, db)
	})

	// Для отладки - выводим все маршруты
	chi.Walk(s.router, func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		fmt.Printf("[%s] %s\n", method, route)
		return nil
	})
}

func (s *Server) setupUserRoutes(r chi.Router, db *gorm.DB) {
	userRepo := repository.NewUserRepository(s.db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	authHandler := &handlers.AuthHandler{DB: db}
	oauthHandler := &handlers.OAuthHandler{DB: db}

	// Публичные маршруты
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
		r.Get("/check", s.healthCheck)

		// OAuth routes
		r.Get("/google", oauthHandler.GoogleLogin)
		r.Get("/google/callback", oauthHandler.GoogleCallback)
		r.Get("/yandex", oauthHandler.YandexLogin)
		r.Get("/yandex/callback", oauthHandler.YandexCallback)
		r.Get("/vk", oauthHandler.VKLogin)
		r.Get("/vk/callback", oauthHandler.VKCallback)
	})

	// Защищенные маршруты
	r.Route("/api", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)

		r.Route("/users", func(r chi.Router) {
			r.Get("/", userHandler.GetAllUsers)
			r.Post("/", userHandler.CreateUser)
			r.Get("/{id}", userHandler.GetUser)
			r.Get("/check", s.healthCheck)
		})

	})
}

func (s *Server) healthCheck(w http.ResponseWriter, r *http.Request) {
	// Проверяем соединение с БД
	sqlDB, err := s.db.DB()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusServiceUnavailable)
		return
	}

	if err := sqlDB.Ping(); err != nil {
		http.Error(w, "Database ping failed", http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":    "healthy",
		"timestamp": time.Now().UTC().Format(time.RFC1123),
	})
}

func (s *Server) Run(port string) error {
	return http.ListenAndServe(":"+port, s.router)
}
