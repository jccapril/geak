package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main(){
	fmt.Println("hello world")
	signalHandler()
}


func signalHandler() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		si := <-ch
		switch si {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:

			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}