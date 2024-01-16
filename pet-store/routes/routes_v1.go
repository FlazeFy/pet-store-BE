package routes

import (
	"net/http"
	animalhandlers "pet-store/modules/animals/http_handlers"
	ctlghandlers "pet-store/modules/catalogs/http_handlers"
	pplhandlers "pet-store/modules/people/http_handlers"
	planthandlers "pet-store/modules/plants/http_handlers"
	stshandlers "pet-store/modules/stats/http_handlers"
	syshandlers "pet-store/modules/systems/http_handlers"
	wohandlers "pet-store/modules/warehouse/http_handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitV1() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to Pet-Store")
	})

	// =============== Public routes ===============

	// Dictionary
	e.GET("api/v1/dct/:type", syshandlers.GetDictionaryByType)
	e.POST("api/v1/dct", syshandlers.PostDct)
	e.DELETE("api/v1/dct/destroy/:id", syshandlers.HardDelDctById)

	// Tag
	e.GET("api/v1/tag", syshandlers.GetAllTag)
	e.POST("api/v1/tag", syshandlers.PostTag)
	e.DELETE("api/v1/tag/destroy/:id", syshandlers.HardDelTagById)

	// Animals
	e.GET("api/v1/animal/:order", animalhandlers.GetAllAnimals)
	e.GET("api/v1/animal/detail/:slug", animalhandlers.GetAnimalDetailBySlug)
	e.POST("api/v1/animal", animalhandlers.PostAnimal)
	e.DELETE("api/v1/animal/destroy/:id", animalhandlers.HardDelAnimalBySlug)
	e.DELETE("api/v1/animal/by/:id", animalhandlers.SoftDelAnimalBySlug)

	// Plants
	e.GET("api/v1/plant/:order", planthandlers.GetAllPlants)
	e.GET("api/v1/plant/detail/:slug", planthandlers.GetPlantDetailBySlug)
	e.POST("api/v1/plant", planthandlers.PostPlant)
	e.DELETE("api/v1/plant/destroy/:id", planthandlers.HardDelPlantBySlug)
	e.DELETE("api/v1/plant/by/:id", planthandlers.SoftDelPlantBySlug)

	// Catalog (Animal & Plants)
	e.GET("api/v1/catalog/:order", ctlghandlers.GetAllCatalogs)

	// Carts
	e.GET("api/v1/cart/:order", ctlghandlers.GetMyCart)
	e.POST("api/v1/cart", ctlghandlers.PostCart)
	e.DELETE("api/v1/cart/destroy/:id", ctlghandlers.HardDelCartById)
	e.PUT("api/v1/cart/by/:id", ctlghandlers.UpdateCartById)

	// Shelf
	e.GET("api/v1/shelf/:order", wohandlers.GetAllActiveShelf)
	e.GET("api/v1/dump/shelf/:order", wohandlers.GetAllTrashShelf)

	// Customer
	e.GET("api/v1/customer/:view", pplhandlers.GetAllCustomer)
	e.GET("api/v1/customer/my/profile", pplhandlers.GetMyProfile)

	// Stats
	e.GET("api/v1/stats/animalgender/:ord", stshandlers.GetTotalAnimalByGender)
	e.GET("api/v1/stats/customerisnotif/:ord", stshandlers.GetTotalCustomerByIsNotif)
	e.GET("api/v1/stats/cartispaid/:ord", stshandlers.GetTotalCartIsPaid)
	e.GET("api/v1/stats/shelfisactive/:ord", stshandlers.GetTotalShelfIsActive)

	// =============== Private routes (Admin) ===============

	return e
}
