package main

import (
	"log"
	"net/http"
	"os"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

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
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = app.cfg.CorsAllowOrigins
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))
	r.Use(sessions.Sessions(app.cfg.Session.CookieName, app.sessionStore))

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	humaConfig := huma.DefaultConfig("Shanepee API", "0.0.0")
	api := humagin.New(r, humaConfig)

	app.authHdr.RegisterChangePassword(api)
	app.authHdr.RegisterLogin(api)
	app.authHdr.RegisterLogout(api)
	app.authHdr.RegisterGetMe(api)
	app.authHdr.RegisterRegister(api)
	app.authHdr.RegisterCreatePasswordChangeRequests(api)

	app.userHdr.UpdateUser(api)

	app.artToyHdr.RegisterGetArtToys(api)
	app.artToyHdr.RegisterCreateArtToy(api)
	app.artToyHdr.RegisterGetArtToyById(api)
	app.artToyHdr.RegisterUpdateArtToy(api)

	if err = r.Run(); err != nil {
		log.Fatal(err)
	}
}
