package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kenmobility/grpc-project/helpers"
	"gopkg.in/go-playground/validator.v9"
)

type Config struct {
	AppEnv            string
	GRPCServerAddress string
	DatabaseHost      string `validate:"required"`
	DatabasePort      string `validate:"required"`
	DatabaseUser      string `validate:"required"`
	DatabasePassword  string `validate:"required"`
	DatabaseName      string `validate:"required"`
}

// LoadConfig loads the env variables
func LoadConfig(path string) *Config {
	var err error

	if path == "" {
		path = ".env"
	}
	if err := godotenv.Load(path); err != nil {
		log.Fatal("env config error: ", err)
	}

	configVar := Config{
		AppEnv:            helpers.Getenv("APP_ENV", "local"),
		GRPCServerAddress: os.Getenv("GRPC_SERVER_ADDRESS"),
		DatabaseHost:      os.Getenv("DATABASE_HOST"),
		DatabasePort:      os.Getenv("DATABASE_PORT"),
		DatabaseUser:      os.Getenv("DATABASE_USER"),
		DatabaseName:      os.Getenv("DATABASE_NAME"),
		DatabasePassword:  os.Getenv("DATABASE_PASSWORD"),
	}

	validate := validator.New()
	err = validate.Struct(configVar)
	if err != nil {
		log.Fatalf("env validation error: %s", err.Error())
	}

	return &configVar
}
