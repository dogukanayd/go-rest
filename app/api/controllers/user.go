package controllers

import (
	"github.com/dogukanayd/go-rest/app/user"
	"github.com/gin-gonic/gin"
)

// Information about user
func User(ctx *gin.Context) {
	user.New().RequestCount()
	ctx.JSON(200, gin.H{"username": "donovan"})
}
