package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "shanepee.com/api/docs"
	"shanepee.com/api/infrastructure/handler"
	"shanepee.com/api/infrastructure/repository"
	"shanepee.com/api/service"
)

//	@title			Shanepee API
//	@version		0.0
//	@description	Shanepee API

// @host	localhost:8080
func main() {
	aRepo := repository.NewARepository()
	aSvc := service.NewAService(aRepo)
	aHdr := handler.NewAHandler(aSvc)
	// TODO: api versioning and prefix
	// TODO: make document dynamic
	r := gin.Default()
	r.GET("/a", aHdr.GetA)
	r.POST("/a", aHdr.CreateA)
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
