package server

import (
	"engagement360/internal/controller"
	"fmt"
	"github.com/gin-gonic/gin"
)

type PssRouter struct {
	rtr *gin.Engine
}

func (pr PssRouter) Init() error {
	fmt.Println("Initializing server")

	pr.rtr = gin.Default()

	// setup router
	pr.rtr = gin.Default()
	pr.rtr.Static("/css", "internal/ui/css")
	pr.rtr.Static("/js", "internal/ui/js")

	pr.rtr.LoadHTMLGlob("internal/ui/*.html")

	pr.rtr.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", nil)
	})

	pr.rtr.GET("/header", func(ctx *gin.Context) {
		ctx.HTML(200, "header.html", nil)
	})

	pr.rtr.GET("/menu", func(ctx *gin.Context) {
		ctx.HTML(200, "menu.html", nil)
	})

	pr.rtr.GET("/welcome", func(ctx *gin.Context) {
		ctx.HTML(200, "welcome.html", nil)
	})

	controller.EngagementControllerInstance(pr.rtr)

	pr.rtr.Run()
	return nil
}
