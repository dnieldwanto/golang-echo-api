package routes

import (
	"fmt"
	"golang-echo-api/domains"
	"golang-echo-api/dto"
	"golang-echo-api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type wallpaperRoutes struct {
	wallpaperService domains.WallpaperService
}

func NewWallpaperRoutes(group *echo.Group, wallpaperService domains.WallpaperService) {
	wallpaperRoutes := wallpaperRoutes{
		wallpaperService: wallpaperService,
	}

	wallpaperApi := group.Group("/wallpapers")
	wallpaperApi.POST("", wallpaperRoutes.CreateWallpaper)
}

func (wallpaperRoutes wallpaperRoutes) CreateWallpaper(c echo.Context) error {
	var request dto.CreateWallpaperDto
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, utils.GenerateResponse(http.StatusBadRequest, err.Error(), nil))
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.GenerateResponse(http.StatusBadRequest, err.Error(), nil))
	}

	fails := utils.Validate(request)
	if len(fails) > 0 {
		return c.JSON(http.StatusBadRequest, utils.GenerateResponse(http.StatusBadRequest, "Bad Request", fails))
	}

	wallpaper, err := wallpaperRoutes.wallpaperService.Create(file, request)
	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, utils.GenerateResponse(http.StatusInternalServerError, err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, utils.GenerateResponse(http.StatusCreated, "Create Wallpaper Successfully", wallpaper))
}
