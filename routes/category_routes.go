package routes

import (
	"golang-echo-api/domains"
	"golang-echo-api/dto"
	"golang-echo-api/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type categoryRoutes struct {
	categoryService domains.CategoryService
}

func NewCategoryRoutes(group *echo.Group, categoryService domains.CategoryService) {
	categoryRoutes := categoryRoutes{
		categoryService: categoryService,
	}

	categoryApi := group.Group("/categories")
	categoryApi.POST("", categoryRoutes.CreateCategory)
	categoryApi.PUT("/:id", categoryRoutes.UpdateCategory)
	categoryApi.DELETE("/:id", categoryRoutes.DeleteCategory)
	categoryApi.GET("", categoryRoutes.FindCategory)
}

func (categoryRoutes categoryRoutes) CreateCategory(c echo.Context) error {
	var request dto.CreateCategoryDto
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, utils.GenerateResponse(http.StatusBadRequest, err.Error(), nil))
	}

	fails := utils.Validate(request)
	if len(fails) > 0 {
		return c.JSON(http.StatusBadRequest, utils.GenerateResponse(http.StatusBadRequest, "Bad Request", fails))
	}

	category, err := categoryRoutes.categoryService.Create(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.GenerateResponse(http.StatusInternalServerError, err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, utils.GenerateResponse(http.StatusCreated, "Create Category Successfully", category))
}

func (categoryRoutes categoryRoutes) UpdateCategory(c echo.Context) error {
	var request dto.CreateCategoryDto
	id, _ := strconv.Atoi(c.Param("id"))
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, utils.GenerateResponse(http.StatusBadRequest, err.Error(), nil))
	}

	fails := utils.Validate(request)
	if len(fails) > 0 {
		return c.JSON(http.StatusBadRequest, utils.GenerateResponse(http.StatusBadRequest, "Bad Request", fails))
	}

	category, err := categoryRoutes.categoryService.Update(id, &request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.GenerateResponse(http.StatusInternalServerError, err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, utils.GenerateResponse(http.StatusCreated, "Update Category Successfully", category))
}

func (categoryRoutes categoryRoutes) DeleteCategory(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := categoryRoutes.categoryService.Delete(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.GenerateResponse(http.StatusInternalServerError, err.Error(), nil))
	}

	return c.JSON(http.StatusCreated, utils.GenerateResponse(http.StatusCreated, "Delete Category Successfully", nil))
}

func (categoryRoutes categoryRoutes) FindCategory(c echo.Context) error {
	var request dto.SearchCategoryDto
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &request); err != nil {
		return c.JSON(http.StatusBadRequest, utils.GenerateResponse(http.StatusBadRequest, err.Error(), nil))
	}

	fails := utils.Validate(request)
	if len(fails) > 0 {
		return c.JSON(http.StatusBadRequest, utils.GenerateResponse(http.StatusBadRequest, "Bad Request", fails))
	}

	category, err := categoryRoutes.categoryService.FindCategory(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.GenerateResponse(http.StatusInternalServerError, err.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.GenerateResponseV2(http.StatusOK, category))
}
