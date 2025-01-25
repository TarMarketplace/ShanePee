package main

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "shanepee.com/api/docs"
)

//	@title			Shanepee API
//	@version		0.0
//	@description	Shanepee API

// @host	localhost:8080
func main() {
	app, err := InitializeApp()
	if err != nil {
		log.Fatal(err)
	}
	// TODO: api versioning and prefix
	// TODO: make document dynamic
	r := gin.Default()
	r.GET("/a", app.aHdr.GetA)
	r.POST("/a", app.aHdr.CreateA)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if err = r.Run(); err != nil {
		log.Fatal(err)
	}
}
