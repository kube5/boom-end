package repo

import (
	"fmt"
	"strings"
)

// ModuleName is the default module name
const ModuleName = "game"

// GetEnvDir is to get the environment variable used to change the path root.
func GetEnvDir() string {
	return fmt.Sprintf("%s_PATH", strings.ToUpper(ModuleName))
}

// GetEnvPrefix is to get the environment variable prefix used by the config.
func GetEnvPrefix() string {
	return strings.ToUpper(ModuleName)
}
