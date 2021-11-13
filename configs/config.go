package configs

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	AppEnv string `mapstructure:"MODE"`

	ServerAddress string `mapstructure:"SERVER_ADDR"`

	DbDriver   string `mapstructure:"DB_DRIVER"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASS"`
	DbPort     string `mapstructure:"DB_PORT"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbName     string `mapstructure:"DB_NAME"`

	SecretKey string `mapstructure:"SECRET_KEY"`

	Root string `mapstructure:"ROOT"`

	ExpiredDuration time.Duration `mapstructure:"EXPIRED_DURATION"`

	ExpCacheTimeDb time.Duration `mapstructure:"CACHE_TIME_DB"`

	TelegramBotToken string `mapstructure:"FILE_MANAGER_TELEGRAM_BOT_TOKEN"`
	TelegramChatID   string `mapstructure:"FILE_MANAGER_TELEGRAM_CHAT_ID"`
}

var Common *Config

func Get() *Config {
	return Common
}

func LoadConfigs(mode string) {
	viper.SetConfigType("env")
	//viper.AddConfigPath("")
	viper.SetConfigFile(".env.dev")

	//var configName string
	// load env from cmd
	viper.AutomaticEnv()

	if mode == "dev" {
		//	configName = "dev"
	} else if mode == "prod" {
		//	configName = "prod"
	} else {
		//	configName = "dev"
	}
	//viper.SetConfigName("")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Common)
	if err != nil {
		panic(err)
	}
	//viper.SetDefault("SERVER_ADDR", "0.0.0.0:8080")
	Common.ExpCacheTimeDb *= time.Minute
	Common.ExpiredDuration *= time.Minute
	return
	/*
		var pathConfig string
		switch mode {
		case "dev":
			pathConfig = `.env.dev`
		case "prod":
			pathConfig = `.env.prod`
		default:
			pathConfig = `.env.dev`
		}
		LoadEnv(pathConfig)
		Common = &Config{
			AppEnv: mode,

			ServerAddress: getStringD("SERVER_ADDR", "0.0.0.0:8080"),

			DbDriver:   getString("DB_DRIVER"),
			DbUser:     getString("DB_USER"),
			DbPassword: getString("DB_PASS"),
			DbPort:     getString("DB_PORT"),
			DbHost:     getString("DB_HOST"),
			DbName:     getString("DB_NAME"),

			SecretKey: getString("SECRET_KEY"),

			Root: getStringD("ROOT", "/home/gumball/"),

			ExpiredDuration: getTimeDurationWithDefault("EXPIRED_DURATION", 15),

			ExpCacheTimeDb: getTimeDurationWithDefault("CACHE_TIME_DB", 60),

			TelegramBotToken: getString("FILE_MANAGER_TELEGRAM_BOT_TOKEN"),
			TelegramChatID:   getString("FILE_MANAGER_TELEGRAM_CHAT_ID"),
		}
	*/
}
