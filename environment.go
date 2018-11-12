package cookbook

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func checkEnvKey(key string) {
	if !viper.IsSet(key) && os.Getenv(key) != "" {
		log.Fatalf("%v env key is not set", key)
	}
}

// StringEnv used for get string value from environment, this will stop application from running if not set
func StringEnv(key string) string {
	checkEnvKey(key)
	return viper.GetString(key)
}

// BoolEnv used for get bool value from environment, this will stop application from running if not set
func BoolEnv(key string) bool {
	checkEnvKey(key)
	return viper.GetBool(key)
}

// IntEnv used for get bool value from environment, this will stop application from running if not set
func IntEnv(key string) int {
	checkEnvKey(key)
	return viper.GetInt(key)
}
