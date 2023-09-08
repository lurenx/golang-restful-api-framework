package config

type App struct {
	AppName string `mapstructure:"name"`
	Port    int    `mapstructure:"port"`
}
