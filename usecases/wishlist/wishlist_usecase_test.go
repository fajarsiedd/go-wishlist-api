package wishlist

import (
	"encoding/json"
	"errors"
	"go-wishlist-api/dto"
	"go-wishlist-api/models"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var (
	wlUsecase   WishlistUsecase
	wlDummyRepo DummyWishlistRepo
)

var defaultTime time.Time = time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC)

type DummyWishlistRepo struct {
	ReturnWishlists []models.Wishlist
	ReturnWishlist  models.Wishlist
	ReturnError     error
}

func (dummy *DummyWishlistRepo) GetAll() ([]models.Wishlist, error) {
	return dummy.ReturnWishlists, dummy.ReturnError
}

func (dummy *DummyWishlistRepo) Create(input dto.WishlistInput) (models.Wishlist, error) {
	return dummy.ReturnWishlist, dummy.ReturnError
}

func setup() {
	wlDummyRepo = DummyWishlistRepo{}

	wlUsecase = NewWishlistUsecase(&wlDummyRepo)
}

func TestGetAll(t *testing.T) {
	setup()

	t.Run("get all wishlist success", func(t *testing.T) {
		wlDummyRepo.ReturnWishlists = []models.Wishlist{
			{
				Base: models.Base{
					ID:        1,
					CreatedAt: defaultTime,
					UpdatedAt: defaultTime,
					DeletedAt: gorm.DeletedAt{Time: defaultTime, Valid: false},
				},
				Title:      "Rumah 2 lantai",
				IsAchieved: false,
			},
		}

		result, err := wlUsecase.GetAll()
		assert.NoError(t, err)
		assert.NotNil(t, result)

		b, err := json.Marshal(result)
		assert.NoError(t, err)
		expected := `[{"id":1,"title":"Rumah 2 lantai","is_achieved":false,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","deleted_at":null}]`
		assert.JSONEq(t, expected, string(b))
	})

	t.Run("get all wishlist error", func(t *testing.T) {
		wlDummyRepo.ReturnError = errors.New("database error")

		_, err := wlUsecase.GetAll()
		assert.Error(t, err)
	})
}

func TestCreate(t *testing.T) {
	setup()

	t.Run("create wishlist success", func(t *testing.T) {
		wlDummyRepo.ReturnWishlist = models.Wishlist{
			Base: models.Base{
				ID:        1,
				CreatedAt: defaultTime,
				UpdatedAt: defaultTime,
				DeletedAt: gorm.DeletedAt{Time: defaultTime, Valid: false},
			},
			Title:      "Rumah 2 lantai",
			IsAchieved: false,
		}

		input := dto.WishlistInput{Title: "Rumah 2 lantai", IsAchieved: false}
		result, err := wlUsecase.Create(input)
		assert.NoError(t, err)
		assert.NotNil(t, result)

		b, err := json.Marshal(result)
		assert.NoError(t, err)
		expected := `{"id":1,"title":"Rumah 2 lantai","is_achieved":false,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z","deleted_at":null}`
		assert.JSONEq(t, expected, string(b))
	})

	t.Run("create wishlist error db", func(t *testing.T) {
		wlDummyRepo.ReturnError = errors.New("database error")

		input := dto.WishlistInput{Title: "Rumah 2 lantai", IsAchieved: false}
		_, err := wlUsecase.Create(input)
		assert.Error(t, err)
	})

	t.Run("create wishlist error title empty", func(t *testing.T) {
		wlDummyRepo.ReturnError = errors.New("email not found")

		input := dto.WishlistInput{IsAchieved: false}
		_, err := wlUsecase.Create(input)
		expectedErr := "title cannot be empty"
		assert.EqualError(t, err, expectedErr)
	})
}
