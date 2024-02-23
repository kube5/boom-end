package repo

import (
	"github.com/Mu-Exchange/Mu-End/common"
	"github.com/mitchellh/go-homedir"
	"os"
)

func PathRoot(envDir string) (string, error) {
	dir := os.Getenv(envDir)
	var err error
	if len(dir) == 0 {
		dir, err = homedir.Expand(common.DefaultPathRoot)
	}
	return dir, err
}

func PathRootWithDefault(path string, envDir string) (string, error) {
	if len(path) == 0 {
		return PathRoot(envDir)
	}

	return path, nil
}
