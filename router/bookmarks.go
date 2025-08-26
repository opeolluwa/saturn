package router

import (
	"github.com/opeolluwa/saturn/handlers"

	"github.com/labstack/echo/v4"
)

func BookmarkRoutes(e *echo.Echo) {

	e.POST("", handlers.CreateBookmark)
	e.GET("", handlers.FetchBookmarks)
	e.DELETE("", handlers.DeleteBookmark)
	e.PUT("", handlers.UpdateBookmarks)
}
