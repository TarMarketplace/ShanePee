package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

  "github.com/gin-contrib/cors"
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

	r := gin.Default()
  r.Use(cors.Default())
	v1 := r.Group("v1")
	v1.GET("/a", app.aHdr.GetA)
	v1.POST("/a", app.aHdr.CreateA)
	v1.POST("/auth/register", app.authHdr.Register)

	if app.cfg.Debug == "1" {
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	if err = r.Run(); err != nil {
		log.Fatal(err)
	}
}
