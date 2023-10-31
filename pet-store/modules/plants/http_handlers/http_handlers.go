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
