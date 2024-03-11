package service

import (
	"context"
	"crud/internal/core/model"
)

type AuthService interface {
	Register(ctx context.Context, login, password string) (string, error)
	GenerateToken(ctx context.Context, login, password string) (string, error)
}

type PostService interface {
	CreatePost(ctx context.Context, post model.Post) (int, error)
	GetPost(ctx context.Context, postId int) (model.Post, error)
}

type BookService interface {
	GetBook(ctx context.Context, bookId int) (model.Book, error)
	GetBookByName(ctx context.Context, bookName string) (model.Book, error)
	GetBooksByAuthor(ctx context.Context, bookAuthor string) ([]model.Book, error)
	GetBooksByGenre(ctx context.Context, bookGenre string) ([]model.Book, error)
	AddBook(ctx context.Context, book model.Book) (int, error)
}
