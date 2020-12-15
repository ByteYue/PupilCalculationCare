package expression

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestExp(t *testing.T) {
	examlevel := LevelChoose{
		Symbol:  []int8{1, 2},
		Level:   4,
		Max:     20,
		Nums:    10,
		Anssize: 10,
	}
	StartExamine(examlevel, false)

}
func TestGinSentMessage(t *testing.T) {
	examlevel := LevelChoose{
		Symbol:  []int8{1, 2},
		Level:   4,
		Max:     20,
		Nums:    10,
		Anssize: 10,
	}
	msg := StartExamine(examlevel, true)
	r := gin.Default()
	//var message map[string]interface{}
	r.GET("/", func(c *gin.Context) {
		//json.Unmarshal(msg, &message)
		c.JSON(200, msg)
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	r.Run(":8080")
}
