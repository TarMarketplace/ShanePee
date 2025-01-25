package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "shanepee.com/api/docs"
)

//	@title			Shanepee API
//	@version		0.0
//	@description	Shanepee API

// @host	localhost:8080
func main() {
	godotenv.Load()
	app, err := InitializeApp()
	if err != nil {
		log.Fatal(err)
	}
	// TODO: api versioning and prefix
	if app.cfg.Debug != "1" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.GET("/a", app.aHdr.GetA)
	r.POST("/a", app.aHdr.CreateA)

	if app.cfg.Debug == "1" {
		r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	if err = r.Run(); err != nil {
		log.Fatal(err)
	}
}
