package grpc

import (
	"gitee.com/jlab/biz/lottery"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Init(){
	lis,err := net.Listen("tcp", ":443")
	if err != nil {
		log.Fatalf("failed to listem: %v",err)
		panic(err)
	}
	s := grpc.NewServer()
	route(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		panic(err)
	}
}

func route(s *grpc.Server) {
	lottery.RegisterLotteryServer(s,&Lottery{})
}