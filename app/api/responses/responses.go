package responses

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RespondWithStatusOK responds with status accepted
func RespondWithStatusOK(ctx *gin.Context, message string) {
	ctx.Writer.WriteHeader(http.StatusOK)
	ctx.JSON(http.StatusOK, gin.H {"status": http.StatusOK, "message": message} )
}

// RespondWithBadRequest responds with status bad request
func RespondWithBadRequest(ctx *gin.Context, message string)  {
	ctx.Writer.WriteHeader(http.StatusBadRequest)
	ctx.JSON(http.StatusBadRequest, gin.H {"status": http.StatusBadRequest, "message": message} )
}
