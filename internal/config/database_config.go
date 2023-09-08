package config

type Mysql struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
type Sqlite struct {
	DBPath string `mapstructure:"db_path"`
}

type Database struct {
	Mysql  Mysql  `mapstructure:"mysql"`
	Sqlite Sqlite `mapstructure:"sqlite"`
}
