package setting

import (
	"io"
	"os"
	"path/filepath"
	"runtime"

	"github.com/pelletier/go-toml/v2"
)

// Config settings
type Config struct{}

// LoadConfigFile reads and parses the specified config file
// param configFile: the name of the config file to load. Must be in the `conf` directory
func LoadConfigFile(configFile string) (Config, error) {
	var cfg Config
	configFP := filepath.Join(getConfDir(), configFile)

	// Get the config file info
	fileInfo, err := os.Stat(configFP)
	if err != nil {
		return cfg, err
	}

	// Open the config file
	file, err := os.Open(configFP)
	if err != nil {
		return cfg, err
	}
	defer file.Close()

	// Parse the config file
	data := make([]byte, fileInfo.Size())
	_, err = io.ReadFull(file, data)
	if err != nil {
		return cfg, err
	}

	// Unmarshal the config file data
	err = toml.Unmarshal(data, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

// getConfDir returns the absolute path to the `conf` directory
func getConfDir() string {
	_, settingFile, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(settingFile), "../../conf/")
}
