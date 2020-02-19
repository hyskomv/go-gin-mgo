package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-mgo/config"
	"go-gin-mgo/database"
	"log"
)

func init()  {
	// init database
	database.Connect(config.DBUri)
	database.InitModels()
}

func main()  {
	env := config.AppEnv
	port := config.Port
	url := "http://localhost:" + port

	app := gin.Default()

	// router
	config.Routes(app)

	// start
	log.Printf("App is running at %s in %s mode", url, env)
	//log.Printf("API explorer is running at %s/explorer", url)

	if err := app.Run(":" + port); err != nil {
		panic(err.Error())
	}
}
