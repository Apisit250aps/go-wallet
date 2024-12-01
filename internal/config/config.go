package config

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
	Secret string
	// กำหนดเวลาหมดอายุเป็น 24 ชั่วโมง
	ExpiryHours int
}

type ServerConfig struct {
	Port string
}

func NewConfig() *Config {
	return &Config{
		MongoDB: MongoDBConfig{
			URI:      "mongodb://localhost:27017/go_wallet",
			Database: "go_wallet",
		},
		JWT: JWTConfig{
			Secret:      "svvjJfXMAna8E6Qrer3XhzyzcWKyqrXHJvSHRRZnNmq4GgvGem634cP+vCSAOQ/DgJwBZ0M24+wbyzty00jJuw==",
			ExpiryHours: 24,
		},
		Server: ServerConfig{
			Port: ":8080",
		},
	}
}
