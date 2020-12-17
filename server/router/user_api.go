package router

import (
	"log"
	"my-app/expression"
	"strconv"

	//"db"
	"github.com/gin-gonic/gin"
)

//DataRestored 缓存
var DataRestored expression.LevelChoose

//PracticeGenerate 生成Practice给client
func PracticeGenerate(c *gin.Context) {
	log.Println(DataRestored)
	msg := expression.StartExamine(DataRestored, true)
	log.Println(msg)
	//c.String(200, "expression传不过来???")
	c.JSON(200, msg)
}

//PracticeDIY 生成DIY给client
func PracticeDIY(c *gin.Context) {
	var ch expression.LevelChoose
	//DataRestored.See = c.PostForm("see")
	symbol := c.PostFormMap("symbol")
	dnum := c.PostForm("dnum")
	dsize := c.PostForm("dsize")
	puznum := c.PostForm("puznum")
	anssize := c.PostForm("anssize")
	log.Println(symbol)
	log.Println(dnum)
	ch.Dnum, _ = strconv.Atoi(dnum)
	ch.Dsize, _ = strconv.Atoi(dsize)
	ch.Puznum, _ = strconv.Atoi(puznum)
	ch.Anssize, _ = strconv.Atoi(anssize)
	ch.Symbol = make([]int8, len(symbol))
	for key, elem := range symbol {
		temp, _ := strconv.Atoi(elem)
		i, _ := strconv.Atoi(key)
		ch.Symbol[i] = int8(temp)
	}
	log.Println(ch)
	DataRestored = ch
	log.Println("---切割线---")
	log.Println(DataRestored)
	//msg := expression.StartExamine(ch, true)
	//DataRestored = ch
	c.JSON(200, nil)
}
