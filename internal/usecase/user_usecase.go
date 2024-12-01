package usecase

import (
    "errors"
    "go-wallet/internal/domain"
    "go-wallet/pkg/auth"
)

type userUsecase struct {
    userRepo    domain.UserRepository
    jwtSecret   string
    jwtExpiry   int
}

func NewUserUsecase(ur domain.UserRepository, secret string, expiry int) domain.UserUsecase {
    return &userUsecase{
        userRepo:    ur,
        jwtSecret:   secret,
        jwtExpiry:   expiry,
    }
}

func (u *userUsecase) Register(user *domain.User) error {
    // Check if username already exists
    existingUser, _ := u.userRepo.GetByUsername(user.Username)
    if existingUser != nil {
        return errors.New("username already exists")
    }

    // Hash password
    hashedPassword, err := auth.HashPassword(user.Password)
    if err != nil {
        return err
    }
    user.Password = hashedPassword

    return u.userRepo.Create(user)
}

func (u *userUsecase) Login(username, password string) (string, error) {
    user, err := u.userRepo.GetByUsername(username)
    if err != nil {
        return "", errors.New("invalid credentials")
    }

    if !auth.CheckPasswordHash(password, user.Password) {
        return "", errors.New("invalid credentials")
    }

    // Generate JWT token
    token, err := auth.GenerateToken(user.ID, u.jwtSecret, u.jwtExpiry)
    if err != nil {
        return "", err
    }

    return token, nil
}

func (u *userUsecase) GetUserByID(id string) (*domain.User, error) {
    return u.userRepo.GetByID(id)
}