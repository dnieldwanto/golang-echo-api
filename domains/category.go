package domains

import (
	"golang-echo-api/dto"
	"golang-echo-api/model"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll(request dto.SearchCategoryDto) ([]model.Category, error)
	FindById(id int) (model.Category, error)
	Create(tx *gorm.DB, category *model.Category) error
	Update(tx *gorm.DB, category *model.Category) error
	Delete(tx *gorm.DB, id int) error
}

type CategoryService interface {
	Create(request *dto.CreateCategoryDto) (dto.CategoryResponseDto, error)
	Update(id int, request *dto.CreateCategoryDto) (dto.CategoryResponseDto, error)
	Delete(id int) error
	FindCategory(request dto.SearchCategoryDto) ([]dto.CategoryResponseDto, error)
}
