package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	trippb "server/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cann't connect server: %v", err)
	}

	client := trippb.NewTripServiceClient(conn)
	resp, err := client.GetTrip(context.Background(), &trippb.GetTripRequest{Id: "trip456"})
	if err != nil {
		log.Fatalf("cann't call GetTrip: %v", err)
	}

	fmt.Println(resp)
}
