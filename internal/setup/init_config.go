package setup

import (
	"fmt"
	"github.com/spf13/viper"
	"golang-restful-api-framework/internal/config"
)

type ConfigFile string

func InitConfigFile(configFile string) ConfigFile {
	return ConfigFile(configFile)
}

func InitConfig(configFile ConfigFile) config.Config {
	var config config.Config
	viper.SetConfigFile(string(configFile))
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("无法加载配置文件: %s\n", err)
		return config
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		fmt.Printf("无法映射配置到结构体: %s\n", err)
		return config
	}
	return config
}
