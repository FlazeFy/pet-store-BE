package httphandlers

import (
	"net/http"
	"pet-store/modules/animals/repositories"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllAnimals(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	ord := c.Param("order")
	result, err := repositories.GetAllAnimals(page, 10, "api/v1/animal/"+ord, ord)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAnimalDetailBySlug(c echo.Context) error {
	slug := c.Param("slug")
	result, err := repositories.GetAnimalDetailBySlug("api/v1/animal/detail/"+slug, slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
