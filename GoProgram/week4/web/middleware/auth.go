package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var loggedInUsers = make(map[string]bool)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, err := c.Cookie("username")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未授权访问，请先登录"})
			c.Abort()
			return
		}

		if !loggedInUsers[username] {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "登录已过期，请重新登录"})
			c.Abort()
			return
		}

		c.Set("username", username)
		c.Next()
	}
}

func LoginUser(username string, c *gin.Context) {
	loggedInUsers[username] = true
	c.SetCookie("username", username, 3600, "/", "", false, true)
	c.SetCookie("is_logged_in", "true", 3600, "/", "", false, false)
}

func LogoutUser(username string, c *gin.Context) {
	delete(loggedInUsers, username)
	c.SetCookie("username", "", -1, "/", "", false, true)
	c.SetCookie("is_logged_in", "", -1, "/", "", false, false)
}

func GetCurrentUser(c *gin.Context) string {
	username, _ := c.Get("username")
	return username.(string)
}
