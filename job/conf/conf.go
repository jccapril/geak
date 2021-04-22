package conf

import (
	"geak/cache"
	"geak/database"
)

var (
	// config
	confPath string
	// Conf .
	Conf = &Config{}
)

type Config struct {
	Env     	string
	DB 			*database.Config
	Redis 		*cache.Config
}

