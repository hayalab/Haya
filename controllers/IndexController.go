package controllers

import (
	"time"

	"github.com/hayalab/Haya/tools"

	"github.com/gin-gonic/gin"
	"github.com/hayalab/Haya/core"
)

type IndexController struct {
	core.BaseController
}

func (ctrl *IndexController) ServerStatus(c *gin.Context) {
	ctrl.JsonSuccess(c, map[string]interface{}{
		"millisecond": tools.GetMillisecond(time.Now()),
	})
}
