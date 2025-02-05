package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"shanepee.com/api/docs"
)

//	@title			Shanepee API
//	@version		0.0
//	@description	Shanepee API

// @host	localhost:8080
func main() {
	if err := godotenv.Load(); err != nil {
		if _, ok := err.(*os.PathError); ok {
			log.Print(".env not found, skipping")
		} else {
			log.Fatal(err)
		}
	}
	app, err := InitializeApp()
	if err != nil {
		log.Fatal(err)
	}
	if app.cfg.Debug != "1" {
		gin.SetMode(gin.ReleaseMode)
	}
	docs.SwaggerInfo.Host = app.cfg.ServerUrl

	r := gin.Default()
	r.Use(cors.Default())

	v1 := r.Group("v1")

	v1.GET("/a", app.aHdr.GetA)
	v1.GET("/a/:id", app.aHdr.GetAById)
	v1.POST("/a", app.aHdr.CreateA)
	v1.PATCH("/a/:id", app.aHdr.UpdateA)
	v1.DELETE("/a/:id", app.aHdr.DeleteA)

	v1.POST("/auth/register", app.authHdr.Register)

	v1.GET("/user", app.userHdr.GetUsers)
	v1.GET("/user/:id", app.userHdr.GetUser)
	v1.PATCH("/user/:id", app.userHdr.UpdateUser)

	if app.cfg.Debug == "1" {
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	if err = r.Run(); err != nil {
		log.Fatal(err)
	}
}
