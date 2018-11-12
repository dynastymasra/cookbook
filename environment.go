package cookbook

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func checkEnvKey(key string) error {
	if !viper.IsSet(key) && os.Getenv(key) == "" {
		return fmt.Errorf("%v env key is not set", key)
	}
	return nil
}

// StringEnv used for get string value from environment
func StringEnv(key string) (string, error) {
	return viper.GetString(key), checkEnvKey(key)
}

// BoolEnv used for get bool value from environment
func BoolEnv(key string) (bool, error) {
	return viper.GetBool(key), checkEnvKey(key)
}

// IntEnv used for get bool value from environment
func IntEnv(key string) (int, error) {
	return viper.GetInt(key), checkEnvKey(key)
}
