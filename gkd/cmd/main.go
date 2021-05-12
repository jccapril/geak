package main

import (
	"flag"
	"geak/gkd"
	"geak/libs/conf"
	"log"

)

func main(){

	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Fatalf("conf.Init() error(%v)", err)
		panic(err)
	}

	gkd.Init(conf.Conf)
	gkd.InitData()

}