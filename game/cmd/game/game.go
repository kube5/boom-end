package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Mu-Exchange/Mu-End/common"
	"github.com/Mu-Exchange/Mu-End/game/repo"
	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = common.AppName
	app.Usage = fmt.Sprintf("%s %s", common.AppName, repo.ModuleName)
	app.Compiled = time.Now()

	// global flags
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "repo",
			Usage: fmt.Sprintf("%s repository path", repo.ModuleName),
		},
	}

	app.Commands = []*cli.Command{
		&startCMD,
		&versionCMD,
		initCMD(),
	}

	err := app.Run(os.Args)
	if err != nil {
		color.Red(err.Error())
		os.Exit(-1)
	}
}
