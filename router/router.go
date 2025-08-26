package router

import (
	"github.com/labstack/echo/v4"
	"github.com/opeolluwa/saturn/handlers"
)

func LoadRoutes(e *echo.Echo) {
	e.GET("/health", handlers.HealthCheck)

	bookmarks := e.Group("/bookmarks")

	bookmarks.POST("", handlers.CreateBookmark)
	bookmarks.GET("", handlers.FetchBookmarks)
	bookmarks.DELETE("", handlers.DeleteBookmark)
	bookmarks.PUT("", handlers.UpdateBookmarks)
}
