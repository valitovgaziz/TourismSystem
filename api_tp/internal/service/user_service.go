package service

import (
	"errors"
	"api_tp/internal/models"
	"api_tp/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
    userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
    return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(req *models.CreateUserRequest) (*models.UserResponse, error) {
    // Проверяем существует ли пользователь с таким email
    existingUser, err := s.userRepo.FindByEmail(req.Email)
    // Проверяем как на nil, так на пустой ID
    if existingUser != nil && existingUser.ID != 0 {
        return nil, errors.New("user with this email already exists")
    }

    // Хешируем пароль
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    user := &models.User{
        Name:     req.Name,
        Email:    req.Email,
        Password: string(hashedPassword),
    }

    if err := s.userRepo.Create(user); err != nil {
        return nil, err
    }

    return s.toUserResponse(user), nil
}

func (s *UserService) GetUserByID(id uint) (*models.UserResponse, error) {
    user, err := s.userRepo.FindByID(id)
    if err != nil {
        return nil, errors.New("user not found")
    }

    return s.toUserResponse(user), nil
}

func (s *UserService) GetAllUsers() ([]models.UserResponse, error) {
    users, err := s.userRepo.FindAll()
    if err != nil {
        return nil, err
    }

    var responses []models.UserResponse
    for _, user := range users {
        responses = append(responses, *s.toUserResponse(&user))
    }

    return responses, nil
}

func (s *UserService) toUserResponse(user *models.User) *models.UserResponse {
    return &models.UserResponse{
        ID:        user.ID,
        CreatedAt: user.CreatedAt,
        UpdatedAt: user.UpdatedAt,
        Name:      user.Name,
        Email:     user.Email,
    }
}