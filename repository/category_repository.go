package repository

import (
	"golang-echo-api/domains"
	"golang-echo-api/dto"
	"golang-echo-api/model"
	"log"

	"gorm.io/gorm"
)

type categoryRepository struct {
	db *gorm.DB
}

// Create implements domains.CategoryRepository.
func (repository *categoryRepository) Create(tx *gorm.DB, category *model.Category) error {
	if err := tx.Create(category).Error; err != nil {
		return err
	}
	return nil
}

// Delete implements domains.CategoryRepository.
func (c *categoryRepository) Delete(tx *gorm.DB, id int) error {
	var category model.Category
	result := tx.First(&category, id)
	if result.Error != nil {
		return result.Error
	}

	category.IsActive = false
	if err := tx.Save(category).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements domains.CategoryRepository.
func (repository *categoryRepository) FindAll(request dto.SearchCategoryDto) (result []model.Category, err error) {
	var query = "SELECT * FROM wallpaper.categories c WHERE c.is_active IS TRUE"
	params := []interface{}{}

	if request.ID != 0 {
		query += " AND c.id = ?"
		params = append(params, request.ID)
	}

	if request.Name != "" {
		query += " AND c.category_name ILIKE ?"
		params = append(params, "%"+request.Name+"%")
	}

	if request.From != "" && request.To != "" {
		query += " AND c.created_at BETWEEN ? AND ?"
		params = append(params, request.From, request.To)
	}

	offset := (request.Page - 1) * request.Limit
	query += " ORDER BY c.id ASC LIMIT ? OFFSET ?"
	params = append(params, request.Limit, offset)

	log.Println("Final Query =>", query)

	repository.db.Raw(query, params...).Scan(&result)
	return result, nil
}

// FindById implements domains.CategoryRepository.
func (repository *categoryRepository) FindById(id int) (model.Category, error) {
	var category model.Category
	result := repository.db.First(&category, id)
	if result.Error != nil {
		return model.Category{}, result.Error
	}

	return category, nil
}

// Update implements domains.CategoryRepository.
func (c *categoryRepository) Update(tx *gorm.DB, category *model.Category) error {
	if err := tx.Save(&category).Error; err != nil {
		return err
	}
	return nil
}

func NewCategoryRepository(db *gorm.DB) domains.CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}
