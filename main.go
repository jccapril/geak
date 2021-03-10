package main

import (
	"context"
	"gitee.com/jlab/biz/login"
	"gitee.com/jlab/biz/user"
	//login "geak/model"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	login.UnimplementedLoginServer
}

func (s *server) Login(ctx context.Context, in *login.NormalLoginRequest) (*login.NormalLoginResponse, error) {
	log.Printf("Received Moble: %v password:%v", in.GetMobile(),in.GetPassword())
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

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}