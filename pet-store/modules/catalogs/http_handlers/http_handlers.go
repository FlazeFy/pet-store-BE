package httphandlers

import (
	"net/http"
	"pet-store/modules/catalogs/models"
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

func GetMyWishlist(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	ord := c.Param("order")
	result, err := repositories.GetMyWishlist(page, 10, "api/v1/catalog/wishlist/my/"+ord, ord, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetCheckWishlist(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	slug := c.Param("slug")
	types := c.Param("type")
	result, err := repositories.GetCheckWishlist(token, slug, types)
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

func PostWishlist(c echo.Context) error {
	var obj models.PostWishlist
	token := c.Request().Header.Get("Authorization")

	// Data
	obj.CatalogType = c.FormValue("catalog_type")
	obj.CatalogId = c.FormValue("catalog_id")

	result, err := repositories.PostWishlist(obj, token)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
