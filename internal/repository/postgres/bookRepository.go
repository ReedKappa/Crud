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

func (bookRepository _bookRepository) AddBook(ctx context.Context, book model.Book) (int, error) {
	bookDb := dbModel.Book(book)

	var id int

	err := bookRepository.db.PgConn.QueryRow(ctx,
		`INSERT INTO public.book(name, author, genre) values ($1,$2,$3) RETURNING id`,
		bookDb.Name,
		bookDb.Author,
		bookDb.Genre).Scan(&id)

	return id, err
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
	var book dbModel.Book

	err := bookRepository.db.PgConn.QueryRow(ctx,
		`SELECT b.name, b.author, b.genre FROM public.book b WHERE b.name LIKE $1`,
		bookName).Scan(&book.Name, &book.Author, &book.Genre)

	if err != nil {
		return model.Book{}, fmt.Errorf("Ошибка при получении книги: %s", err.Error())
	}

	return model.Book(book), nil
}

func (bookRepository _bookRepository) GetBooksByAuthor(ctx context.Context, bookAuthor string) ([]model.Book, error) {
	row, _ := bookRepository.db.PgConn.Query(ctx,
		`SELECT b.name, b.author, b.genre FROM public.book b WHERE b.author LIKE $1`,
		bookAuthor)

	var models []model.Book

	for row.Next() {
		var book dbModel.Book
		err := row.Scan(&book.Name, &book.Author, &book.Genre)

		if err != nil {
			return nil, fmt.Errorf("Ошибка при получении книги: %s", err.Error())
		}

		models = append(models, model.Book(book))
	}

	return models, nil
}

func (bookRepository _bookRepository) GetBooksByGenre(ctx context.Context, bookGenre string) ([]model.Book, error) {
	row, _ := bookRepository.db.PgConn.Query(ctx,
		`SELECT b.name, b.author, b.genre FROM public.book b WHERE b.genre LIKE $1`,
		bookGenre)

	var models []model.Book

	for row.Next() {
		var book dbModel.Book
		err := row.Scan(&book.Name, &book.Author, &book.Genre)

		if err != nil {
			return nil, fmt.Errorf("Ошибка при получении книги: %s", err.Error())
		}

		models = append(models, model.Book(book))
	}

	return models, nil
}

func NewBookRepo(db *db.Db) repository.BookRepository { return _bookRepository{db} }
