package models

type Wishlist struct {
	Base
	Title      string `json:"title"`
	IsAchieved bool   `json:"is_achieved" gorm:"default:false"`
}
