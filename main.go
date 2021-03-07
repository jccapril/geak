package main

import (
	"context"
	"errors"
	login "geak/model"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	login.UnimplementedLoginServer
}

func (s *server) Login(ctx context.Context, in *login.LoginRequest) (*login.LoginResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return nil,errors.New("1111")
}

func main()  {

	lis,err := net.Listen("tcp", ":18080")
	if err != nil {
		log.Fatal("failed to listem: %v",err)
	}
	s := grpc.NewServer()
	login.RegisterLoginServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}



}