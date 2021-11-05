package configs

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

	ExpiredDuration int
}

var Common *Config

func Get() *Config {
	return Common
}

func LoadConfigs(mode string) {
	var pathConfig string
	switch mode {
	case "dev":
		pathConfig = `.env.dev`
	case "prod":
		pathConfig = `config/.env.dev`
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

		ExpiredDuration: getIntD("EXPIRED_DURATION", 8760),
	}
}
