package config

import "github.com/spf13/viper"

type Application struct {
	IsInit        bool
	ReadTimeout   int
	WriterTimeout int
	Host          string
	Port          string
	Name          string
	JwtSecret     string
	Log           logConf
	Env           string
}

func InitApplication(cfg *viper.Viper) *Application {
	return &Application{
		IsInit:        cfg.GetBool("isInit"),
		ReadTimeout:   cfg.GetInt("readTimeout"),
		WriterTimeout: cfg.GetInt("writerTimeout"),
		Host:          cfg.GetString("host"),
		Port:          cfg.GetString("port"),
		Name:          cfg.GetString("name"),
		JwtSecret:     cfg.GetString("jwtSecret"),
		Env:           cfg.GetString("env"),
	}
}

var ApplicationConfig = new(Application)
