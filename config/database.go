package config

type Database struct {
	Host                  string `env:"DB_HOST,required" json:"DB_HOST"`
	Port                  string `env:"DB_PORT"`
	User                  string `env:"DB_USER,required"`
	Password              string `env:"DB_PASSWORD,required"`
	Name                  string `env:"DB_NAME,required"`
	MaxOpenConnections    int    `env:"DB_MAX_OPEN_CONNECTIONS" envDefault:"100"`
	MaxIdleConnections    int    `env:"DB_MAX_IDLE_CONNECTIONS" envDefault:"5"`
	ConnectionMaxLifeTime int    `env:"DB_CONNECTION_MAX_LIFETIME" envDefault:"300"`
}

type RedisDatabase struct {
	Host     string `env:"REDIS_DB_HOST,required"`
	Port     string `env:"REDIS_DB_PORT"`
	User     string `env:"REDIS_DB_USER,required"`
	Password string `env:"REDIS_DB_PASSWORD,required"`
	Name     int    `env:"REDIS_DB_NAME,required"`
}
