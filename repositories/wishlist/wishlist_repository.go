package wishlist

import (
	"go-wishlist-api/dto"
	"go-wishlist-api/models"
)

type WishlistRepository interface {
	GetAll() ([]models.Wishlist, error)
	Create(input dto.WishlistInput) (models.Wishlist, error)
}
