package config

type Http struct {
	Port    string `env:"PORT" envDefault:"3000"`
	Address string `env:"ADDRESS" envDefault:"0.0.0.0"`
}
