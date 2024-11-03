package wishlist

import (
	"go-wishlist-api/dto"
	"go-wishlist-api/handlers/base"
	"go-wishlist-api/usecases/wishlist"

	"github.com/labstack/echo/v4"
)

type wishlistHandler struct {
	usecase wishlist.WishlistUsecase
}

func NewWishlistHandler(u wishlist.WishlistUsecase) *wishlistHandler {
	return &wishlistHandler{
		usecase: u,
	}
}

func (handler wishlistHandler) GetAll(c echo.Context) error {
	result, err := handler.usecase.GetAll()

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponse(c, result)
}

func (handler wishlistHandler) Create(c echo.Context) error {
	input := dto.WishlistInput{}

	if err := c.Bind(&input); err != nil {
		return base.BadRequestResponse(c, err)
	}

	result, err := handler.usecase.Create(input)

	if err != nil {
		return base.ErrorResponse(c, err)
	}

	return base.SuccesResponse(c, result)
}
