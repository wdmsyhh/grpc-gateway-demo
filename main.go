package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	trippb "grpc-gateway-demo/proto"
	trip "grpc-gateway-demo/tripservice"
	"log"
	"net"
	"net/http"
)

func main() {
	log.SetFlags(log.Lshortfile)
	go startGRPCGateway()
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	trippb.RegisterTripServiceServer(server, &trip.Service{})
	log.Fatal(server.Serve(lis))
}

func startGRPCGateway() {
	c := context.Background()
	c, cancel := context.WithCancel(c)
	defer cancel()
	mux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseEnumNumbers: true,
			UseProtoNames:  true,
		},
	}))
	err := trippb.RegisterTripServiceHandlerFromEndpoint(c, mux, ":8081", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatalf("cann't start grpc gateway: %v", err)
	}

	err = http.ListenAndServe(":8080", mux) // grpc gateway 的
	if err != nil {
		log.Fatalf("cann't listen and serve: %v", err)
	}
}
