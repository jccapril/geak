package main

import (
	"flag"
	"geak/gkd/mysql"
	"geak/libs/conf"
	"geak/libs/log"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main(){

	flag.Parse()
	if err := conf.Init(); err != nil {
		panic(err)
	}
	log.Init(conf.Conf.Log)
	mysql.Init(conf.Conf)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		mysql.InitSSQData()
		wg.Done()
	}()

	go func() {
		mysql.ImportDLTData()
		wg.Done()
	}()

	wg.Wait()


	//signalHandler()
}

func signalHandler() {
	var (
		ch  = make(chan os.Signal, 1)
	)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			log.Info("get a signal %s, stop the consume process", zap.Any("sigal",si.String()))
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}