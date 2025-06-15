package dto

type CreateWallpaperDto struct {
	Title       string `form:"title" validate:"required"`
	Description string `form:"description" validate:"required"`
	Filename    string
	CategoryId  uint `form:"categoryId" validate:"required"`
}

type WallpaperDto struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Filename    string `json:"filename"`
	Category    string `json:"category"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type FindWallpaperDto struct {
	ID       int    `query:"id"`
	Category string `query:"category"`
	Search   string `query:"search"`
	Page     int    `query:"page"`
	Limit    int    `query:"limit"`
	SortBy   string `query:"sortBy"`
}
