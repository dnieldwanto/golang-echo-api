package services

import (
	"database/sql"
	"errors"
	"golang-echo-api/domains"
	"golang-echo-api/dto"
	"golang-echo-api/model"
	"io"
	"math"
	"mime/multipart"
	"os"
	"time"

	"gorm.io/gorm"
)

type wallpaperService struct {
	db                  *gorm.DB
	wallpaperRepository domains.WallpaperRepository
	categoryRepository  domains.CategoryRepository
}

// FindAll implements domains.WallpaperService.
func (service *wallpaperService) FindAll(request dto.FindWallpaperDto) (dto.Paginate[dto.WallpaperDto], error) {
	if request.Page == 0 {
		request.Page = 1
	}
	if request.Limit == 0 {
		request.Limit = 10
	}
	var wallpaperDto []dto.WallpaperDto

	wallpapers, err := service.wallpaperRepository.FindAll(request)
	if err != nil {
		return dto.Paginate[dto.WallpaperDto]{}, err
	}

	for _, data := range wallpapers.Wallpapers {
		category, err := service.categoryRepository.FindById(int(data.CategoryId))
		if err != nil {
			return dto.Paginate[dto.WallpaperDto]{}, nil
		}

		if category.ID == 0 {
			return dto.Paginate[dto.WallpaperDto]{}, errors.New("Category not found")
		}

		wallpaperDto = append(wallpaperDto, dto.WallpaperDto{
			ID:          data.ID,
			Title:       data.Title,
			Description: data.Description,
			Filename:    data.Filename,
			Category:    category.CategoryName,
			CreatedAt:   data.CreatedAt.Time.String(),
			UpdatedAt:   data.UpdatedAt.Time.String(),
		})
	}

	return dto.Paginate[dto.WallpaperDto]{
		TotalData:   wallpapers.Count,
		TotalPage:   uint(math.Ceil(float64(wallpapers.Count) / float64(request.Limit))),
		CurrentData: uint(len(wallpapers.Wallpapers)),
		CurrentPage: uint(request.Page),
		Content:     wallpaperDto,
	}, nil
}

// Update implements domains.WallpaperService.
func (service *wallpaperService) Update(id int, file *multipart.FileHeader, request dto.CreateWallpaperDto) (result dto.WallpaperDto, err error) {
	err = service.db.Transaction(func(tx *gorm.DB) error {
		wallpaper, err := service.wallpaperRepository.FindById(id)
		if err != nil {
			return err
		}

		if file != nil {
			err := os.Remove("uploads/" + wallpaper.Filename)
			if err != nil {
				return err
			}
			wallpaper.Filename = file.Filename
		}

		wallpaper.Title = request.Title
		wallpaper.Description = request.Description
		wallpaper.CategoryId = request.CategoryId
		wallpaper.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}
		wallpaper.UpdatedBy = sql.NullString{Valid: true, String: "ADMIN"}

		if err := service.wallpaperRepository.Update(tx, &wallpaper); err != nil {
			return err
		}

		category, err := service.categoryRepository.FindById(int(request.CategoryId))
		if err != nil {
			return err
		}

		if category.ID == 0 {
			return errors.New("Category not found")
		}

		if file != nil {
			src, err := file.Open()
			if err != nil {
				return err
			}
			defer src.Close()

			dst, err := os.Create("uploads/" + file.Filename)
			if err != nil {
				return err
			}
			defer dst.Close()

			if _, err = io.Copy(dst, src); err != nil {
				return err
			}
		}

		result.ID = wallpaper.ID
		result.Title = wallpaper.Title
		result.Description = wallpaper.Description
		result.Filename = wallpaper.Filename
		result.Category = category.CategoryName
		result.CreatedAt = wallpaper.CreatedAt.Time.String()
		result.UpdatedAt = wallpaper.UpdatedAt.Time.String()
		return nil
	})

	if err != nil {
		return dto.WallpaperDto{}, err
	}

	return result, nil
}

// Create implements domains.WallpaperService.
func (service *wallpaperService) Create(file *multipart.FileHeader, request dto.CreateWallpaperDto) (result dto.WallpaperDto, err error) {
	err = service.db.Transaction(func(tx *gorm.DB) error {
		request.Filename = file.Filename
		wallpaper := &model.Wallpaper{
			Title:       request.Title,
			Description: request.Description,
			Filename:    request.Filename,
			CategoryId:  request.CategoryId,
			IsActive:    true,
			CreatedAt:   sql.NullTime{Valid: true, Time: time.Now()},
			CreatedBy:   "ADMIN",
		}
		if err := service.wallpaperRepository.Create(tx, wallpaper); err != nil {
			return err
		}

		category, err := service.categoryRepository.FindById(int(request.CategoryId))
		if err != nil {
			return err
		}

		if category.ID == 0 {
			return errors.New("Category not found")
		}

		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		dst, err := os.Create("uploads/" + file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		result.ID = wallpaper.ID
		result.Title = wallpaper.Title
		result.Description = wallpaper.Description
		result.Filename = wallpaper.Filename
		result.Category = category.CategoryName
		result.CreatedAt = wallpaper.CreatedAt.Time.String()
		result.UpdatedAt = wallpaper.UpdatedAt.Time.String()
		return nil
	})

	if err != nil {
		return dto.WallpaperDto{}, err
	}

	return result, nil
}

func NewWallpaperService(db *gorm.DB, wallpaperRepository domains.WallpaperRepository, categoryRepository domains.CategoryRepository) domains.WallpaperService {
	return &wallpaperService{
		db:                  db,
		wallpaperRepository: wallpaperRepository,
		categoryRepository:  categoryRepository,
	}
}
