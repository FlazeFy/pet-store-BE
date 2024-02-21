package httphandlers

import (
	"net/http"
	"pet-store/modules/goods/repositories"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllGoods(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	ord := c.Param("order")
	result, err := repositories.GetAllGoods(page, 10, "api/v1/goods/"+ord, ord)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
