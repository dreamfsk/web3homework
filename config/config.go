package config

import (
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
)

var (
	config Config
	// export db
	DB    *gorm.DB
	B_Zap Zap
)

type DBTypeConfig string

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Mysql    MysqlConfig    `mapstructure:"mysql"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Zap      Zap            `mapstructure:"zap" json:"zap" yaml:"zap"`
	Sqlite   SqliteConfig   `mapstructure:"sqlite"`
	DBType   DBTypeConfig   `mapstructure:"dbtype"`
	Language LanguageConfig `mapstructure:"language"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Host string `mapstructure:"host"`
	Mode string `mapstructure:"mode"`
}

type SqliteConfig struct {
	DBName  string `mapstructure:"dbname"`
	Migrate bool   `mapstructure:"migrate"`
}
type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	Migrate  bool   `mapstructure:"migrate"`
}

type JWTConfig struct {
	Secret string `mapstructure:"secret"`
	Expire string `mapstructure:"expire"`
}

type LanguageConfig struct {
	Local string `mapstructure:"local"`
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

	err := initDatabase()
	if err != nil {
		return
	}
	B_Zap = config.Zap
}

func Env() *Config {
	return &config
}
