package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hayalab/Haya/crond"
	"github.com/hayalab/Haya/sdk/awshelper"
	"github.com/hayalab/Haya/sdk/chatgptapi"
	"github.com/hayalab/Haya/sdk/twitterapi"

	"github.com/hayalab/Haya/commands"
	"github.com/hayalab/Haya/conf"
	"github.com/hayalab/Haya/core"
	"github.com/hayalab/Haya/models"
	"github.com/hayalab/Haya/router"
	"github.com/hayalab/Haya/taskpool"
	"github.com/hayalab/Haya/tools/log"
	"github.com/hayalab/Haya/tools/logger"
	"github.com/urfave/cli"
)

var (
	sigs     = make(chan os.Signal)
	done     = make(chan bool)
	confpath = ""
)

func main() {
	app := &cli.App{
		Name:   "AgentService",
		Usage:  "the main backend of AgentService",
		Action: version,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "config",
				Usage:    "the path of the config file",
				Required: true,
			},
		},
		Before: initProc,
		Commands: []cli.Command{
			{
				Name:   "start",
				Usage:  "start the main server",
				Action: start,
			},
			{
				Name:   "version",
				Usage:  "get the version",
				Action: version,
			},
			{
				Name:  "tool",
				Usage: "all tools commands",
				After: func(c *cli.Context) error {
					// wait for all logs to be written
					for i := 2; i > 0; i-- {
						fmt.Printf("process will exit %d second later\n", i)
						time.Sleep(time.Second)
					}
					return nil
				},
				Subcommands: commands.ToolCommands,
			},
		},
	}

	app.Run(os.Args)
}

func version(c *cli.Context) {
	fmt.Println("0.0.1")
}

func initProc(c *cli.Context) error {

	confpath = c.String("config")

	// initialize configuration
	err := conf.ParseConfigINI(confpath)
	if err != nil {
		fmt.Println("err : parse config failed", err.Error())
		os.Exit(1)
	}

	logPath := conf.GetConfigString("app", "log_path")

	logger.InitLogger(logPath, "20060102") // log
	el, err := conf.GetConfigInt("log", "level")
	if err != nil {
		el = 0
	}
	log.SetLogErrorLevel(int(el))

	// initialize database
	models.InitModel()

	// initialize Twitter configuration
	twitterapi.InitConfig()

	err = chatgptapi.InitChatGPT()
	if err != nil {
		panic(err)
	}

	// initialize Twitter V1
	twitterapi.InitTwitterAPIV1()

	awshelper.InitAwsSDK()

	return nil
}

func start(c *cli.Context) {

	core.InitScheduler()

	// initialize task
	taskpool.InitTaskListeners()

	// initialize scheduled tasks
	crond.InitCrond()

	// initialize router
	router.Router()

	// start program
	go core.Run()

	// serve
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) // ctrl+c, kill, kill -2,
	go sigAwaiter()

	<-done
}

func sigAwaiter() {
	sig := <-sigs
	fmt.Printf("recv signal %s\n", sig.String())

	// todo process exit logic
	logger.DestroyLogger() // flush log

	done <- true
}
