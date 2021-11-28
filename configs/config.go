package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"time"
)

type Config struct {
	AppEnv string

	ServerAddress string

	DbDriver   string
	DbUser     string
	DbPassword string
	DbPort     string
	DbHost     string
	DbName     string

	SecretKey string

	Root string

	ExpAccessTokenDuration time.Duration

	ExpCacheTimeDb time.Duration

	TelegramBotToken string
	TelegramChatID   string
}

var Common *Config

func Get() *Config {
	return Common
}

func LoadConfigs(mode string) {
	viper.SetConfigType("env")
	curDir, _ := os.Getwd()
	fmt.Printf("cur dir: %s \n", curDir)
	viper.SetConfigName("dev")
	viper.SetConfigFile(".env.dev")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Common)
	if err != nil {
		panic(err)
	}

	//viper.SetDefault("SERVER_ADDR", "0.0.0.0:8080")

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
			DbPassword: getString("DB_PASSWORD"),
			DbPort:     getString("DB_PORT"),
			DbHost:     getString("DB_HOST"),
			DbName:     getString("DB_NAME"),

			SecretKey: getString("SECRET_KEY"),

			Root: getStringD("ROOT", "/home/gumball/"),

			ExpAccessTokenDuration: getTimeDurationWithDefault("EXPIRED_DURATION", 15),

			ExpCacheTimeDb: getTimeDurationWithDefault("CACHE_TIME_DB", 60),

			TelegramBotToken: getString("FILE_MANAGER_TELEGRAM_BOT_TOKEN"),
			TelegramChatID:   getString("FILE_MANAGER_TELEGRAM_CHAT_ID"),
		}
	*/
}
