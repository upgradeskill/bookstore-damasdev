package services

import (
	"context"
	"database/sql"

	"github.com/upgradeskill/bookstore-damasdev/dto"
	"github.com/upgradeskill/bookstore-damasdev/helper"
	"github.com/upgradeskill/bookstore-damasdev/interfaces"
	"github.com/upgradeskill/bookstore-damasdev/models"
)

func NewBookService(db *sql.DB, bookRepository interfaces.BookRepository) interfaces.BookService {
	return &BookServiceImpl{
		DB:             db,
		BookRepository: bookRepository,
	}
}

type BookServiceImpl struct {
	BookRepository interfaces.BookRepository
	DB             *sql.DB
}

func (service *BookServiceImpl) Create(ctx context.Context, request dto.BookRequestBody) dto.BookResponseBody {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	book := models.BookModel{
		Name: request.Name,
	}

	result := service.BookRepository.Save(ctx, tx, book)
	return helper.ToBookResponse(result)
}

func (service *BookServiceImpl) Update(ctx context.Context, params dto.BookRequestParams, request dto.BookRequestBody) dto.BookResponseBody {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindById(ctx, tx, params.Id)
	helper.PanicIfError(err)
	book.Name = request.Name

	result := service.BookRepository.Update(ctx, tx, book)
	return helper.ToBookResponse(result)
}

func (service *BookServiceImpl) Delete(ctx context.Context, params dto.BookRequestParams) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	book, err := service.BookRepository.FindById(ctx, tx, params.Id)
	helper.PanicIfError(err)

	service.BookRepository.Delete(ctx, tx, book)
}

func (service *BookServiceImpl) FindById(ctx context.Context, params dto.BookRequestParams) dto.BookResponseBody {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	book, err := service.BookRepository.FindById(ctx, tx, params.Id)
	helper.PanicIfError(err)

	return helper.ToBookResponse(book)
}

func (service *BookServiceImpl) FindAll(ctx context.Context) []dto.BookResponseBody {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	categories := service.BookRepository.FindAll(ctx, tx)
	return helper.ToBookResponses(categories)
}
