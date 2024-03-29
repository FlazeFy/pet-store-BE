package routes

import (
	"net/http"
	middlewares "pet-store/middlewares/jwt"
	animalhandlers "pet-store/modules/animals/http_handlers"
	authhandlers "pet-store/modules/auth/http_handlers"
	ctlghandlers "pet-store/modules/catalogs/http_handlers"
	gdshandlers "pet-store/modules/goods/http_handlers"
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

	// Auth
	e.POST("api/v1/login", authhandlers.PostLoginUser)
	e.POST("api/v1/register", authhandlers.PostRegister)

	// Dictionary
	e.GET("api/v1/dct/:type", syshandlers.GetDictionaryByType)
	e.POST("api/v1/dct", syshandlers.PostDct)
	e.DELETE("api/v1/dct/destroy/:id", syshandlers.HardDelDctById)

	// Tag
	e.GET("api/v1/tag", syshandlers.GetAllTag)
	e.POST("api/v1/tag", syshandlers.PostTag)
	e.DELETE("api/v1/tag/destroy/:id", syshandlers.HardDelTagById)

	// Feedbacks
	e.POST("api/v1/feedbacks", syshandlers.PostFeedback)
	e.GET("api/v1/feedbacks/:ord_obj/:ord", syshandlers.GetAllFeedback)

	// Animals
	e.GET("api/v1/animal/:order", animalhandlers.GetAllAnimals)
	e.POST("api/v1/animal", animalhandlers.PostAnimal)

	// Plants
	e.GET("api/v1/plant/:order", planthandlers.GetAllPlants)
	e.POST("api/v1/plant", planthandlers.PostPlant)

	// Catalog (Animal & Plants)
	e.GET("api/v1/catalog/:order", ctlghandlers.GetAllCatalogs)

	// Goods
	e.GET("api/v1/goods/:order", gdshandlers.GetAllGoods)
	e.DELETE("api/v1/goods/destroy/:id", gdshandlers.HardDelGoodBySlug)

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
	e.DELETE("api/v1/customer/destroy/:slug", pplhandlers.HardDelCustomerBySlug)

	// Doctor
	e.GET("api/v1/doctor/schedule", pplhandlers.GetAllDoctorSchedule)
	e.GET("api/v1/doctor/data/:ord", pplhandlers.GetAllDoctors)
	e.DELETE("api/v1/doctor/destroy/:slug", pplhandlers.HardDelDoctorBySlug)

	// Stats
	e.GET("api/v1/stats/animalgender/:ord", stshandlers.GetTotalAnimalByGender)
	e.GET("api/v1/stats/customerisnotif/:ord", stshandlers.GetTotalCustomerByIsNotif)
	e.GET("api/v1/stats/cartispaid/:ord", stshandlers.GetTotalCartIsPaid)
	e.GET("api/v1/stats/shelfisactive/:ord", stshandlers.GetTotalShelfIsActive)
	e.GET("api/v1/stats/goodscategory/:ord", stshandlers.GetTotalGoodsCategory)
	e.GET("api/v1/stats/wishtype/:ord", stshandlers.GetTotalWishlistType)
	e.GET("api/v1/stats/docready/:ord", stshandlers.GetTotalDoctorReady)

	// =============== Private routes ===============

	// Auth
	e.POST("api/v1/logout", authhandlers.SignOut, middlewares.CustomJWTAuth)
	e.GET("api/v1/check", authhandlers.CheckRole, middlewares.CustomJWTAuth)

	// Animals
	e.GET("api/v1/animal/detail/:slug", animalhandlers.GetAnimalDetailBySlug, middlewares.CustomJWTAuth)
	e.DELETE("api/v1/animal/destroy/:slug", animalhandlers.HardDelAnimalBySlug, middlewares.CustomJWTAuth)
	e.DELETE("api/v1/animal/by/:slug", animalhandlers.SoftDelAnimalBySlug, middlewares.CustomJWTAuth)
	e.POST("api/v1/animal/recover/:slug", animalhandlers.RecoverAnimalBySlug, middlewares.CustomJWTAuth)

	// Plants
	e.GET("api/v1/plant/detail/:slug", planthandlers.GetPlantDetailBySlug, middlewares.CustomJWTAuth)
	e.DELETE("api/v1/plant/destroy/:slug", planthandlers.HardDelPlantBySlug, middlewares.CustomJWTAuth)
	e.DELETE("api/v1/plant/by/:slug", planthandlers.SoftDelPlantBySlug, middlewares.CustomJWTAuth)
	e.POST("api/v1/plant/recover/:slug", planthandlers.RecoverPlantBySlug, middlewares.CustomJWTAuth)

	// Customer
	e.GET("api/v1/customer/my/profile", pplhandlers.GetMyProfile, middlewares.CustomJWTAuth)

	// Admin
	e.GET("api/v1/admin/my/profile", pplhandlers.GetMyProfileAdmin, middlewares.CustomJWTAuth)

	// Catalog (Animal & Plants)
	e.GET("api/v1/catalog/wishlist/my/:order", ctlghandlers.GetMyWishlist, middlewares.CustomJWTAuth)
	e.GET("api/v1/catalog/wishlist/check/:type/:slug", ctlghandlers.GetCheckWishlist, middlewares.CustomJWTAuth)
	e.POST("api/v1/catalog/wishlist/add", ctlghandlers.PostWishlist, middlewares.CustomJWTAuth)
	e.DELETE("api/v1/catalog/wishlist/destroy/:id", ctlghandlers.HardDelWishlistById, middlewares.CustomJWTAuth)

	return e
}
