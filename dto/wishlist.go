package dto

import "go-wishlist-api/models"

type WishlistInput struct {
	Title      string `json:"title"`
	IsAchieved bool   `json:"is_achieved"`
}

func (w WishlistInput) ToModel() models.Wishlist {
	return models.Wishlist{
		Title:      w.Title,
		IsAchieved: w.IsAchieved,
	}
}
