package router

import (
    "go-wallet/internal/delivery/http/handler"
    "go-wallet/internal/delivery/http/middleware"
    "github.com/gin-gonic/gin"
)

func SetupRouter(
    userHandler *handler.UserHandler,
    walletHandler *handler.WalletHandler,
    jwtSecret string,
) *gin.Engine {
    router := gin.Default()

    // Public routes
    public := router.Group("/api")
    {
        public.POST("/register", userHandler.Register)
        public.POST("/login", userHandler.Login)
    }

    // Protected routes
    protected := router.Group("/api")
    protected.Use(middleware.AuthMiddleware(jwtSecret))
    {
        protected.POST("/transactions", walletHandler.CreateTransaction)
        protected.GET("/transactions", walletHandler.GetTransactions)
        protected.GET("/transactions/:id", walletHandler.GetTransaction)
        protected.PUT("/transactions/:id", walletHandler.UpdateTransaction)
        protected.DELETE("/transactions/:id", walletHandler.DeleteTransaction)
    }

    return router
}