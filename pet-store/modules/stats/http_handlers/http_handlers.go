package httphandlers

import (
	"net/http"
	"pet-store/modules/stats/repositories"

	"github.com/labstack/echo"
)

func GetTotalAnimalByGender(c echo.Context) error {
	ord := c.Param("ord")
	view := "animals_gender"
	table := "animals"

	result, err := repositories.GetTotalStats("api/v1/stats/animalgender/"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalCustomerByIsNotif(c echo.Context) error {
	ord := c.Param("ord")
	view := "is_notifable"
	table := "customers"

	result, err := repositories.GetTotalStats("api/v1/stats/customerisnotif/"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalCartIsPaid(c echo.Context) error {
	ord := c.Param("ord")
	view := "is_paid"
	table := "carts"

	result, err := repositories.GetTotalStats("api/v1/stats/cartispaid/"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetTotalShelfIsActive(c echo.Context) error {
	ord := c.Param("ord")
	view := "is_active"
	table := "shelfs"

	result, err := repositories.GetTotalStats("api/v1/stats/shelfisactive/"+ord, ord, view, table)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
