package utils

import "github.com/gin-gonic/gin"

func ErrorResponse(c *gin.Context, err error) {
	c.JSON(100001, gin.H{
		"error": err.Error(),
	})
}
