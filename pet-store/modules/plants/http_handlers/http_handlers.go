package httphandlers

import (
	"net/http"
	"pet-store/modules/plants/repositories"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllPlants(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	ord := c.Param("order")
	result, err := repositories.GetAllPlants(page, 10, "api/v1/plant/"+ord, ord)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetPlantDetailBySlug(c echo.Context) error {
	slug := c.Param("slug")
	result, err := repositories.GetPlantDetailBySlug("api/v1/plant/detail/"+slug, slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func HardDelPlantBySlug(c echo.Context) error {
	slug := c.Param("slug")
	result, err := repositories.HardDelPlantBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func SoftDelPlantBySlug(c echo.Context) error {
	slug := c.Param("slug")
	result, err := repositories.SoftDelPlantBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostPlant(c echo.Context) error {
	result, err := repositories.PostPlant(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
