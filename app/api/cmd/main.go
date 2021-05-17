package main

import (
	"flag"
	"geak/api/grpc"
	"geak/libs/conf"
	"geak/libs/log"
)

func main(){
	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Log)
	grpc.Init(conf.Conf)
}
