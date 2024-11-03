package routes

import (
	handler "go-wishlist-api/handlers/wishlist"
	repository "go-wishlist-api/repositories/wishlist"
	usecase "go-wishlist-api/usecases/wishlist"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitWishlistRoute(e *echo.Echo, db *gorm.DB) {
	wishlistRepo := repository.NewWishlistRepository(db)
	wishlistUsecase := usecase.NewWishlistUsecase(wishlistRepo)
	wishlistHandler := handler.NewWishlistHandler(wishlistUsecase)

	wishlist := e.Group("/api/v1/wishlist")
	wishlist.POST("", wishlistHandler.Create)
	wishlist.GET("", wishlistHandler.GetAll)
}
