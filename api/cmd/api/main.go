package main

import (
	"log"
	"os"

	"github.com/gin-contrib/sessions"
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
	r.Use(sessions.Sessions(app.cfg.Session.CookieName, app.sessionStore))

	v1 := r.Group("v1")

	v1.POST("/auth/register", app.authHdr.Register)
	v1.POST("/auth/login", app.authHdr.Login)
	v1.POST("/auth/logout", app.authHdr.Logout)
	v1.POST("/auth/password-change-requests", app.authHdr.CreatePasswordChangeRequests)
	v1.POST("/auth/change-password", app.authHdr.ChangePassword)
	v1.GET("/auth/me", app.authHdr.GetMe)

	v1.PATCH("/user", app.userHdr.UpdateUser)

	if app.cfg.Debug == "1" {
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	if err = r.Run(); err != nil {
		log.Fatal(err)
	}
}
