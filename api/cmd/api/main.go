package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humagin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"shanepee.com/api/infrastructure/handler"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
)

var cmd = flag.String("command", "server", "command to run")
var openApiOutDir = flag.String("output", "./docs", "output openapi file")

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
	r.Use(handler.GetUserSession())

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	AddSwagger(r)
	AddSpotlight(r)
	AddScalar(r)

	humaConfig := huma.DefaultConfig("Shanepee API", "0.0.0")
	humaConfig.DocsPath = ""
	humaConfig.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"sessionId": {
			Type:        "apiKey",
			Description: "session from login route",
			Name:        "session",
			In:          "cookie",
		},
	}
	api := humagin.New(r, humaConfig)

	app.authHdr.RegisterChangePassword(api)
	app.authHdr.RegisterCreatePasswordResetRequests(api)
	app.authHdr.RegisterLogin(api)
	app.authHdr.RegisterLogout(api)
	app.authHdr.RegisterGetMe(api)
	app.authHdr.RegisterRegister(api)
	app.authHdr.RegisterResetPassword(api)

	app.userHdr.UpdateUser(api)

	app.artToyHdr.RegisterGetArtToys(api)
	app.artToyHdr.RegisterGetMyArtToys(api)
	app.artToyHdr.RegisterCreateArtToy(api)
	app.artToyHdr.RegisterGetArtToyByID(api)
	app.artToyHdr.RegisterUpdateArtToy(api)
	app.artToyHdr.RegisterDeleteArtToy(api)
	app.artToyHdr.RegisterSearchArtToys(api)

	app.reviewHdr.RegisterCreateReview(api)
	app.reviewHdr.RegisterGetReview(api)
	app.reviewHdr.RegisterGetSellerRating(api)
	app.reviewHdr.RegisterUpdateReview(api)
	app.reviewHdr.RegisterDeleteReview(api)

	app.cartHdr.RegisterAddItemToCart(api)
	app.cartHdr.RegisterRemoveItemFromCart(api)
	app.cartHdr.RegisterClearItemsFromCart(api)
	app.cartHdr.RegisterGetCart(api)
	app.cartHdr.RegisterCheckout(api)

	app.orderHdr.RegisterGetOrdersByStatus(api)
	app.orderHdr.RegisterGetOrdersOfSeller(api)
	app.orderHdr.RegisterGetOrderOfSeller(api)
	app.orderHdr.RegisterGetOrdersOfBuyer(api)
	app.orderHdr.RegisterCompleteOrder(api)
	app.orderHdr.RegisterUpdateOrder(api)

	flag.Parse()

	if cmd == nil {
		log.Fatal("Missing command")
	} else if *cmd == "server" {
		if err = r.Run(); err != nil {
			log.Fatal(err)
		}
	} else if *cmd == "openapi" {
		jsonData, err := api.OpenAPI().MarshalJSON()
		if err != nil {
			log.Fatal(err)
		}
		var remarshal any
		err = json.Unmarshal(jsonData, &remarshal)
		if err != nil {
			log.Fatal(err)
		}
		prettied, err := json.MarshalIndent(remarshal, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		if openApiOutDir == nil {
			log.Fatal("no output directory provide")
		}
		outDir := *openApiOutDir
		oaiJsonPath := path.Join(outDir, "openapi.json")
		jsonDocs, err := os.Create(oaiJsonPath)
		if err != nil {
			log.Fatal(err)
		}
		if openApiOutDir == nil {
			log.Fatal("unable to create json docs file")
		}
		defer jsonDocs.Close()

		_, err = jsonDocs.Write(prettied)
		if err != nil {
			log.Fatal(err)
		}
	}
}
