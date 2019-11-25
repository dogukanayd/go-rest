package controllers

import (
	"github.com/dogukanayd/go-rest/app/api/responses"
	"github.com/dogukanayd/go-rest/app/healt"
	"github.com/gin-gonic/gin"
)

// Health handles health requests
func Health(ctx *gin.Context) {
	h := healt.New().Check()

	if h != nil {
		responses.RespondWithBadRequest(ctx, "database connection failed")
	}

	responses.RespondWithStatusOK(ctx, "application is running and db connection success")
}
