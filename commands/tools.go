package commands

import (
	"github.com/melodydev777/Melody/core"
	"github.com/melodydev777/Melody/tools/log"

	"github.com/urfave/cli"
)

func ParseUserToken(c *cli.Context) {
	loginToken := c.String("login-token")

	tokenInfo, e := core.ParseLoginToken(loginToken)
	if e != nil {
		log.Error("", "parse login token failed %v", tokenInfo)
		return
	}

	log.Info("", "token info %v", tokenInfo)
}
