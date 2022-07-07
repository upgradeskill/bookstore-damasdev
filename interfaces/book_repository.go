package interfaces

import (
	"context"
	"database/sql"

	"github.com/upgradeskill/bookstore-damasdev/models"
)

type BookRepository interface {
	Save(ctx context.Context, tx *sql.Tx, book models.BookModel) models.BookModel
	Update(ctx context.Context, tx *sql.Tx, book models.BookModel) models.BookModel
	Delete(ctx context.Context, tx *sql.Tx, book models.BookModel)
	FindById(ctx context.Context, tx *sql.Tx, bookId uint) (models.BookModel, error)
	FindAll(ctx context.Context, tx *sql.Tx) []models.BookModel
}
