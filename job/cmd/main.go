package main

import (
	"flag"
	"geak/job/conf"
	"log"
)

func main(){
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Fatalf("conf.Init() error(%v)", err)
		panic(err)
	}
}