package controllers

import (
	"github.com/dogukanayd/go-rest/app/api/responses"
	"github.com/dogukanayd/go-rest/app/health"
	"github.com/gin-gonic/gin"
	"log"
)

// Health handles health requests
func Health(ctx *gin.Context) {
	h := health.New().Check()

	if h != nil {
		log.Println(h)
		responses.RespondWithBadRequest(ctx, "database connection failed")
	}

	responses.RespondWithStatusOK(ctx, "application is running and db connection success")
}
