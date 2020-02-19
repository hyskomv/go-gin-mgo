package config

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

var (
	AppEnv = "development"
	Port   = "7745"
	DBUri  = ""
)

func init() {
	gin.SetMode(gin.ReleaseMode)

	// environments
	err := godotenv.Load()
	if err != nil {
		println(err.Error())
	}

	env := os.Getenv("APP_ENV")
	if env != "" {
		AppEnv = env
	} else {
		err := os.Setenv("APP_ENV", AppEnv)
		if err != nil {
			panic(err.Error())
		}
	}

	// parse DB URI
	DBUri = os.Getenv("MONGODB_URI")
	if DBUri == "" {
		panic("No mongo connection string. Set MONGODB_URI environment variable.")
	}

	// app port
	port := os.Getenv("PORT")
	if port != "" {
		Port = port
	}
}
