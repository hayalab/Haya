package router

import (
	"github.com/hayalab/Haya/conf"
	"github.com/hayalab/Haya/controllers"
	"github.com/hayalab/Haya/core"
	"github.com/hayalab/Haya/middlewares"
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
