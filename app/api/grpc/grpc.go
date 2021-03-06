package grpc

import (
	"geak/api/service/lottery"
	gateway_lottery "geak/biz/lottery"
	"geak/libs/conf"
	"google.golang.org/grpc"
	"net"
)

var (
	lotterySrv *lottery.Service
)

func Init(c *conf.Config){

	initService(c)
	lis,err := net.Listen("tcp", ":443")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	route(s)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}

func initService(c *conf.Config){
	lotterySrv = lottery.New(c)
}

func route(s *grpc.Server) {
	gateway_lottery.RegisterLotteryServer(s,&Lottery{})
}