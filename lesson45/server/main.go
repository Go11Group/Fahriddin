package main

import (
	"context"
	"database/sql"
	"log"
	pb "my_module/genproto/library"
	"my_module/storage/postgres"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedLibraryServer
	DB *sql.DB
}

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
	pb.RegisterLibraryServer(server, &Server{DB: db})

	log.Println("Server is listening on port 7070")

	err = server.Serve(listener)
	if err != nil{
		panic(err)
	}

}

func (s *Server)AddBook(ctx context.Context, req *pb.AddBookRequest)(*pb.AddBookResponse,error){
	tr,err := s.DB.Begin()
	if err != nil{
		return nil,err
	}
	defer tr.Commit()

	result := pb.AddBookResponse{}
	row := s.DB.QueryRow("INSER INTO Book(title,author,year_published) VALUES($1,$2) returning id",req.Title,req.Author,req.YearPublished)
	err = row.Scan(result.BookId)
	if err != nil{
		return nil,err
	}
	return &result,err
}

func (s *Server) SearchBook(ctx context.Context, req *pb.SearchBookRequest)(*pb.SearchBookResponse,error){
	tr,err := s.DB.Begin()
	if err != nil{
		return nil,err
	}
	defer tr.Commit()

	rows,err := s.DB.Query(req.Query)
	if err != nil{
		return nil,err
	}
	books := pb.SearchBookResponse{}
	for rows.Next(){
		book := pb.Book{}
		err = rows.Scan(&book.BookId,&book.Title,&book.Title,&book.YearPublished)
		if err != nil{
			return nil,err
		}
		books.Books = append(books.Books,&book)
	}
	return &books,err
}

func (s *Server) Borrowbook(ctx context.Context, req *pb.BorrowBookRequest)(*pb.BorrowBookResponse,error){
	tr,err := s.DB.Begin()
	if err != nil{
		return nil,err
	}
	defer tr.Commit()

	_,err = s.DB.Exec("INSER INTO Borrowbook(book_id,user_id) VALUES($1,$2)",req.BookId,req.UserId)
	if err != nil{
		return &pb.BorrowBookResponse{Success: false},err
	}
	return &pb.BorrowBookResponse{Success: true}, nil
}