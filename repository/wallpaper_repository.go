package repository

import (
	"golang-echo-api/domains"
	"golang-echo-api/model"

	"gorm.io/gorm"
)

type wallpaperRepository struct {
	db *gorm.DB
}

// Create implements domains.WallpaperRepository.
func (repository *wallpaperRepository) Create(tx *gorm.DB, wallpaper *model.Wallpaper) error {
	if err := tx.Create(wallpaper).Error; err != nil {
		return err
	}
	return nil
}

func NewWallpaperRepository(db *gorm.DB) domains.WallpaperRepository {
	return &wallpaperRepository{
		db: db,
	}
}
