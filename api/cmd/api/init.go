package main

import (
	"fmt"

	repo2 "github.com/Mu-Exchange/Mu-End/api/repo"
	"github.com/Mu-Exchange/Mu-End/common/repo"
	"github.com/urfave/cli/v2"
)

func initCMD() *cli.Command {
	return &cli.Command{
		Name:   "init",
		Usage:  fmt.Sprintf("Initialize %s local configuration", repo2.ModuleName),
		Action: initialize,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "config",
				Value: "",
				Usage: repo2.ModuleName + " config repo path",
			},
			&cli.BoolFlag{
				Name:  "interactive,i",
				Usage: fmt.Sprintf("configure %s in interactive mode", repo2.ModuleName),
			},
		},
	}
}

func initialize(ctx *cli.Context) error {
	repoRoot, err := repo.PathRootWithDefault(ctx.String("repo"), repo2.GetEnvDir())
	if err != nil {
		return err
	}

	fmt.Printf("initializing %s at %s\n", repo2.ModuleName, repoRoot)

	return repo2.Initialize(repoRoot)
}
