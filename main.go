package main

import (
	"go-wishlist-api/database"
	"go-wishlist-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	db, _ := database.InitDB()
	database.MigrateDB(db)

	e := echo.New()

	routes.InitWishlistRoute(e, db)

	e.Logger.Fatal(e.Start(":8000"))
}
