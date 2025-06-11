package model

import "database/sql"

type Category struct {
	ID           uint         `json:"id" gorm:"primaryKey;column:id"`
	CategoryName string       `json:"categoryName" gorm:"not null;column:category_name"`
	IsActive     bool         `json:"isActive" gorm:"column:is_active"`
	CreatedAt    sql.NullTime `json:"createdAt" gorm:"column:created_at"`
	CreatedBy    string       `json:"createdBy" gorm:"column:created_by"`
}

type CategoryPage struct {
	Category []Category
	Count    uint
}
