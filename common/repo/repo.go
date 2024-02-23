package repo

import (
	"bufio"
	"context"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/coreos/etcd/pkg/fileutil"
	"github.com/sirupsen/logrus"
	"go-micro.dev/v4/client"

	"github.com/Mu-Exchange/Mu-End/common/config"
	"github.com/Mu-Exchange/Mu-End/common/utils/log"
)

type CommonComponents struct {
	Ctx      context.Context
	RootPath string
	Logger   *logrus.Logger
	Client   client.Client
	Cfg      *config.Config
}

func SetupLogger(rootPath string, config *config.Log, moduleName string) *logrus.Logger {
	return log.New(config.Level,
		path.Join(rootPath, "logs"),
		moduleName,
		config.MaxSize,
		config.MaxAge.Duration,
		config.RotationTime.Duration)
}

func Initialized(repoRoot string, configName string) bool {
	return fileutil.Exist(filepath.Join(repoRoot, configName))
}

func ReadYes() bool {
	s := readInput()
	return strings.EqualFold(s, "y") || strings.EqualFold(s, "")
}

func readInput() string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}

func Init(ctx context.Context, config *config.Config, logger *logrus.Logger) CommonComponents {
	return CommonComponents{
		Ctx:      ctx,
		RootPath: config.RepoRoot,
		Logger:   logger,
		Cfg:      config,
	}
}
