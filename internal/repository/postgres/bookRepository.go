package postgres

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/model"
	"crud/internal/lib/db"
	"crud/internal/repository/dbModel"
	"fmt"
)

type _bookRepository struct {
	db *db.Db
}

func (bookRepository _bookRepository) GetBook(ctx context.Context, bookId int) (model.Book, error) {
	var book dbModel.Book

	err := bookRepository.db.PgConn.QueryRow(ctx,
		`SELECT b.name, b.author, b.genre FROM public.book b WHERE b.id=$1`,
		bookId).Scan(&book.Name, &book.Author, &book.Genre)

	if err != nil {
		return model.Book{}, fmt.Errorf("Ошибка при получении книги: %s", err.Error())
	}

	return model.Book(book), nil
}

func (bookRepository _bookRepository) GetBookByName(ctx context.Context, bookName string) (model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (bookRepository _bookRepository) GetBooksByAuthor(ctx context.Context, bookAuthor string) ([]model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (bookRepository _bookRepository) GetBooksByGenre(ctx context.Context, bookGenre string) ([]model.Book, error) {
	//TODO implement me
	panic("implement me")
}

func NewBookRepo(db *db.Db) repository.BookRepository { return _bookRepository{db} }
