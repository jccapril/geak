package main

import (
	"flag"
	"geak/gkd"
	"geak/libs/conf"
	"geak/libs/log"
)

func main(){

	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Log)
	gkd.Init(conf.Conf)
	gkd.InitData()

}