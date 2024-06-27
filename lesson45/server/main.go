package main

import (
	"log"
	pb "my_module/genproto/library"
	"my_module/service"
	"my_module/storage/postgres"
	"net"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":7070")
	if err != nil {
		panic(err)
	}

	defer listener.Close()

	server := grpc.NewServer()

	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}

	serve := service.NewBookService(db)

	pb.RegisterLibraryServer(server, serve)

	log.Println("Server is listening on port 7070")

	err = server.Serve(listener)
	if err != nil {
		panic(err)
	}
}
