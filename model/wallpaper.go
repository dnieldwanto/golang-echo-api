package model

import (
	"database/sql"
)

type Wallpaper struct {
	ID          uint           `json:"id" gorm:"primaryKey;column:id"`
	Title       string         `json:"title" gorm:"not null;column:title"`
	Description string         `json:"description" gorm:"not null;column:description"`
	Filename    string         `json:"filename" gorm:"not null;column:filename"`
	IsActive    bool           `json:"isActive" gorm:"column:is_active"`
	CategoryId  uint           `json:"categoryId" gorm:"not null;column:category_id"`
	CreatedAt   sql.NullTime   `json:"createdAt" gorm:"column:created_at"`
	CreatedBy   string         `json:"createdBy" gorm:"column:created_by"`
	UpdatedAt   sql.NullTime   `json:"updatedAt" gorm:"column:updated_at"`
	UpdatedBy   sql.NullString `json:"updatedBy" gorm:"column:updated_by"`
}

type WallpaperPage struct {
	Wallpapers []Wallpaper
	Count      uint
}
