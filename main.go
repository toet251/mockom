package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	orders "toet.io/mockom/pb"
)

func main() {
	port := ":11111"
	f, err := os.OpenFile("./server.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println("---------------------------------------------------------------------------")
	log.Println("START SERVER")

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	omv3Server := orders.NewServerImpl()
	orders.RegisterOrderCapturingServiceServer(s, omv3Server)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	log.Println("SERVER LISTENING AT %s", port)
}
