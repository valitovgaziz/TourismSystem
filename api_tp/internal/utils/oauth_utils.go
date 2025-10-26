// utils/oauth_utils.go
package utils

import (
	"crypto/rand"
	"fmt"
	"api_tp/internal/models"

	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type OAuthHandler struct {
	DB *gorm.DB
}

// GenerateState generates a random state string for OAuth
func GenerateState() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func (h *OAuthHandler) findOrCreateOAuthUser(provider, providerID, email, name string, token *oauth2.Token) (*models.User, error) {
	var oauthProvider models.OAuthProvider

	err := h.DB.Where("provider = ? AND provider_id = ?", provider, providerID).
		Preload("User").
		First(&oauthProvider).Error

	if err == nil {
		// Обновляем токены существующей привязки
		oauthProvider.AccessToken = token.AccessToken
		oauthProvider.RefreshToken = token.RefreshToken
		oauthProvider.ExpiresAt = token.Expiry
		if err := h.DB.Save(&oauthProvider).Error; err != nil {
			return nil, err
		}

		var user models.User
		if err := h.DB.First(&user, oauthProvider.UserID).Error; err != nil {
			return nil, err
		}
		return &user, nil
	}

	// Ищем пользователя по email
	var user models.User
	err = h.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		// Создаем нового пользователя
		user = models.User{
			Email:    email,
			Name:     name,
			Password: GenerateRandomPassword(),
		}
		if err := h.DB.Create(&user).Error; err != nil {
			return nil, err
		}
	}

	// Создаем новую привязку OAuth с токенами
	oauthProvider = models.OAuthProvider{
		UserID:       user.ID,
		Provider:     provider,
		ProviderID:   providerID,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresAt:    token.Expiry,
	}

	if err := h.DB.Create(&oauthProvider).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
