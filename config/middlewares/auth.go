package middlewares

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func Authenticate(c *gin.Context) {
	authStr := c.GetHeader("Authorization")

	tokenStr := ""
	if strings.HasPrefix(authStr, "Bearer ") {
		tokenStr = authStr[7:]
	}

	if	tokenStr == "" {
		c.Next()
		return
	}

	//token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, e error) {
	//
	//})

	//c.Set("user", )

	//if accessToken != currentToken {
	//	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{ "message": "Unauthorized!" })
	//	return
	//}

	c.Next()
}

func RequiresLogin(c *gin.Context) {
	c.Next()
}

func RequiresAdmin(c *gin.Context) {
	c.Next()
}
