package config

type Config struct {
	App      App      `mapstructure:"app"`
	Database Database `mapstructure:"database"`
	Log      Log      `mapstructure:"log"`
}
