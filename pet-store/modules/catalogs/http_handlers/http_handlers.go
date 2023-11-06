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
