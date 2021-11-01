package thought

import (
	"github.com/gin-gonic/gin"
	"thoughts/app/controllers/thought"
)

func Routes(base string, router *gin.Engine) {
	router.Group(base + "/t").
		GET("/all", thought.GetAllThoughts).
		POST("/new", thought.AddNewThought)
}