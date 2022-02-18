package events

import "github.com/gin-gonic/gin"

func Setup() *gin.Engine {
	engine := gin.Default()
	engine.GET("/events", getAllEventsHandler)

	return engine
}