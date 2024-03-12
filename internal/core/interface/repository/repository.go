package repository

import (
	"context"
	"crud/internal/core/model"
)

type AuthRepository interface {
	GetUser(ctx context.Context, login, hashPassword string) (string, error)
	Register(ctx context.Context, login, hashPassword string) (string, error)
}

type PostRepository interface {
	CreatePost(ctx context.Context, post model.Post) (int, error)
	GetPost(ctx context.Context, postId int) (model.Post, error)
}

type BookRepository interface {
	GetBook(ctx context.Context, bookId int) (model.Book, error)
	GetBookByName(ctx context.Context, bookName string) (model.Book, error)
	GetBooksByAuthor(ctx context.Context, bookAuthor string) ([]model.Book, error)
	GetBooksByGenre(ctx context.Context, bookGenre string) ([]model.Book, error)
	AddBook(ctx context.Context, book model.Book) (int, error)
}

type FavoriteRepository interface {
	AddFavorite(ctx context.Context, login string, bookId int) error
	GetFavorite(ctx context.Context, login string) ([]model.Book, error)
}
