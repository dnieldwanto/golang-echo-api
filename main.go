package main

import (
	"golang-echo-api/config"
	"golang-echo-api/migrate"
	"golang-echo-api/repository"
	"golang-echo-api/routes"
	"golang-echo-api/services"
	"log"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db := config.NewDB()
	db.AutoMigrate(migrate.LoadAllModels()...)
	api := e.Group("/api/v1")

	// create directory
	err := os.MkdirAll("uploads", os.ModePerm)
	if err != nil {
		log.Fatalln(err.Error())
	}

	// repository
	categoryRepository := repository.NewCategoryRepository(db)
	wallpaperRepository := repository.NewWallpaperRepository(db)

	// servicee
	categoryService := services.NewCategoryService(db, categoryRepository)
	wallpaperService := services.NewWallpaperService(db, wallpaperRepository, categoryRepository)

	// routes
	routes.NewCategoryRoutes(api, categoryService)
	routes.NewWallpaperRoutes(api, wallpaperService)
	e.Logger.Fatal(e.Start(":8080"))
}
