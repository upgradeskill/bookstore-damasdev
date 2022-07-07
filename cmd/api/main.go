package main

import (
	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/bookstore-damasdev/router"
)

func main() {
	e := echo.New()
	router.Router(e)

	e.Logger.Fatal(e.Start(":3000"))
}
