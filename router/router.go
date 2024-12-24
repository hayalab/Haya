package router

import (
	"github.com/melodydev777/Melody/conf"
	"github.com/melodydev777/Melody/controllers"
	"github.com/melodydev777/Melody/core"
	"github.com/melodydev777/Melody/middlewares"
)

func Router() {
	core.GetEngine().Use(middlewares.Cors)

	templatesPath := conf.GetConfigString("app", "templates")
	core.GetEngine().LoadHTMLGlob(templatesPath + "/*.tmpl")

	// register routes
	core.AutoRoute(&controllers.IndexController{})
	middlewareInst := new(middlewares.Middleware)

	// /security/**
	securityRouterGroup := core.GetEngine().Group("/security")
	securityRouterGroup.Use(middlewareInst.AdminToken)
}
