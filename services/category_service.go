package services

import (
	"database/sql"
	"golang-echo-api/domains"
	"golang-echo-api/dto"
	"golang-echo-api/model"
	"time"

	"gorm.io/gorm"
)

type categoryService struct {
	db                 *gorm.DB
	categoryRepository domains.CategoryRepository
}

// FindCategory implements domains.CategoryService.
func (service *categoryService) FindCategory(request dto.SearchCategoryDto) ([]dto.CategoryResponseDto, error) {
	var categoryDto []dto.CategoryResponseDto

	if request.From != "" && request.To != "" {
		layout := "2006-01-02"

		parseFrom, err := time.Parse(layout, request.From)
		if err != nil {
			return nil, err
		}

		parseTo, err := time.Parse(layout, request.To)
		if err != nil {
			return nil, err
		}

		startOfDay := parseFrom.Format("2006-01-02 15:04:05")
		endOfDay := parseTo.Add(24*time.Hour - time.Nanosecond).Format("2006-01-02 15:04:05")

		request.From = startOfDay
		request.To = endOfDay
	}
	category, err := service.categoryRepository.FindAll(request)
	if err != nil {
		return nil, err
	}

	for _, data := range category {
		categoryDto = append(categoryDto, dto.CategoryResponseDto{
			ID:           data.ID,
			CategoryName: data.CategoryName,
			CreatedAt:    data.CreatedAt.Time.String(),
			CreatedBy:    data.CreatedBy,
		})
	}

	return categoryDto, nil
}

// Delete implements domains.CategoryService.
func (service *categoryService) Delete(id int) error {
	return service.db.Transaction(func(tx *gorm.DB) error {
		return service.categoryRepository.Delete(tx, id)
	})
}

// Update implements domains.CategoryService.
func (service *categoryService) Update(id int, request *dto.CreateCategoryDto) (result dto.CategoryResponseDto, err error) {
	err = service.db.Transaction(func(tx *gorm.DB) error {
		category, err := service.categoryRepository.FindById(id)
		if err != nil {
			return err
		}

		category.CategoryName = request.CategoryName
		if err := service.categoryRepository.Update(tx, &category); err != nil {
			return err
		}

		result.ID = category.ID
		result.CategoryName = category.CategoryName
		result.CreatedAt = category.CreatedAt.Time.String()
		result.CreatedBy = category.CreatedBy
		return nil
	})

	if err != nil {
		return dto.CategoryResponseDto{}, err
	}

	return result, nil
}

// Create implements domains.CategoryService.
func (service *categoryService) Create(request *dto.CreateCategoryDto) (result dto.CategoryResponseDto, err error) {
	err = service.db.Transaction(func(tx *gorm.DB) error {
		category := &model.Category{
			CategoryName: request.CategoryName,
			IsActive:     true,
			CreatedAt:    sql.NullTime{Valid: true, Time: time.Now()},
			CreatedBy:    "ADMIN",
		}
		if err := service.categoryRepository.Create(tx, category); err != nil {
			return err
		}

		result.ID = category.ID
		result.CategoryName = category.CategoryName
		result.CreatedAt = category.CreatedAt.Time.String()
		result.CreatedBy = category.CreatedBy
		return nil
	})

	if err != nil {
		return dto.CategoryResponseDto{}, err
	}

	return result, nil
}

func NewCategoryService(db *gorm.DB, categoryRepository domains.CategoryRepository) domains.CategoryService {
	return &categoryService{
		db:                 db,
		categoryRepository: categoryRepository,
	}
}
