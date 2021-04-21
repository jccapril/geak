package main

import (
	"context"
	"fmt"
	"geak/api/controllers"
	"gitee.com/jlab/biz/login"
	"gitee.com/jlab/biz/lottery"
	"gitee.com/jlab/biz/user"
	"google.golang.org/grpc/metadata"

	//login "geak/model"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	login.UnimplementedLoginServer
}

func (s *server) NormalLogin(ctx context.Context, in *login.NormalLoginRequest) (*login.NormalLoginResponse, error) {
	log.Printf("Received Moble: %v password:%v", in.GetMobile(),in.GetPassword())

	// 获取请求头
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Printf("get metadata error")
	}
	fmt.Println(md["x-jeak-bid"])
	
	// 返回数据
	return &login.NormalLoginResponse{
		Guid:"100",
		UserInfo:&user.BasicProfile{
			Guid:"100",
			Username:"Admin",
			UserID:100,
		},
	},nil
}

func main()  {

	lis,err := net.Listen("tcp", ":443")
	if err != nil {
		log.Fatal("failed to listem: %v",err)
	}
	s := grpc.NewServer()
	login.RegisterLoginServer(s, &server{})
	lottery.RegisterLotteryServer(s,&controllers.LotteryController{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}