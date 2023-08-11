package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type TomlConfig struct {
	AppName string
	MySQL   MySQLConfig
	Log     LogConfig
	Redis   RedisConfig
}

// MySQLConfig MySQL相关配置
type MySQLConfig struct {
	Host        string
	Name        string
	Password    string
	Port        int
	TablePrefix string
	User        string
}

// LogConfig 日志保存地址
type LogConfig struct {
	Path  string
	Level string
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
	DB       int
}

var Config TomlConfig

func InitConfig() {
	// 设置文件名
	viper.SetConfigName("config")
	// 设置文件类型
	viper.SetConfigType("toml")
	// 设置文件路径，可以多个viper会根据设置顺序依次查找
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	if err := viper.Unmarshal(&Config); err != nil {
		panic(err)
	}

}
