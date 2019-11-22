package responses

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RespondWithStatusOK responds with status accepted
func RespondWithStatusOK(ctx *gin.Context, message string) {
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": message})
}
