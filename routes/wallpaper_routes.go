package routes

import (
	"golang-echo-api/domains"
	"golang-echo-api/dto"
	"golang-echo-api/utils"
	"net/http"
	"strconv"

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
	wallpaperApi.PUT("/:id", wallpaperRoutes.UpdateWallpaper)
	wallpaperApi.GET("", wallpaperRoutes.FindAllData)
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
		return c.JSON(http.StatusInternalServerError, utils.GenerateResponse(http.StatusInternalServerError, err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, utils.GenerateResponse(http.StatusCreated, "Create Wallpaper Successfully", wallpaper))
}

func (wallpaperRoutes wallpaperRoutes) UpdateWallpaper(c echo.Context) error {
	var request dto.CreateWallpaperDto
	file, _ := c.FormFile("file")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.GenerateResponse(http.StatusBadRequest, err.Error(), nil))
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, utils.GenerateResponse(http.StatusBadRequest, err.Error(), nil))
	}

	fails := utils.Validate(request)
	if len(fails) > 0 {
		return c.JSON(http.StatusBadRequest, utils.GenerateResponse(http.StatusBadRequest, "Bad Request", fails))
	}

	wallpaper, err := wallpaperRoutes.wallpaperService.Update(id, file, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.GenerateResponse(http.StatusInternalServerError, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.GenerateResponse(http.StatusOK, "Update Wallpaper Successfully", wallpaper))
}

func (wallpaperRoutes wallpaperRoutes) FindAllData(c echo.Context) error {
	var request dto.FindWallpaperDto

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, utils.GenerateResponse(http.StatusBadRequest, err.Error(), nil))
	}

	wallpapers, err := wallpaperRoutes.wallpaperService.FindAll(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.GenerateResponse(http.StatusInternalServerError, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.GenerateResponse(http.StatusOK, "OK", wallpapers))
}
