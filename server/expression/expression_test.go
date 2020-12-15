package expression

import (
	"encoding/json"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestExp(t *testing.T) {
	examlevel := LevelChoose{
		Level: 4,
		Max:   20,
		Nums:  10,
	}
	StartExamine(examlevel)

}
func TestGinSentMessage(t *testing.T) {
	examlevel := LevelChoose{
		Level: 4,
		Max:   20,
		Nums:  10,
	}
	msg := StartExamine(examlevel)
	r := gin.Default()
	var message map[string]interface{}
	r.GET("/", func(c *gin.Context) {
		json.Unmarshal(msg, &message)
		c.JSON(200, message)
		c.JSON(200, gin.H{
			"status":"ok",
		})
	})
	r.Run(":8080")
}
