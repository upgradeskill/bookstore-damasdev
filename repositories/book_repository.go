package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/upgradeskill/bookstore-damasdev/helper"
	"github.com/upgradeskill/bookstore-damasdev/interfaces"
	"github.com/upgradeskill/bookstore-damasdev/models"
)

func NewBookRepository() interfaces.BookRepository {
	return &BookRepositoryImpl{}
}

type BookRepositoryImpl struct{}

func (repository *BookRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, book models.BookModel) models.BookModel {
	SQL := "INSERT INTO books (name) VALUES (?)"
	result, err := tx.ExecContext(ctx, SQL, book.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	book.Id = uint(id)
	return book
}

func (repository *BookRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, book models.BookModel) models.BookModel {
	SQL := "UPDATE books SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, book.Name, book.Id)
	helper.PanicIfError(err)

	return book
}

func (repository *BookRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, book models.BookModel) {
	SQL := "DELETE FROM books WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, book.Id)
	helper.PanicIfError(err)
}

func (repository *BookRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, bookId uint) (models.BookModel, error) {
	SQL := "SELECT id, name FROM books WHERE id = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, SQL, bookId)
	helper.PanicIfError(err)

	book := models.BookModel{}
	if rows.Next() {
		err := rows.Scan(&book.Id, &book.Name)
		helper.PanicIfError(err)
		return book, nil
	} else {
		return book, errors.New("book is not found")
	}
}

func (repository *BookRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []models.BookModel {
	SQL := "SELECT id, name FROM books"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var books []models.BookModel
	for rows.Next() {
		book := models.BookModel{}
		err := rows.Scan(&book.Id, &book.Name)
		helper.PanicIfError(err)

		books = append(books, book)
	}

	return books
}
