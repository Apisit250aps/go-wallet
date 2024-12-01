package handler

import (
    "net/http"
    "go-wallet/internal/domain"
    "go-wallet/pkg/response"
    "github.com/gin-gonic/gin"
)

type WalletHandler struct {
    walletUsecase domain.WalletUsecase
}

func NewWalletHandler(wu domain.WalletUsecase) *WalletHandler {
    return &WalletHandler{
        walletUsecase: wu,
    }
}

func (h *WalletHandler) CreateTransaction(c *gin.Context) {
    var wallet domain.Wallet
    if err := c.ShouldBindJSON(&wallet); err != nil {
        c.JSON(http.StatusBadRequest, response.Error("invalid request body"))
        return
    }

    userID := c.GetString("userID")
    wallet.UserID = userID

    err := h.walletUsecase.CreateTransaction(&wallet)
    if err != nil {
        c.JSON(http.StatusBadRequest, response.Error(err.Error()))
        return
    }

    c.JSON(http.StatusCreated, response.Success("transaction created successfully", wallet))
}

func (h *WalletHandler) GetTransactions(c *gin.Context) {
    userID := c.GetString("userID")
    transactions, err := h.walletUsecase.GetTransactions(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, response.Error("failed to get transactions"))
        return
    }

    c.JSON(http.StatusOK, response.Success("transactions retrieved successfully", transactions))
}

func (h *WalletHandler) GetTransaction(c *gin.Context) {
    id := c.Param("id")
    transaction, err := h.walletUsecase.GetTransaction(id)
    if err != nil {
        c.JSON(http.StatusNotFound, response.Error("transaction not found"))
        return
    }

    c.JSON(http.StatusOK, response.Success("transaction retrieved successfully", transaction))
}

func (h *WalletHandler) UpdateTransaction(c *gin.Context) {
    id := c.Param("id")
    var wallet domain.Wallet
    if err := c.ShouldBindJSON(&wallet); err != nil {
        c.JSON(http.StatusBadRequest, response.Error("invalid request body"))
        return
    }

    wallet.ID = id
    wallet.UserID = c.GetString("userID")

    err := h.walletUsecase.UpdateTransaction(&wallet)
    if err != nil {
        c.JSON(http.StatusBadRequest, response.Error(err.Error()))
        return
    }

    c.JSON(http.StatusOK, response.Success("transaction updated successfully", wallet))
}

func (h *WalletHandler) DeleteTransaction(c *gin.Context) {
    id := c.Param("id")
    err := h.walletUsecase.DeleteTransaction(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, response.Error(err.Error()))
        return
    }

    c.JSON(http.StatusOK, response.Success("transaction deleted successfully", nil))
}