package config

import (
    "log"
    "os"
    "strconv"

    "github.com/joho/godotenv"
)

type Config struct {
    MongoDB MongoDBConfig
    JWT     JWTConfig
    Server  ServerConfig
}

type MongoDBConfig struct {
    URI      string
    Database string
}

type JWTConfig struct {
    Secret      string
    ExpiryHours int
}

type ServerConfig struct {
    Port string
}

func NewConfig() *Config {
    err := godotenv.Load()
    if err != nil {
        log.Printf("Warning: .env file not found, using default values")
    }

    // MongoDB Config
    mongoURI := getEnv("MONGODB_URI", "mongodb://localhost:27017")
    mongoDatabase := getEnv("MONGODB_DATABASE", "go_wallet")

    // JWT Config
    jwtSecret := getEnv("JWT_SECRET", "your-secret-key")
    jwtExpiryHours := getEnvAsInt("JWT_EXPIRY_HOURS", 24)

    // Server Config
    serverPort := getEnv("SERVER_PORT", ":8080")

    return &Config{
        MongoDB: MongoDBConfig{
            URI:      mongoURI,
            Database: mongoDatabase,
        },
        JWT: JWTConfig{
            Secret:      jwtSecret,
            ExpiryHours: jwtExpiryHours,
        },
        Server: ServerConfig{
            Port: serverPort,
        },
    }
}

// Helper function to get env variable with default value
func getEnv(key string, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}

// Helper function to get env variable as integer with default value
func getEnvAsInt(key string, defaultValue int) int {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }

    intValue, err := strconv.Atoi(value)
    if err != nil {
        return defaultValue
    }

    return intValue
}