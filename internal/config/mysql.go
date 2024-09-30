package config

type Mysql struct {
	DSN string `env:"DSN" envDefault:"root:1234@tcp(127.0.0.1:3316)/exify?charset=utf8&parseTime=True&loc=Local"`
}
