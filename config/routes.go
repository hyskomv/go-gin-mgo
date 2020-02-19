package config

import (
	"github.com/gin-gonic/gin"
	"go-gin-mgo/app/user"
	"go-gin-mgo/config/middlewares"
)

func Routes(app *gin.Engine) {
	app.POST("/login", user.Login)

	app.Use(middlewares.Authenticate, middlewares.RequiresLogin)
	{
		app.GET("/logout", middlewares.RequiresLogin, user.Logout)
		app.GET("/users", middlewares.RequiresAdmin, user.List)
		app.GET("/users/:userId", middlewares.RequiresLogin, user.Get)
	}

	app.Use(middlewares.ErrorHandler)
}
