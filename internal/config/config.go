package config

type Config struct {
	Addr string
}

func New() Config {
	return Config{Addr: ":8080"}
}
