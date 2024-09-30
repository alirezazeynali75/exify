package config

type App struct {
	Env  string `env:"ENV" envDefault:"development"`
	Name string `env:"NAME" envDefault:"EXIFY"`
}