package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`

	JWTSecret      string `mapstructure:"JWT_SECRET"`
	JWTExpireHours int    `mapstructure:"JWT_EXPIRE_HOURS"`

	DeepSeekAPIKey  string `mapstructure:"DEEPSEEK_API_KEY"`
	DeepSeekBaseURL string `mapstructure:"DEEPSEEK_BASE_URL"`

	TongyiAPIKey  string `mapstructure:"TONGYI_API_KEY"`
	TongyiBaseURL string `mapstructure:"TONGYI_BASE_URL"`

	ServerPort string `mapstructure:"SERVER_PORT"`
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: .env file not found, using environment variables: %v", err)
	}

	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "3306")
	viper.SetDefault("DB_USER", "root")
	viper.SetDefault("DB_PASSWORD", "root123")
	viper.SetDefault("DB_NAME", "wordbook")
	viper.SetDefault("JWT_SECRET", "default-secret")
	viper.SetDefault("JWT_EXPIRE_HOURS", 24)
	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("DEEPSEEK_BASE_URL", "https://api.deepseek.com")
	viper.SetDefault("TONGYI_BASE_URL", "https://dashscope.aliyuncs.com/compatible-mode")

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}
}

func GetDSN() string {
	return AppConfig.DBUser + ":" + AppConfig.DBPassword +
		"@tcp(" + AppConfig.DBHost + ":" + AppConfig.DBPort + ")/" +
		AppConfig.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
}
