package service

import (
	"log"
	"net"
	"fmt"
	"google.golang.org/grpc"
)

// StartGRPCServer starts gRPC server for rest wrapper.
func (srv *Server) StartGRPCServer(gRPCPort uint) {

	s := grpc.NewServer()
	auth_gw.RegisterAuthServer(
		s,
		&action.Rpc{},
	)

	//bro := BrokerSrv()
	//go BrokerSrvStart(bro)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", gRPCPort))
	if err != nil {
		uErr.Fatal(err, "failed to listen")
	}

	log.Println(fmt.Sprintf("gRPC - Start listening on port: %d", gRPCPort))
	if err := s.Serve(lis); err != nil {
		uErr.Fatal(err, "failed to serve")
	}
}
