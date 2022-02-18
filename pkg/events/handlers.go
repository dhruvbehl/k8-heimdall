package events

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getAllEventsHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, EventList)
}