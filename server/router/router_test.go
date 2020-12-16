package router_test

import (
	"log"
	"my-app/expression"
	request "my-app/model/request"
	"my-app/router"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGinSentMessage(t *testing.T) {
	examlevel := expression.LevelChoose{
		Level: 4,
		Max:   20,
		Nums:  10,
	}
	msg := expression.StartExamine(examlevel, false)
	r := gin.Default()
	r.GET("/practice", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"work": "done",
		})
		c.JSON(200, msg)
	})
	r.Run(":8080")
}

func TestSendHtml(t *testing.T) {
	r := gin.Default()
	r.GET("", func(c *gin.Context) {
		c.File("./index.html")
	})
	r.Run(":8080")
}

func TestCreate(t *testing.T) {
	r := gin.Default()
	r.POST("/create", router.PracticeDIY)
	r.Run(":8080")
}

type Region struct {
	conutry  string `json:"country"`
	province string `json:"province"`
}
type School struct {
	location string `json:"Location"`
	name     string `json:"Name"`
}
type Datas struct {
	names   []string `json:"names"`
	schools []School `json:"schools"`
	region  Region   `json:"Region"`
}

func TestJsonTrans(t *testing.T) {
	var transdata Datas
	r := gin.Default()
	r.POST("/see", func(c *gin.Context) {
		if c.ShouldBindJSON(&transdata) == nil {
			log.Println("what?")
		}
		c.JSON(200, gin.H{"status": "GOT"})
		log.Println(transdata)
	})
	r.Run(":8080")
}

func TestCommitTrans(t *testing.T) {
	var datafromfront request.CommitMessage
	r := gin.Default()
	r.POST("/commit", func(c *gin.Context) {
		if c.ShouldBindJSON(&datafromfront) == nil {
			log.Println(datafromfront.Expression)
			log.Println(datafromfront.Owner)
		}
		log.Println("Done")
	})
	r.Run(":8080")
}
