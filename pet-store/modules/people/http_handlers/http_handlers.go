package httphandlers

import (
	"net/http"
	"pet-store/modules/people/repositories"
	"strconv"

	"github.com/labstack/echo"
)

func GetAllCustomer(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	view := c.Param("view")
	result, err := repositories.GetAllCustomer(page, 10, "api/v1/customer/"+view, view)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetMyProfile(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	result, err := repositories.GetMyProfile("api/v1/customer/my/profile", token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
func GetMyProfileAdmin(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	result, err := repositories.GetMyProfileAdmin("api/v1/admin/my/profile", token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllDoctorSchedule(c echo.Context) error {
	result, err := repositories.GetAllDoctorSchedule()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func HardDelDoctorBySlug(c echo.Context) error {
	slug := c.Param("slug")
	result, err := repositories.HardDelDoctorBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func HardDelCustomerBySlug(c echo.Context) error {
	slug := c.Param("slug")
	result, err := repositories.HardDelCustomerBySlug(slug)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
