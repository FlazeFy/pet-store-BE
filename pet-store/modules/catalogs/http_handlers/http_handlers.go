package httphandlers

import (
	"net/http"
	"pet-store/modules/catalogs/repositories"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllCatalogs(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	ord := c.Param("order")
	result, err := repositories.GetAllCatalogs(page, 10, "api/v1/catalog/"+ord, ord)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetMyCart(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	ord := c.Param("order")
	result, err := repositories.GetMyCart(page, 10, "api/v1/cart/"+ord, ord)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func HardDelCartById(c echo.Context) error {
	id := c.Param("id")
	result, err := repositories.HardDelCartById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostCart(c echo.Context) error {
	result, err := repositories.PostCart(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateCartById(c echo.Context) error {
	id := c.Param("id")

	result, err := repositories.UpdateCartById(id, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
