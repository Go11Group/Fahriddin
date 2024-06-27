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

	contex, consel := context.WithTimeout(context.Background(), time.Second)

	defer consel()

	book := pb.AddBookRequest{
		Title:         "Layli va Majnun",
		Author:        "AlisherNavoiy",
		YearPublished: 1544,
	}
	res, err := c.AddBook(contex, &book)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

}
