package routes

import (
	"net/http"
	fehandlers "pet-store/modules/doc/http_handlers"
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

	// Tag
	e.GET("api/v1/animal", fehandlers.GetAllAnimals)

	// =============== Private routes (Admin) ===============

	return e
}
