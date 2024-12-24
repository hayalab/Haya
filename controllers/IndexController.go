package controllers

import (
	"time"

	"github.com/melodydev777/Melody/tools"

	"github.com/gin-gonic/gin"
	"github.com/melodydev777/Melody/core"
)

type IndexController struct {
	core.BaseController
}

func (ctrl *IndexController) ServerStatus(c *gin.Context) {
	ctrl.JsonSuccess(c, map[string]interface{}{
		"millisecond": tools.GetMillisecond(time.Now()),
	})
}
