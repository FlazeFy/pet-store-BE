package routes

import (
	"net/http"
	animalhandlers "pet-store/modules/animals/http_handlers"
	planthandlers "pet-store/modules/plants/http_handlers"
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

	// Animals
	e.GET("api/v1/animal/:order", animalhandlers.GetAllAnimals)
	e.GET("api/v1/animal/detail/:slug", animalhandlers.GetAnimalDetailBySlug)

	// Plants
	e.GET("api/v1/plant/:order", planthandlers.GetAllPlants)
	e.GET("api/v1/plant/detail/:slug", planthandlers.GetPlantDetailBySlug)

	// =============== Private routes (Admin) ===============

	return e
}
