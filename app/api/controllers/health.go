package controllers

import (
	"github.com/dogukanayd/go-rest/app/api/responses"
	"github.com/gin-gonic/gin"
)

// Health handles health requests
func Health(ctx *gin.Context)  {
	responses.RespondWithStatusOK(ctx, "application is running")
}
