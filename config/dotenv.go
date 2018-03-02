package config

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	im "github.com/grokify/gotilla/io/ioutilmore"
	"github.com/grokify/gotilla/os/osutil"
	"github.com/joho/godotenv"
)

var (
	EnvPathVar = "ENV_PATH"
	LocalPath  = "./.env"
)

func DefaultPaths() []string {
	return []string{os.Getenv(EnvPathVar), LocalPath}
}

func LoadEnvDefaults() error {
	envPathsSet := []string{}

	for _, defaultPath := range DefaultPaths() {
		exists, err := osutil.Exists(defaultPath)
		if err == nil && exists {
			envPathsSet = append(envPathsSet, defaultPath)
		}
	}

	if len(envPathsSet) > 0 {
		return godotenv.Load(envPathsSet...)
	}
	return godotenv.Load()
}

func LoadDotEnvSkipEmpty(paths ...string) error {
	if len(paths) == 0 {
		paths = DefaultPaths()
	}

	envPaths := []string{}

	for _, envPathVal := range paths {
		envPathVals := strings.Split(envPathVal, ",")
		for _, envPath := range envPathVals {
			envPath = strings.TrimSpace(envPath)

			good, err := im.IsFileWithSizeGtZero(envPath)
			if err == nil && good {
				envPaths = append(envPaths, envPath)
			}
		}
	}

	if len(envPaths) > 0 {
		return godotenv.Load(envPaths...)
	}
	return nil
}

// GetDotEnvVal retrieves a single var from a `.env` file path
func GetDotEnvVal(envPath, varName string) (string, error) {
	cmd := fmt.Sprintf("grep %s '%s' | rev | cut -d= -f1 | rev", varName, envPath)

	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return "", fmt.Errorf("Failed to execute command: %s", cmd)
	}
	return string(out), nil
}
