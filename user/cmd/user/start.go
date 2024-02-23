package main

import (
	"context"
	"fmt"

	"github.com/Mu-Exchange/Mu-End/common"
	"github.com/Mu-Exchange/Mu-End/common/config"
	"github.com/Mu-Exchange/Mu-End/common/proto"
	"github.com/Mu-Exchange/Mu-End/common/repo"
	"github.com/Mu-Exchange/Mu-End/common/utils/async"
	_ "github.com/Mu-Exchange/Mu-End/user/handler"
	"github.com/Mu-Exchange/Mu-End/user/pkg/base"
	repo2 "github.com/Mu-Exchange/Mu-End/user/repo"
	logrus_v4 "github.com/go-micro/plugins/v4/logger/logrus"
	"github.com/go-micro/plugins/v4/registry/consul"
	_ "github.com/go-micro/plugins/v4/transport/grpc"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
)

var (
	startCMD = cli.Command{
		Name:   "start",
		Usage:  "Start a long-running daemon process",
		Action: start,
	}
)

func start(ctx *cli.Context) error {

	fmt.Println(getVersion(true))

	repoRoot, err := repo.PathRootWithDefault(ctx.String("repo"), repo2.GetEnvDir())
	if err != nil {
		return err
	}

	cfg, err := config.LoadConfig(repoRoot, repo2.ModuleName, repo2.GetEnvPrefix())
	if err != nil {
		return fmt.Errorf("load config failed: %w", err)
	}

	log := repo.SetupLogger(repoRoot, cfg.Log, repo2.ModuleName)
	components := repo.Init(context.Background(), cfg, log)

	return run(components)
}

func run(components repo.CommonComponents) error {
	logger.DefaultLogger = logrus_v4.NewLogger(logrus_v4.WithLogger(components.Logger))

	// Create service
	srv := micro.NewService(
		micro.Name(common.UserService),
		micro.Version(common.Version),
		micro.Registry(consul.NewRegistry(registry.Addrs(components.Cfg.Consul.Addrs...))),
		micro.Flags(&cli.StringFlag{
			Name:  "repo",
			Usage: "User repository path",
		}),
	)
	srv.Init()

	components.Client = srv.Client()

	app, err := base.BuildComponents(components, func(handler proto.UserHandler, commonComponents repo.CommonComponents) {
		// Register handler
		if err := proto.RegisterUserHandler(srv.Server(), handler); err != nil {
			commonComponents.Logger.Fatal(err)
			return
		}
		// Run service
		commonComponents.Logger.Infof("%s service start....", repo2.ModuleName)
		async.SafeAsyncExecute(components.Logger, func() {
			if err := srv.Run(); err != nil {
				commonComponents.Logger.Fatal(err)
				return
			}
		})

	})
	if err != nil {
		return err
	}
	app.Run()

	return nil
}
