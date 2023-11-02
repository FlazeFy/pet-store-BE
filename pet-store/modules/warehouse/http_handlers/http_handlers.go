package httphandlers

import (
	"net/http"
	"pet-store/modules/warehouse/repositories"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllActiveShelf(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	ord := c.Param("order")
	result, err := repositories.GetAllShelf(page, 10, "api/v1/shelf/"+ord, "active", ord)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllTrashShelf(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	ord := c.Param("order")
	result, err := repositories.GetAllShelf(page, 10, "api/v1/dump/shelf/"+ord, "trash", ord)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
