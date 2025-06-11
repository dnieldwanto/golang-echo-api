package repository

import (
	"golang-echo-api/domains"
	"golang-echo-api/model"

	"gorm.io/gorm"
)

type wallpaperRepository struct {
	db *gorm.DB
}

// FindById implements domains.WallpaperRepository.
func (repository *wallpaperRepository) FindById(id int) (result model.Wallpaper, err error) {
	query := "SELECT * FROM wallpaper.wallpapers w WHERE w.id = ?"
	params := []interface{}{}

	params = append(params, id)
	res := repository.db.Raw(query, params...).Scan(&result)
	if res.Error != nil {
		return model.Wallpaper{}, res.Error
	}

	return result, nil
}

// Update implements domains.WallpaperRepository.
func (repository *wallpaperRepository) Update(tx *gorm.DB, wallpaper *model.Wallpaper) error {
	if err := tx.Save(&wallpaper).Error; err != nil {
		return err
	}
	return nil
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
