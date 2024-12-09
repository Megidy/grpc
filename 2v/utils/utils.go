package utils

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func WriteResponse(c *gin.Context, statusCode int, data map[string]any) {
	c.JSON(statusCode, data)
}

func WriteError(c *gin.Context, statusCode int, err error) {
	slog.Info("error : ", err.Error())
	c.JSON(statusCode, gin.H{
		"error ": err.Error(),
	})
}
