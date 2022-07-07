package interfaces

import (
	"context"

	"github.com/upgradeskill/bookstore-damasdev/dto"
)

type BookService interface {
	Create(ctx context.Context, request dto.BookRequestBody) dto.BookResponseBody
	Update(ctx context.Context, params dto.BookRequestParams, request dto.BookRequestBody) dto.BookResponseBody
	Delete(ctx context.Context, params dto.BookRequestParams)
	FindById(ctx context.Context, params dto.BookRequestParams) dto.BookResponseBody
	FindAll(ctx context.Context) []dto.BookResponseBody
}
