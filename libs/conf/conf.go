package conf

import (
	"flag"
	"geak/libs/cache"
	"geak/libs/database"
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
}



func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
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