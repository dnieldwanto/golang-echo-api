package main

import (
	"golang-echo-api/config"
	"golang-echo-api/migrate"
	"golang-echo-api/repository"
	"golang-echo-api/routes"
	"golang-echo-api/services"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db := config.NewDB()
	db.AutoMigrate(migrate.LoadAllModels()...)
	api := e.Group("/api/v1")

	// repository
	categoryRepository := repository.NewCategoryRepository(db)

	// servicee
	categoryService := services.NewCategoryService(db, categoryRepository)

	// routes
	routes.NewCategoryRoutes(api, categoryService)
	e.Logger.Fatal(e.Start(":8080"))
}
