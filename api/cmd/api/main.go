package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"shanepee.com/api/docs"
)

// @title       Shanepee API
// @version     0.0
// @description Shanepee API
//
// @host        localhost:8080
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
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = app.cfg.CorsAllowOrigins
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))
	r.Use(sessions.Sessions(app.cfg.Session.CookieName, app.sessionStore))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	v1 := r.Group("v1")

	v1.POST("/auth/change-password", app.authHdr.ChangePassword)
	v1.POST("/auth/login", app.authHdr.Login)
	v1.POST("/auth/logout", app.authHdr.Logout)
	v1.GET("/auth/me", app.authHdr.GetMe)
	v1.POST("/auth/register", app.authHdr.Register)
	v1.POST("/auth/password-change-requests", app.authHdr.CreatePasswordChangeRequests)

	v1.PATCH("/user", app.userHdr.UpdateUser)

	v1.GET("/art-toy", app.artToyHdr.GetArtToys)
	v1.POST("/art-toy", app.artToyHdr.CreateArtToy)
	v1.GET("/art-toy/:id", app.artToyHdr.GetArtToyById)
	v1.PUT("/art-toy/:id", app.artToyHdr.UpdateArtToy)

	if app.cfg.Debug == "1" {
		v1.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	if err = r.Run(); err != nil {
		log.Fatal(err)
	}
}
