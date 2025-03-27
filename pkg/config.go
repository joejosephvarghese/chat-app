package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port       string `mapstructure:"PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBDialect  string `mapstructure:"DB_DIALECT"`
	JWTKey     string `mapstructure:"JWT_KEY"`
}

func LoadConfig() (Config, error) {
	var config Config

	// Check if the .env file exists
	if _, err := os.Stat(".env"); err != nil {
		if os.IsNotExist(err) {
			fmt.Println(".env file does not exist, using system environment variables.")
		} else {
			fmt.Println("Error checking .env file:", err)
			return config, err
		}
	} else {
		fmt.Println(".env file found.")
	}

	// Load environment variables
	viper.SetConfigFile(".env")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv() // Load from environment variables

	// Read .env file if it exists
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Warning: Failed to read .env file, using only system environment variables.")
	}

	// Unmarshal to Config struct
	if err := viper.Unmarshal(&config); err != nil {
		return config, err
	}

	// Print the loaded configuration (debugging)
	fmt.Printf("Loaded Config: %+v\n", config)

	return config, nil
}
