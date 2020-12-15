package router

import (
	"db"
	"strings"

	"github.com/gin-gonic/gin"
)

type Server struct {
	queries *db.Queries
	router  *gin.Engine
}

func NewServer(queries *db.Queries) *Server {
	server := &Server{queries: queries}

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		//将首页HTML传给CLIENT
		//TODO 等前端把资源文件发过来
	})
	//router.POST("./accounts", server.CreateAccount)
	router.GET("/login", server.LogIn)
	router.GET("/register", server.Register)
	router.GET("/practice", PracticeGenerate)
	router.POST("/DIY", PracticeDIY)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func InitMessageFilename(Owner string) string {
	s1 := "_message.txt"
	var build strings.Builder
	build.WriteString(Owner)
	build.WriteString(s1)
	return build.String()
}

func InitMistakesFilename(Owner string) string {
	s1 := "_mistake.txt"
	var build strings.Builder
	build.WriteString(Owner)
	build.WriteString(s1)
	return build.String()
}
