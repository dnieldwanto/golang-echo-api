package repository

import (
	"golang-echo-api/domains"
	"golang-echo-api/dto"
	"golang-echo-api/model"

	"gorm.io/gorm"
)

type wallpaperRepository struct {
	db *gorm.DB
}

// FindAll implements domains.WallpaperRepository.
func (repository *wallpaperRepository) FindAll(request dto.FindWallpaperDto) (model.WallpaperPage, error) {
	var countQuery = "SELECT COUNT(w.*) as total FROM wallpaper.wallpapers w INNER JOIN wallpaper.categories c ON c.id = w.category_id WHERE w.is_active IS TRUE"
	var query = "SELECT w.* FROM wallpaper.wallpapers w INNER JOIN wallpaper.categories c ON c.id = w.category_id WHERE w.is_active IS TRUE"
	var condition string
	var count uint
	var sortBy string = " ORDER BY w.ID ASC"
	var wallpapers []model.Wallpaper
	params := []interface{}{}

	if request.Category != "" {
		condition += " AND c.category_name ILIKE ?"
		params = append(params, "%"+request.Category+"%")
	}

	if request.Search != "" {
		condition += " AND (w.title ILIKE ? OR w.description ILIKE ?)"
		params = append(params, "%"+request.Search+"%", "%"+request.Search+"%")
	}

	if request.ID != 0 {
		condition += " AND w.id = ?"
		params = append(params, request.ID)
	}

	if request.SortBy != "" {
		switch request.SortBy {
		case "title:ASC":
			sortBy = " ORDER BY w.title ASC"
		case "title:DESC":
			sortBy = " ORDER BY w.title DESC"
		}
	}

	countQuery += condition
	repository.db.Raw(countQuery, params...).Scan(&count)

	offset := (request.Page - 1) * request.Limit
	query = query + condition + sortBy + " LIMIT ? OFFSET ?"
	params = append(params, request.Limit, offset)

	repository.db.Raw(query, params...).Scan(&wallpapers)

	result := model.WallpaperPage{
		Wallpapers: wallpapers,
		Count:      count,
	}

	return result, nil
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
