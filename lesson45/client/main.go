package main

import (
	"context"
	pb "my_module/genproto/library"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:707", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := pb.NewLibraryClient(conn)

	conte,consel := context.WithTimeout(context.Background(),time.Second)

	defer consel()

	

}
