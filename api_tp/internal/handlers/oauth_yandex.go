package handlers

import (
	"encoding/json"
	"net/http"
	"api_tp/internal/config"
	"api_tp/internal/models"
	"api_tp/internal/utils"

	"golang.org/x/oauth2"
)

type YandexUserInfo struct {
	ID            string `json:"id"`
	Login         string `json:"login"`
	Email         string `json:"default_email"`
	DisplayName   string `json:"display_name"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	RealName      string `json:"real_name"`
	IsAvatarEmpty bool   `json:"is_avatar_empty"`
}

// YandexLogin initiates Yandex OAuth flow
func (h *OAuthHandler) YandexLogin(w http.ResponseWriter, r *http.Request) {
	url := config.YandexOAuthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// YandexCallback handles Yandex OAuth callback
func (h *OAuthHandler) YandexCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	if state != "state" {
		utils.WriteError(w, http.StatusBadRequest, "Invalid state parameter")
		return
	}

	token, err := config.YandexOAuthConfig.Exchange(r.Context(), code)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Failed to exchange token: "+err.Error())
		return
	}

	client := config.YandexOAuthConfig.Client(r.Context(), token)

	// Получаем информацию о пользователе
	resp, err := client.Get("https://login.yandex.ru/info?format=json")
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Failed to get user info: "+err.Error())
		return
	}
	defer resp.Body.Close()

	var userInfo YandexUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Failed to decode user info: "+err.Error())
		return
	}

	// Формируем имя пользователя
	name := h.getYandexUserName(userInfo)

	// Создаем или находим пользователя
	user, err := h.findOrCreateOAuthUser("yandex", userInfo.ID, userInfo.Email, name, token)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error processing user: "+err.Error())
		return
	}

	jwtToken, err := utils.GenerateJWT(user.ID, user.Email)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, "Error generating token: "+err.Error())
		return
	}

	h.handleOAuthSuccess(w, r, jwtToken, user)
}

func (h *OAuthHandler) handleOAuthSuccess(w http.ResponseWriter, r *http.Request, jwtToken string, user *models.User) {
	panic("unimplemented")
}

// getYandexUserName формирует имя пользователя из данных Yandex
func (h *OAuthHandler) getYandexUserName(userInfo YandexUserInfo) string {
	if userInfo.RealName != "" {
		return userInfo.RealName
	}
	if userInfo.DisplayName != "" {
		return userInfo.DisplayName
	}
	if userInfo.FirstName != "" && userInfo.LastName != "" {
		return userInfo.FirstName + " " + userInfo.LastName
	}
	if userInfo.FirstName != "" {
		return userInfo.FirstName
	}
	return userInfo.Login
}
