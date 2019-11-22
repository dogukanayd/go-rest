package controllers

import (
	"github.com/gin-gonic/gin"
)

// Information about user
func User(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"username": "dogukan"})
}
