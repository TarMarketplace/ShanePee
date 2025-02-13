package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddSwagger(router *gin.Engine) {
	router.GET("/docs/swagger", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`<!DOCTYPE html>
		<html lang="en">
		<head>
		  <meta charset="utf-8" />
		  <meta name="viewport" content="width=device-width, initial-scale=1" />
		  <meta name="description" content="SwaggerUI" />
		  <title>SwaggerUI</title>
		  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui.css" />
		</head>
		<body>
		<div id="swagger-ui"></div>
		<script src="https://unpkg.com/swagger-ui-dist@5.11.0/swagger-ui-bundle.js" crossorigin></script>
		<script>
		  window.onload = () => {
			window.ui = SwaggerUIBundle({
			  url: '/openapi.json',
			  dom_id: '#swagger-ui',
			});
		  };
		</script>
		</body>
		</html>`))
	})
}

func AddSpotlight(router *gin.Engine) {
	router.GET("/docs/spotlight", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`<!doctype html>
		<html lang="en">
		  <head>
			<meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
			<title>Elements in HTML</title>
			<script src="https://unpkg.com/@stoplight/elements/web-components.min.js"></script>
			<link rel="stylesheet" href="https://unpkg.com/@stoplight/elements/styles.min.css">
		  </head>
		  <body style="height: 100vh;">
			<elements-api
			  apiDescriptionUrl="/openapi.yaml"
			  router="hash"
			  layout="sidebar"
			/>
		  </body>
		</html>`))
	})
}

func AddScalar(router *gin.Engine) {
	router.GET("/docs/scalar", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`<!doctype html>
		<html>
		  <head>
			<title>API Reference</title>
			<meta charset="utf-8" />
			<meta
			  name="viewport"
			  content="width=device-width, initial-scale=1" />
		  </head>
		  <body>
			<script
			  id="api-reference"
			  data-url="/openapi.json"></script>
			<script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
		  </body>
		</html>`))
	})
}
