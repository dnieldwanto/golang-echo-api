package dto

type CreateCategoryDto struct {
	CategoryName string `json:"categoryName" validate:"required"`
}

type SearchCategoryDto struct {
	ID    uint   `json:"id" query:"id"`
	Name  string `json:"name" query:"name"`
	From  string `json:"from" query:"from"`
	To    string `json:"to" query:"to"`
	Page  int    `json:"page" query:"page" validate:"required"`
	Limit int    `json:"limit" query:"limit" validate:"required"`
}

type CategoryResponseDto struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"categoryName"`
	CreatedAt    string `json:"createdAt"`
	CreatedBy    string `json:"createdBy"`
}
