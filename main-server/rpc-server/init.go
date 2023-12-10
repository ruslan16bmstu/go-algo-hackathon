package rpc_server

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"pocket-trader-backend/db"
	pb "pocket-trader-backend/pb/gen"
)

func RunRest(ip string, grpcPort, httpPort int) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterTraderHandlerFromEndpoint(ctx, mux, fmt.Sprintf("%s:%d", ip, grpcPort), opts)
	if err != nil {
		panic(err)
	}
	log.Printf("http-server listening at port %d", httpPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), mux); err != nil {
		panic(err)
	}
}

func RunGrpc(grpcPort int, jsonSrc string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTraderServer(s, &RPCServer{
		DB: db.Init(jsonSrc),
	})
	log.Printf("grpc-server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
