package main

import (
    "log"
    "api_tp/internal/config"
    "api_tp/internal/server"
    "api_tp/pkg/database"
)

func main() {
    // Загрузка конфигурации
    cfg := config.Load()

    // Подключение к БД
    db, err := database.NewPostgresConnection(cfg)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Создание и запуск сервера
    srv := server.New(db)
    
    log.Printf("Server starting on port %s", cfg.AppPort)
    if err := srv.Run(cfg.AppPort); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}