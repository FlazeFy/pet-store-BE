package httphandlers

import (
	"net/http"
	"pet-store/modules/systems/repositories"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllTag(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	result, err := repositories.GetAllTag(page, 10, "api/v1/tag")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetDictionaryByType(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	dctType := c.Param("type")
	result, err := repositories.GetDictionaryByType(page, 10, "api/v1/dct:"+dctType, dctType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func HardDelTagById(c echo.Context) error {
	id := c.Param("id")
	result, err := repositories.HardDelTagById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostTag(c echo.Context) error {
	result, err := repositories.PostTag(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
