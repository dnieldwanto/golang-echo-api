package domains

import (
	"golang-echo-api/dto"
	"golang-echo-api/model"
	"mime/multipart"

	"gorm.io/gorm"
)

type WallpaperRepository interface {
	Create(tx *gorm.DB, wallpaper *model.Wallpaper) error
}

type WallpaperService interface {
	Create(file *multipart.FileHeader, request dto.CreateWallpaperDto) (dto.WallpaperDto, error)
}
