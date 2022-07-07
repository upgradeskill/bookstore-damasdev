package router

import (
	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/bookstore-damasdev/controllers"
)

func Router(e *echo.Echo) {

	g := e.Group("/api")

	g.GET("/books", controllers.FindAll)
	g.POST("/books", controllers.Create)
	g.GET("/books/:id", controllers.FindById)
	g.PUT("/books/:id", controllers.Update)
	g.DELETE("/books/:id", controllers.Delete)
}
