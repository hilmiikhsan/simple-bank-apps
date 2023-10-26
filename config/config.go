package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App         App         `yaml:"app"`
	DB          DB          `yaml:"db"`
	JWT         JWT         `yaml:"jwt"`
	RedisServer RedisServer `yaml:"redis"`
	Log         Log         `yaml:"log"`
}

type App struct {
	Port string `yaml:"port"`
}

type DB struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

type JWT struct {
	Issuer            string `yaml:"issuer"`
	Secret            string `yaml:"secret"`
	TokenLifeTimeHour int    `yaml:"tokenLifeTimeHour"`
	MaxRefresh        int    `yaml:"maxRefresh"`
}

type RedisServer struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	Timeout  int    `yaml:"timeout"`
	MaxIdle  int    `yaml:"maxIdle"`
}

type Log struct {
	Path string `yaml:"path"`
}

var Cfg *Config

// var Db *sql.DB

func LoadConfig(filename string) (err error) {
	viper.SetConfigFile(filename)
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		return
	}

	return
}
