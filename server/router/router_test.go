package router_test

import (
	"my-app/expression"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGinSentMessage(t *testing.T) {
	examlevel := expression.LevelChoose{
		Level: 4,
		Max:   20,
		Nums:  10,
	}
	msg := expression.StartExamine(examlevel)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, msg)
	})
	r.Run(":8080")
}

