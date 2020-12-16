package expression

import (
	"log"
	request "my-app/model/request"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCommitTrans(t *testing.T) {
	var datafromfront request.CommitMessage
	r := gin.Default()
	r.POST("/commit", func(c *gin.Context) {
		if c.ShouldBindJSON(&datafromfront) == nil {
			log.Println(datafromfront.Expression)
			log.Println(datafromfront.Owner)
			c.JSON(200,gin.H{
				"status":"GOT",
			})
		}
		log.Println("Done")
	})
	r.Run(":8080")
}

func TestGenerateJSONToFront(t *testing.T) {
	var transdata ExJSON
	allExpressions := make([]Ex, 0)
	opers := "+-*/"
	ch := LevelChoose{
		Symbol:  []int8{1, 2},
		Level:   4,
		Max:     20,
		Nums:    10,
		Anssize: 10,
	}
	ExpressionGenerate(ch.Level, ch.Max, ch.Nums, &allExpressions, &opers)
	transdata.Expressions = allExpressions
	r := gin.Default()
	r.GET("/see", func(c *gin.Context) {
		c.JSON(200, transdata)
	})
	r.Run(":8080")
}

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
