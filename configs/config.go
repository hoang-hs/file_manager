package configs

type Config struct {
	Port string

	DbDriver   string
	DbUser     string
	DbPassword string
	DbPort     string
	DbHost     string
	DbName     string

	SecretKey string

	ExpiredDuration int
}

var Common *Config

func Get() *Config {
	return Common
}

func LoadConfigs() {
	LoadEnv()
	Common = &Config{
		Port: getStringD("PORT", "8080"),

		DbDriver:   getString("DB_DRIVER"),
		DbUser:     getString("DB_USER"),
		DbPassword: getString("DB_PASSWORD"),
		DbPort:     getString("DB_PORT"),
		DbHost:     getString("DB_HOST"),
		DbName:     getString("DB_NAME"),

		SecretKey: getString("SECRET_KEY"),

		ExpiredDuration: getIntD("EXPIRED_DURATION", 8760),
	}
}
