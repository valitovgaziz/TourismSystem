package models

import (
	"time"

	"gorm.io/gorm"
)

type OAuthProvider struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	UserID       uint           `json:"user_id" gorm:"not null;index:idx_user_provider"`          // Уникальный индекс с провайдером
	Provider     string         `json:"provider" gorm:"not null;index:idx_user_provider;size:50"` // Ограничение длины
	ProviderID   string         `json:"provider_id" gorm:"not null;uniqueIndex:uix_provider_id"`  // Уникальный идентификатор
	AccessToken  string         `json:"access_token"`
	RefreshToken string         `json:"refresh_token"`
	ExpiresAt    time.Time      `json:"expires_at"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"` // Добавлено для отслеживания изменений
	DeletedAt    gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}