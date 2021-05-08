package main

import (
	"flag"
	"geak/job/conf"
	"geak/job/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	s *service.Service
)

func main(){
	flag.Parse()
	if err := conf.Init(); err != nil {
		log.Fatalf("conf.Init() error(%v)", err)
		panic(err)
	}

	service.New(conf.Conf)

	//http.Init(conf.Conf, s)

	signalHandler()

}


func signalHandler() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			s.Close()
			s.Wait()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}