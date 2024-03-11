package service

import (
	"crud/internal/core/interface/repository"
	"crud/internal/core/interface/service"
	"crud/internal/core/model"
	"errors"
	"golang.org/x/net/context"
	"log/slog"
)

type _bookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) service.BookService {
	return _bookService{repo: repo}
}

func (bookService _bookService) GetBook(ctx context.Context, bookId int) (model.Book, error) {
	return bookService.repo.GetBook(ctx, bookId)
}

func (bookService _bookService) GetBookByName(ctx context.Context, bookName string) (model.Book, error) {
	return bookService.repo.GetBookByName(ctx, bookName)
}

func (bookService _bookService) GetBooksByAuthor(ctx context.Context, bookAuthor string) ([]model.Book, error) {
	return bookService.repo.GetBooksByAuthor(ctx, bookAuthor)
}

func (bookService _bookService) GetBooksByGenre(ctx context.Context, bookGenre string) ([]model.Book, error) {
	return bookService.repo.GetBooksByGenre(ctx, bookGenre)
}

func (bookService _bookService) AddBook(ctx context.Context, book model.Book) (int, error) {
	id, err := bookService.repo.AddBook(ctx, book)

	if err != nil {
		slog.Error(err.Error())
		return 0, errors.New("Ошибка добавления книги")
	}

	return id, nil
}
