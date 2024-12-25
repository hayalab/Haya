package core

import (
	"github.com/hayalab/Haya/tools/timer"
)

type Crond interface {
	GetDurationMillisecond() uint32
	Init()
	Worker()
}

func RegisterCrond(c Crond) {

	c.Init()

	timer.DoTimer(c.GetDurationMillisecond(), c.Worker)
}
