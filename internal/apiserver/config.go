package apiserver

type Config struct {
	BindAddr    string `toml:"bind_addr"`
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr:    ":9000",
		DatabaseURL: "postgres://postgres:@localhost:5432/tasksellerx?sslmode=disable",
	}
}
