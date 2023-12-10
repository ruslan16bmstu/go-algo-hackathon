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

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		h.ServeHTTP(w, r)
	})
}

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
	if err := http.ListenAndServe(fmt.Sprintf(":%d", httpPort), cors(mux)); err != nil {
		panic(err)
	}
}

func RunGrpc(grpcPort int, stocksSrc, predictionsSrc string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTraderServer(s, &RPCServer{
		DB: db.Init(stocksSrc, predictionsSrc),
	})
	log.Printf("grpc-server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
