// handlers/auth.go
package handlers

import (
    "net/http"
	"api_tp/internal/models"
	"api_tp/internal/utils"
    
    "gorm.io/gorm"
)

type AuthHandler struct {
    DB *gorm.DB
}

type RegisterRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6"`
    Name     string `json:"name" validate:"required"`
}

type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
    var req RegisterRequest
    if err := utils.DecodeJSON(r, &req); err != nil {
        utils.WriteError(w, http.StatusBadRequest, "Invalid request")
        return
    }
    
    // Проверяем, существует ли пользователь
    var existingUser models.User
    if err := h.DB.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
        utils.WriteError(w, http.StatusConflict, "User already exists")
        return
    }
    
    // Хешируем пароль
    hashedPassword, err := utils.HashPassword(req.Password)
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, "Error creating user")
        return
    }
    
    // Создаем пользователя
    user := models.User{
        Email:    req.Email,
        Password: hashedPassword,
        Name:     req.Name,
    }
    
    if err := h.DB.Create(&user).Error; err != nil {
        utils.WriteError(w, http.StatusInternalServerError, "Error creating user")
        return
    }
    
    // Генерируем JWT токен
    token, err := utils.GenerateJWT(user.ID, user.Email)
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, "Error generating token")
        return
    }
    
    utils.WriteJSON(w, http.StatusCreated, map[string]interface{}{
        "token": token,
        "user":  user,
    })
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    var req LoginRequest
    if err := utils.DecodeJSON(r, &req); err != nil {
        utils.WriteError(w, http.StatusBadRequest, "Invalid request")
        return
    }
    
    // Ищем пользователя
    var user models.User
    if err := h.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
        utils.WriteError(w, http.StatusUnauthorized, "Invalid credentials")
        return
    }
    
    // Проверяем пароль
    if !utils.CheckPasswordHash(req.Password, user.Password) {
        utils.WriteError(w, http.StatusUnauthorized, "Invalid credentials")
        return
    }
    
    // Генерируем JWT токен
    token, err := utils.GenerateJWT(user.ID, user.Email)
    if err != nil {
        utils.WriteError(w, http.StatusInternalServerError, "Error generating token")
        return
    }
    
    utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
        "token": token,
        "user":  user,
    })
}