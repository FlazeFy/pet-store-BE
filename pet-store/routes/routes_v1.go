package routes

import (
	"net/http"
	syshandlers "pet-store/modules/systems/http_handlers"

	"github.com/labstack/echo"
)

func InitV1() *echo.Echo {
	e := echo.New()

	e.GET("api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Pet-Store")
	})

	// =============== Public routes ===============

	// Dictionary
	e.GET("api/v1/dct/:type", syshandlers.GetDictionaryByType)

	// Tag
	e.GET("api/v1/tag", syshandlers.GetAllTag)

	// =============== Private routes (Admin) ===============

	return e
}
