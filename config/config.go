package config

import (
	"github.com/spf13/viper"
	"log"
)

var (
	config Config
)

type DBTypeConfig string

type Config struct {
	Server ServerConfig `mapstructure:"server"`
	Mysql  MysqlConfig  `mapstructure:"mysql"`
	JWT    JWTConfig    `mapstructure:"jwt"`
	Sqlite SqliteConfig `mapstructure:"sqlite"`
	DBType DBTypeConfig `mapstructure:"dbtype"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
	Mode string `mapstructure:"mode"`
}

type SqliteConfig struct {
	DBName string `mapstructure:"dbname"`
}
type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
	Expire string `mapstructure:"expire"`
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %s", err)
	}
	// 将配置反序列化到结构体
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("反序列化配置失败: %s", err)
	}
}

func Env() *Config {
	return &config
}
