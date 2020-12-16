package router

import (
	"db"
	"my-app/middleware"
	request "my-app/model/request"
	"strings"

	"github.com/gin-gonic/gin"
)

//Server JSON
type Server struct {
	queries *db.Queries
	router  *gin.Engine
}

//NewServer newServer
func NewServer(queries *db.Queries) *Server {
	server := &Server{queries: queries}

	router := gin.Default()
	//config := cors.DefaultConfig()
	router.Use(middleware.Cors())
	//router.GET("/", func(c *gin.Context) {
	//将首页HTML传给CLIENT
	//TODO 等前端把资源文件发过来
	//})
	//router.POST("./accounts", server.CreateAccount)
	//router.Use(middleware())

	router.POST("/login", server.Register)
	router.GET("/practice", PracticeGenerate)
	router.GET("/test", PracticeGenerate)
	router.POST("/create", PracticeDIY)
	router.POST("/commit", func(c *gin.Context) {
		var ch request.CommitMessage
		if c.ShouldBindJSON(&ch) != nil {
			c.JSON(400, nil)
		}
		name := server.Getfilename(ch.Owner, c)
		exps := request.ItemToString(&ch.Expression)
		//TODO  邱邱人写文件操作
		request.SaveToFile(exps, name)
	})
	server.router = router
	return server
}

//Start start router
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

//errorResponse error handles
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

//InitMessageFilename init filename
func InitMessageFilename(Owner string) string {
	s1 := "_message.txt"
	var build strings.Builder
	build.WriteString(Owner)
	build.WriteString(s1)
	return build.String()
}

//InitMistakesFilename init mistake file
func InitMistakesFilename(Owner string) string {
	s1 := "_mistake.txt"
	var build strings.Builder
	build.WriteString(Owner)
	build.WriteString(s1)
	return build.String()
}
