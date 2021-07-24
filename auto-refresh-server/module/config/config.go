package config

import (
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type config struct {
	Apps []*AppConfig `mapstructure:"apps"`
}

type AppConfig struct {
	ID                  string `mapstructure:"id"`
	Secret              string `mapstructure:"secret"`
	RedisAccessTokenKey string `mapstructure:"redis_access_token_key"`
	RedisJsAPITicketKey string `mapstructure:"redis_js_api_ticket_key"`
	RefreshDuration     int64  `mapstructure:"refresh_duration"`
}

var (
	Config *config
)

func Init() {
	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	configFilename := "./config.yaml"
	if v := os.Getenv("WECHAT_REFRESH_CONFIG"); v != "" {
		configFilename = v
	}
	viper.SetConfigFile(configFilename)
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s\n", err)
	}
	viper.WatchConfig()

	Config = new(config)
	err = viper.Unmarshal(Config)
	if err != nil { // Handle errors reading the config file
		log.Fatalf("Fatal error config file: %s\n", err)
	}
	viper.OnConfigChange(func(in fsnotify.Event) {
		err = viper.Unmarshal(Config)
		if err != nil { // Handle errors reading the config file
			log.Printf("Fatal error config file: %s\n", err)
		}
	})
}

func GetApps() []*AppConfig {
	return Config.Apps
}

func GetApp(appID string) *AppConfig {
	apps := GetApps()
	for _, v := range apps {
		if v.ID == appID {
			return v
		}
	}
	return nil
}

func ExistAppID(appID string) bool {
	apps := GetApps()
	for _, v := range apps {
		if v.ID == appID {
			return true
		}
	}
	return false
}
