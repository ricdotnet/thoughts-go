package api

import (
	"github.com/gin-gonic/gin"
	"thoughts/app/api/thought"
)

func Routes(router *gin.Engine) {
	main := router.Group("/api")
	main.GET("/", invalid)

	thought.Routes(main.BasePath(), router)
}

func invalid(c *gin.Context) {
	c.JSON(204, "No content")
}