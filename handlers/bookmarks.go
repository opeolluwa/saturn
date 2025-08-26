package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateBookmark(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Created")

}

func FetchBookmarks(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Fetched")

}

func DeleteBookmark(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Deleted")
}

func UpdateBookmarks(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Updated")
}
