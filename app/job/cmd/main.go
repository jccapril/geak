/*
 蒋晨成 @copyright
 */

package main

import (
	"flag"
	"geak/libs/conf"
	"geak/job/service"
	"geak/libs/log"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	s *service.Service
)



func main(){
	flag.Parse()
	if err := conf.Init(); err != nil {

		panic(err)
	}
	log.Init(conf.Conf.Log)
	service.New(conf.Conf)

	signalHandler()

}


func signalHandler() {
	var (
		err error
		ch  = make(chan os.Signal, 1)
	)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("get a signal %s, stop the consume process", zap.Any("sigal",si.String()))
			if err = s.Close(); err != nil {
				log.Error("srv close consumer error(%v)", zap.Error(err))
			}
			time.Sleep(5 * time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}