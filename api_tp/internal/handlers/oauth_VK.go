package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"api_tp/internal/config"
	"api_tp/internal/utils"
)

// VKUserInfo представляет данные пользователя от VK
type VKUserInfo struct {
    Response []struct {
        ID        int    `json:"id"`
        FirstName string `json:"first_name"`
        LastName  string `json:"last_name"`
        Email     string `json:"email"`
        Photo     string `json:"photo_200"`
    } `json:"response"`
}

// VKEmailResponse представляет ответ с email от VK
type VKEmailResponse struct {
    Email string `json:"email"`
}

// VKLogin initiates VK OAuth flow
func (h *OAuthHandler) VKLogin(w http.ResponseWriter, r *http.Request) {
    url := config.VKOAuthConfig.AuthCodeURL("state")
    http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// VKCallback handles VK OAuth callback
func (h *OAuthHandler) VKCallback(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    
    token, err := config.VKOAuthConfig.Exchange(r.Context(), code)
    if err != nil {
        utils.WriteError(w, http.StatusBadRequest, "Failed to exchange token: "+err.Error())
        return
    }
    
    // VK не возвращает email в основном токене, нужно получить его отдельно
    email, err := h.getVKEmail(token.AccessToken)
    if err != nil {
        utils.WriteError(w, http.StatusBadRequest, "Failed to get email from VK: "+err.Error())
        return
    }
    
    client := config.VKOAuthConfig.Client(r.Context(), token)
    
    // Получаем основную информацию о пользователе
    userInfoURL := fmt.Sprintf("https://api.vk.com/method/users.get?fields=photo_200,email&v=5.131&access_token=%s", token.AccessToken)
    resp, err := client.Get(userInfoURL)
    if err != nil {
        utils.WriteError(w, http.StatusBadRequest, "Failed to get user info: "+err.Error())
        return
    }
    defer resp.Body.Close()
    
    var vkUserInfo VKUserInfo
    if err := json.NewDecoder(resp.Body).Decode(&vkUserInfo); err != nil {
        utils.WriteError(w, http.StatusBadRequest, "Failed to decode user info: "+err.Error())
        return
    }
    
    if len(vkUserInfo.Response) == 0 {
        utils.WriteError(w, http.StatusBadRequest, "No user data received from VK")
        return
    }
    
    vkUser := vkUserInfo.Response[0]
    userID := fmt.Sprintf("%d", vkUser.ID)
    name := vkUser.FirstName + " " + vkUser.LastName
    
    // Используем email из отдельного запроса
    if email == "" && vkUser.Email != "" {
        email = vkUser.Email
    }
    
    // Если email все еще пустой, создаем временный
    if email == "" {
        email = fmt.Sprintf("vk_%s@temp.vk", userID)
    }
    
    // Создаем или находим пользователя
    user, err := h.findOrCreateOAuthUser("vk", userID, email, name, token)
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

// getVKEmail получает email из VK OAuth
func (h *OAuthHandler) getVKEmail(accessToken string) (string, error) {
    // VK возвращает email в ответе на запрос токена, но если его нет,
    // можно попробовать получить через API
    emailURL := fmt.Sprintf("https://api.vk.com/method/account.getProfileInfo?v=5.131&access_token=%s", accessToken)
    
    resp, err := http.Get(emailURL)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    
    var emailResp struct {
        Response struct {
            Email string `json:"email"`
        } `json:"response"`
    }
    
    if err := json.NewDecoder(resp.Body).Decode(&emailResp); err != nil {
        return "", err
    }
    
    return emailResp.Response.Email, nil
}