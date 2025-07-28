package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WriteSuccessOutput[T any](c *gin.Context, output *T) {
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    output,
	})
}

func WriteCreatedOutput[T any](c *gin.Context, output *T) {
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    output,
	})
}
