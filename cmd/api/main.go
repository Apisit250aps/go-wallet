package main

import (
    "log"
    "go-wallet/internal/config"
    "go-wallet/internal/delivery/http/handler"
    "go-wallet/internal/delivery/http/router"
    "go-wallet/internal/repository/mongodb"
    "go-wallet/internal/usecase"
    "go-wallet/pkg/database/mongdb"
)

func main() {
    // Load config
    cfg := config.NewConfig()

    // Connect to MongoDB
    client, err := mongdb.NewMongoDBConnection(cfg.MongoDB.URI)
    if err != nil {
        log.Fatal(err)
    }
    db := client.Database(cfg.MongoDB.Database)

    // Initialize repositories
    userRepo := mongodb.NewUserRepository(db)
    walletRepo := mongodb.NewWalletRepository(db)

    // Initialize usecases
    userUsecase := usecase.NewUserUsecase(userRepo, cfg.JWT.Secret, cfg.JWT.ExpiryHours)
    walletUsecase := usecase.NewWalletUsecase(walletRepo)

    // Initialize handlers
    userHandler := handler.NewUserHandler(userUsecase)
    walletHandler := handler.NewWalletHandler(walletUsecase)

    // Setup router
    r := router.SetupRouter(userHandler, walletHandler, cfg.JWT.Secret)

    // Start server
    log.Printf("Server starting on port %s\n", cfg.Server.Port)
    if err := r.Run(cfg.Server.Port); err != nil {
        log.Fatal(err)
    }
}