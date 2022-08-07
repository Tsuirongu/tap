package impl

import (
	"github.com/Tsuirongu/tap/service/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

// SayHello "/" health check
func SayHello(c *gin.Context) {
	c.String(200, "Hello, Geektutu")
}

// Login fake login
func Login(c *gin.Context) {
	user := &entity.User{}
	if err := c.BindJSON(user); err != nil {
		c.Error(err)
		return
	}
	if user.Username == "savagelu" {
		c.JSON(http.StatusOK, gin.H{
			"msg": "success",
		})
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"msg": "unknown user",
		})
	}
}
