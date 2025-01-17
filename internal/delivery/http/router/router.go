package router

import (
	"go-wallet/internal/delivery/http/handler"
	"go-wallet/internal/delivery/http/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(
    userHandler *handler.UserHandler,
    walletHandler *handler.WalletHandler,
    jwtSecret string,
) *gin.Engine {
    router := gin.Default()
    router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                   // Allow frontend origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, // Include OPTIONS for preflight
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Include required headers
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

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