package controllers

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/bookstore-damasdev/dto"
	"github.com/upgradeskill/bookstore-damasdev/helper"
	"github.com/upgradeskill/bookstore-damasdev/infrastructures/database"
	"github.com/upgradeskill/bookstore-damasdev/repositories"
	"github.com/upgradeskill/bookstore-damasdev/services"
)

func FindAll(c echo.Context) error {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()
	result := services.NewBookService(db, repositories.NewBookRepository()).FindAll(ctx)

	return c.JSON(http.StatusOK, result)
}

func Create(c echo.Context) error {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()
	book := dto.BookRequestBody{}
	err := c.Bind(&book)
	helper.PanicIfError(err)

	result := services.NewBookService(db, repositories.NewBookRepository()).Create(ctx, book)

	return c.JSON(http.StatusOK, result)
}

func FindById(c echo.Context) error {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()
	id, err := strconv.Atoi(c.Param("id"))
	helper.PanicIfError(err)
	params := dto.BookRequestParams{
		Id: uint(id),
	}

	result := services.NewBookService(db, repositories.NewBookRepository()).FindById(ctx, params)

	return c.JSON(http.StatusOK, result)
}

func Delete(c echo.Context) error {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()
	id, err := strconv.Atoi(c.Param("id"))
	helper.PanicIfError(err)
	params := dto.BookRequestParams{
		Id: uint(id),
	}

	services.NewBookService(db, repositories.NewBookRepository()).Delete(ctx, params)

	return c.JSON(http.StatusOK, params)
}

func Update(c echo.Context) error {
	db := database.GetConnection()
	defer db.Close()

	ctx := context.Background()
	book := dto.BookRequestBody{}
	err := c.Bind(&book)
	helper.PanicIfError(err)

	id, err := strconv.Atoi(c.Param("id"))
	helper.PanicIfError(err)
	params := dto.BookRequestParams{
		Id: uint(id),
	}

	result := services.NewBookService(db, repositories.NewBookRepository()).Update(ctx, params, book)

	return c.JSON(http.StatusOK, result)
}
