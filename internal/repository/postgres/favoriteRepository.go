package postgres

import (
	"context"
	"crud/internal/core/interface/repository"
	"crud/internal/core/model"
	"crud/internal/lib/db"
	"crud/internal/repository/dbModel"
	"fmt"
)

type _favoriteRepository struct {
	db *db.Db
}

func (favoriteRepository _favoriteRepository) AddFavorite(ctx context.Context, login string, bookId int) error {

	_, err := favoriteRepository.db.PgConn.Exec(ctx,
		`INSERT INTO public.favorite(login, id) values ($1,$2)`,
		login,
		bookId)

	return err
}

func (favoriteRepository _favoriteRepository) GetFavorite(ctx context.Context, login string) ([]model.Book, error) {
	row, _ := favoriteRepository.db.PgConn.Query(ctx,
		`SELECT f.login, b.author, b.genre, b.name FROM public.favorite f join public.book b on f.book_id = b.id WHERE f.login LIKE $1`,
		login)

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

func NewFavoriteRepository(db *db.Db) repository.FavoriteRepository { return _favoriteRepository{db} }
