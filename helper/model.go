package helper

import (
	"github.com/upgradeskill/bookstore-damasdev/dto"
	"github.com/upgradeskill/bookstore-damasdev/models"
)

func ToBookResponse(category models.BookModel) dto.BookResponseBody {
	return dto.BookResponseBody{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToBookResponses(categories []models.BookModel) []dto.BookResponseBody {
	var categoryResponses []dto.BookResponseBody
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToBookResponse(category))
	}

	return categoryResponses
}
