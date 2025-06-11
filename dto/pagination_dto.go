package dto

type Paginate[T any] struct {
	TotalData   uint `json:"totalData"`
	TotalPage   uint `json:"totalPage"`
	CurrentPage uint `json:"currentPage"`
	CurrentData uint `json:"currentData"`
	Content     []T  `json:"content"`
}
