package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context)  {
	usrCreds := UserCredentials{}

	if err := c.Bind(&usrCreds); err != nil {
		_ = c.Error(err)
		return
	}

	user, err := Dao.FindByUsername(usrCreds.Username)
	if err == nil {
		user, err = user.SignIn(usrCreds.Password)
	}

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{ "message": "Unauthorized" })
		return
	}

	c.JSON(http.StatusOK, gin.H{ "user": user })
}

func Logout(c *gin.Context)  {
	u := c.MustGet("user").(*User)

	if err:= u.SignOut(); err !=nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{ "message": "Server error" })
		return
	}

	c.JSON(http.StatusOK, gin.H{ "message": "OK" })
}

func Get(c *gin.Context)  {
	user, err := Dao.FindById(c.Param("userId"))

	if err != nil {
		_ = c.Error(err)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{ "message": "Not found" })
		return
	}

	c.JSON(http.StatusOK, gin.H{ "user": user })
}

func List(c *gin.Context)  {
	users, err := Dao.List(nil)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{ "users": users })
}
