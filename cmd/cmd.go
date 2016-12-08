package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/eapastor/shop/cmd/routes"
	eztemplate "github.com/michelloworld/ez-gin-template"
)

const (
	port string = `3000`
)

var router *gin.Engine

func init() {

	gin.SetMode("debug")

	router = gin.Default()

	render := eztemplate.New()


	render.TemplatesDir = "templates/" // default

	render.Debug = gin.IsDebugging()

	setRoutes(router)

	router.Static("/assets", "./assets")
	router.HTMLRender = render.Init()

}

func setRoutes (router *gin.Engine) {
	router.GET("/", routes.IndexHandler)
	router.GET("/cart", routes.Cart.Show)
	router.POST("/cart", routes.Cart.Add)
}

func Start() {
	router.Run(":" + port)
}