package service

import (
	"context"
	"database/sql"
	pb "my_module/genproto/library"
)

type bookService struct {
	pb.UnimplementedLibraryServer
	DB *sql.DB
}

func NewBookService(db *sql.DB) *bookService{
	return &bookService{DB: db}
}

func (s *bookService) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.AddBookResponse, error) {
	tr, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Commit()

	result := pb.AddBookResponse{}
	row := s.DB.QueryRow("INSERT INTO Book(title,author,year_published) VALUES($1,$2,$3) returning id", req.Title, req.Author, req.YearPublished)
	err = row.Scan(&result.BookId)
	if err != nil {
		return nil, err
	}
	return &result, err
}

func (s *bookService) SearchBook(ctx context.Context, req *pb.SearchBookRequest) (*pb.SearchBookResponse, error) {
	tr, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Commit()

	rows, err := s.DB.Query(req.Query)
	if err != nil {
		return nil, err
	}
	books := pb.SearchBookResponse{}
	for rows.Next() {
		book := pb.Book{}
		err = rows.Scan(&book.BookId, &book.Title, &book.Title, &book.YearPublished)
		if err != nil {
			return nil, err
		}
		books.Books = append(books.Books, &book)
	}
	return &books, err
}

func (s *bookService) Borrowbook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	tr, err := s.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tr.Commit()

	_, err = s.DB.Exec("INSER INTO Borrowbook(book_id,user_id) VALUES($1,$2)", req.BookId, req.UserId)
	if err != nil {
		return &pb.BorrowBookResponse{Success: false}, err
	}
	return &pb.BorrowBookResponse{Success: true}, nil
}
