package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var (
	CONFIG_FILE = "./configs/config.yml"
)

func LoadConfigs() {
	viper.SetConfigFile(CONFIG_FILE)
	viper.ReadInConfig()

	env := getEnvironment()

	loglevel := viper.Get(fmt.Sprintf("%s.loglevel", env))
	fmt.Println(loglevel)
}

func getEnvironment() string {
	return getEnvVarOrDefault("env", "local").(string)
}

func getEnvVarOrDefault(path string, defaultValue interface{}) interface{} {
	value := os.Getenv(path)
	if value != "" {
		return value
	}
	return defaultValue
}
