package wishlist

import (
	"go-wishlist-api/dto"
	"go-wishlist-api/models"

	"gorm.io/gorm"
)

type wishlistRepositoryImpl struct {
	db *gorm.DB
}

func NewWishlistRepository(db *gorm.DB) *wishlistRepositoryImpl {
	return &wishlistRepositoryImpl{
		db: db,
	}
}

func (repository *wishlistRepositoryImpl) GetAll() ([]models.Wishlist, error) {
	wishlists := []models.Wishlist{}
	result := repository.db.Find(&wishlists)

	if result.Error != nil {
		return nil, result.Error
	}

	return wishlists, nil
}

func (repository *wishlistRepositoryImpl) Create(input dto.WishlistInput) (models.Wishlist, error) {
	wishlist := input.ToModel()
	result := repository.db.Create(&wishlist)

	if result.Error != nil {
		return models.Wishlist{}, result.Error
	}

	return wishlist, nil
}
