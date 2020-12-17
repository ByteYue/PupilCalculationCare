package expression

import (
	"log"
	request "my-app/model/request"
	"strconv"
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
			c.JSON(200, gin.H{
				"status": "GOT",
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
		//Symbol:  1111,
		Dnum:    4,
		Dsize:   20,
		Puznum:  10,
		Anssize: 10,
	}
	ExpressionGenerate(ch.Dnum, ch.Dsize, ch.Puznum, &allExpressions, &opers)
	transdata.Expressions = allExpressions
	r := gin.Default()
	r.GET("/see", func(c *gin.Context) {
		c.JSON(200, transdata)
	})
	r.Run(":8080")
}

func TestExp(t *testing.T) {
	examlevel := LevelChoose{
		//Symbol:  1111,
		Dnum:    4,
		Dsize:   20,
		Puznum:  10,
		Anssize: 10,
	}
	msg := StartExamine(examlevel, false)
	log.Println(msg)
}
func TestGinSentMessage(t *testing.T) {
	//var examlevel LevelChoose
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		var ch LevelChoose
		symbol := c.PostFormMap("symbol")
		level := c.PostForm("level")
		max := c.PostForm("max")
		nums := c.PostForm("nums")
		anssize := c.PostForm("anssize")
		log.Println(symbol)
		log.Println(level)
		ch.Dnum, _ = strconv.Atoi(level)
		ch.Dsize, _ = strconv.Atoi(max)
		ch.Puznum, _ = strconv.Atoi(nums)
		ch.Anssize, _ = strconv.Atoi(anssize)
		for key, elem := range symbol {
			temp, _ := strconv.Atoi(elem)
			i, _ := strconv.Atoi(key)
			ch.Symbol[i] = int8(temp)
		}
		log.Println(ch)
		msg := StartExamine(ch, true)
		c.JSON(200, msg)
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
	r.Run(":8080")
}
