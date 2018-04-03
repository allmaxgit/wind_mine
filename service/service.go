package service

import (
	"fmt"
	"net"
	"net/http"
	"os"

	buyerGW "WindToken/api/buyer"
	"WindToken/controllers/buyer"

	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// StartGRPCServer starts gRPC server for REST wrapper.
func StartGRPCServer(gRPCPort uint) (err error) {
	s := grpc.NewServer()
	buyerGW.RegisterAuthServer(
		s,
		&buyer.RPC{},
	)

	lis, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", gRPCPort))
	if err != nil {
		return
	}

	// Start listening.
	err = s.Serve(lis)
	return
}

// StartRESTGetaway starts REST wrapper for gRPC.
func StartRESTGetaway(RESTPort, gRPCPort uint) (err error) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	err = buyerGW.RegisterAuthHandlerFromEndpoint(ctx, mux, fmt.Sprintf(":%d", gRPCPort), opts)
	if err != nil {
		return
	}

	methods := handlers.AllowedMethods([]string{"POST"})
	hCORS := handlers.CORS(methods)(mux)

	// Start listening.
	err = http.ListenAndServe(fmt.Sprintf(":%d", RESTPort), handlers.LoggingHandler(os.Stdout, hCORS))
	return
}
