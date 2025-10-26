// handlers/oauth.go
package handlers

import (
	"encoding/json"
	"net/http"
	"api_tp/internal/config"
	"api_tp/internal/models"
	"api_tp/internal/utils"

	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

type OAuthHandler struct {
    DB *gorm.DB
}

type GoogleUserInfo struct {
    ID    string `json:"id"`
    Email string `json:"email"`
    Name  string `json:"name"`
}


func (h *OAuthHandler) GoogleLogin(w http.ResponseWriter, r *http.Request) {
    url := config.GoogleOAuthConfig.AuthCodeURL("state")
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func (h *OAuthHandler) GoogleCallback(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    
    token, err := config.GoogleOAuthConfig.Exchange(r.Context(), code)
    if err != nil {
        utils.WriteError(w, http.StatusBadRequest, "Failed to exchange token")
        return
    }
    
    client := config.GoogleOAuthConfig.Client(r.Context(), token)
    resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
    if err != nil {
        utils.WriteError(w, http.StatusBadRequest, "Failed to get user info")
        return
    }
    defer resp.Body.Close()
    
    var userInfo GoogleUserInfo
    if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
        utils.WriteError(w, http.StatusBadRequest, "Failed to decode user info")
        return
    }
    
    // Создаем или находим пользователя
    user, err := h.findOrCreateOAuthUser("google", userInfo.ID, userInfo.Email, userInfo.Name, token)
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, "Error processing user")
        return
    }
    
    jwtToken, err := utils.GenerateJWT(user.ID, user.Email)
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, "Error generating token")
        return
    }
    
    // Редирект или возврат токена
    utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
        "token": jwtToken,
        "user":  user,
    })
}

// Аналогичные методы для Yandex и VK...

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
            Password: utils.GenerateRandomPassword(),
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