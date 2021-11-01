package thought

import (
	"github.com/gin-gonic/gin"
	"thoughts/app/database"
	"thoughts/app/models"
)

type thought struct {
	UserId string `json:"user_id"`
	Thought string `json:"thought"`
}

func GetAllThoughts(c *gin.Context) {

	var Thoughts []models.Thought
	var t models.Thought

	result, err := database.Connect().Query("select * from Thoughts where isprivate = ?", 0)
	if err != nil {
		println(err)
		return
	}
	for result.Next() {
		err = result.Scan(&t.ThoughtId, &t.UserId, &t.Thought, &t.IsPrivate, &t.CreatedAt)
		if err != nil {
			println(err)
			return
		}
		Thoughts = append(Thoughts, t)
	}
	c.JSON(200, Thoughts)

	err = database.Connect().Close()
	if err != nil {
		return
	}
}

func AddNewThought(c *gin.Context) {

	if c.ContentType() != "application/json" {
		c.JSON(401, "Need to send json bodies")
	}

	var t thought
	err := c.ShouldBind(&t)
	if err != nil {
		return
	}

	c.JSON(200, t)
}