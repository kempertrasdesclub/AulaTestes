package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	endpointUser "github.com/kempertrasdesclub/AulaTestes/aulaInterface/gin/endpoint/user"
	"log"
	"net/http"
)

func ConfigAndStart() (err error) {
	var epUser = endpointUser.DataSource{}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.StaticFS("/static", http.Dir("./cmd/static"))
	r.GET("/local/file", func(c *gin.Context) {
		c.File("local/file.go")
	})
	r.GET("/datasource/user/:mail", epUser.UserByEmail)

	r.POST("/saveTimeLine", func(c *gin.Context) {
		var a interface{}
		c.BindJSON(&a)
		fmt.Printf("bind: %+v\n", a)
	})

	log.Println("Listening on :3000...")
	err = r.Run(":3000")
	return
}
