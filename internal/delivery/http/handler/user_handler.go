package handler

import (
    "net/http"
    "go-wallet/internal/domain"
    "go-wallet/pkg/response"
    "github.com/gin-gonic/gin"
)

type UserHandler struct {
    userUsecase domain.UserUsecase
}

func NewUserHandler(uu domain.UserUsecase) *UserHandler {
    return &UserHandler{
        userUsecase: uu,
    }
}

func (h *UserHandler) Register(c *gin.Context) {
    var user domain.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, response.Error("invalid request body"))
        return
    }

    err := h.userUsecase.Register(&user)
    if err != nil {
        c.JSON(http.StatusBadRequest, response.Error(err.Error()))
        return
    }

    c.JSON(http.StatusCreated, response.Success("user registered successfully", user))
}

func (h *UserHandler) Login(c *gin.Context) {
    var credentials struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&credentials); err != nil {
        c.JSON(http.StatusBadRequest, response.Error("invalid request body"))
        return
    }

    token, err := h.userUsecase.Login(credentials.Username, credentials.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, response.Error(err.Error()))
        return
    }

    c.JSON(http.StatusOK, response.Success("login successful", gin.H{"token": token}))
}