package config

import (
	"log"
	"os"

	"github.com/emilpoppler134/api-template/internal/utils"
	"github.com/joho/godotenv"
)

type EnviromentVariable struct {
	key   string
	value string
}

type Config struct {
	Port        int
	DatabaseDSN string
}

func Load() Config {
	godotenv.Load()

	var PORT int = getEnviromentVariable("PORT").asInteger()

	var DATABASE_USERNAME string = getEnviromentVariable("DATABASE_USERNAME").asString()
	var DATABASE_PASSWORD string = getEnviromentVariable("DATABASE_PASSWORD").asString()
	var DATABASE_HOST string = getEnviromentVariable("DATABASE_HOST").asString()
	var DATABASE_PORT int = getEnviromentVariable("DATABASE_PORT").asInteger()
	var DATABASE_NAME string = getEnviromentVariable("DATABASE_NAME").asString()
	var DATABASE_SSLMODE string = getEnviromentVariable("DATABASE_SSLMODE").asString()
	var DATABASE_TIMEZONE string = getEnviromentVariable("DATABASE_TIMEZONE").asString()

	var dsn string = utils.FormatDSN(DATABASE_USERNAME, DATABASE_PASSWORD, DATABASE_HOST, DATABASE_PORT, DATABASE_NAME, DATABASE_SSLMODE, DATABASE_TIMEZONE)

	return Config{
		Port:        PORT,
		DatabaseDSN: dsn,
	}
}

func getEnviromentVariable(key string) EnviromentVariable {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("Error loading enviroment variables, %s is required", key)
	}
	return EnviromentVariable{key, value}
}

func (variable EnviromentVariable) asString() string {
	return variable.value
}

func (variable EnviromentVariable) asInteger() int {
	value, err := utils.ParseInt(variable.value)
	if err != nil {
		log.Fatalf("Error loading enviroment variables, %s must be an integer", variable.key)
	}
	return value
}
