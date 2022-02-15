package configs

import "time"

type Config struct {
	AppEnv string

	ServerAddress string

	DbDriver   string
	DbUser     string
	DbPassword string
	DbPort     string
	DbHost     string
	DbName     string

	GraphiteHost string
	GraphitePort int

	SecretKey string

	Root string

	ExpiredDuration time.Duration

	ExpCacheTimeDb time.Duration

	TelegramBotToken string
	TelegramChatID   string
}

var Common *Config

func Get() *Config {
	return Common
}

func LoadConfigs(mode string) {
	var pathConfig string
	switch mode {
	case "dev":
		pathConfig = `.env`
	case "prod":
		pathConfig = `.env.prod`
	default:
		pathConfig = `.env`
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

		GraphiteHost: getString("GRAPHITE_HOST"),
		GraphitePort: getInt("GRAPHITE_PORT"),

		SecretKey: getString("SECRET_KEY"),

		Root: getStringD("ROOT", "/home/gumball/"),

		ExpiredDuration: getTimeDurationWithDefault("EXPIRED_DURATION", 15),

		ExpCacheTimeDb: getTimeDurationWithDefault("CACHE_TIME_DB", 60),

		TelegramBotToken: getString("FILE_MANAGER_TELEGRAM_BOT_TOKEN"),
		TelegramChatID:   getString("FILE_MANAGER_TELEGRAM_CHAT_ID"),
	}
}
