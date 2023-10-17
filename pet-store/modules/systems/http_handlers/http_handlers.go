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
