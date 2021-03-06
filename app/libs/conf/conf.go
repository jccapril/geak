package conf

import (
	"flag"
	"geak/libs/cache"
	"geak/libs/database"
	"geak/libs/log"
	"github.com/BurntSushi/toml"
)

var (
	// config
	confPath string
	// Conf .
	Conf = &Config{}
)

type Config struct {
	DB 			*database.Config
	Redis 		*cache.Config
	App			*App
	Log			*log.Config
}

type App struct {
	Bid		  string
	Resources string
}


func init() {
	flag.StringVar(&confPath, "conf", "./conf.toml", "default config path")
}

// Init init conf.
func Init() (err error) {
	if confPath != "" {
		return local()
	}
	return remote()
}


func local()(err error){
	_, err = toml.DecodeFile(confPath, &Conf)
	return
}

func remote()(err error){
	return
}