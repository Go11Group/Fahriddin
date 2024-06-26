package main

import (
	"context"
	"fmt"
	pb "my_module/genproto/library"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:7070", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := pb.NewLibraryClient(conn)

	contex,consel := context.WithTimeout(context.Background(),time.Second)

	defer consel()

	book := pb.AddBookRequest{
		Title: "Xamsa",
		Author: "AlisherNavoiy",
		YearPublished: 1744,
	}
	res,err := c.AddBook(contex,&book)
	if err != nil{
		panic(err)
	}
	fmt.Println(res)

}
