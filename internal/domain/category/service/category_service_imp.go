package service

import (
	"errors"

	"github.com/gosimple/slug"
	"github.com/phn00dev/golang-news-app/internal/domain/category/dto"
	"github.com/phn00dev/golang-news-app/internal/domain/category/repository"
	"github.com/phn00dev/golang-news-app/internal/models"
)

type categoryServiceImp struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) CategoryService {
	return categoryServiceImp{
		categoryRepo: categoryRepo,
	}
}

func (categoryService categoryServiceImp) GetAllCategories() ([]dto.CategoryResponse, error) {
	categories, err := categoryService.categoryRepo.GetAll()
	if err != nil {
		return nil, err
	}

	categoryResponses := dto.GetAllCategoryResponses(categories)
	return categoryResponses, nil
}

func (categoryService categoryServiceImp) GetOneCategory(categoryID int) (*dto.CategoryResponse, error) {
	if categoryID <= 0 {
		return nil, errors.New("category id not null")
	}

	category, err := categoryService.categoryRepo.GetOne(categoryID)
	if err != nil {
		return nil, err
	}
	categoryResponse := dto.GetOneCategoryResponse(category)
	return categoryResponse, nil
}

func (categoryService categoryServiceImp) CreateCategory(createRequest dto.CreateCategoryRequest) error {

	// get category with category_name
	// eger-de category ady on bar bolsa onda error bermeli
	verifyCategoryName := categoryService.categoryRepo.VerifyCategoryName(createRequest.CategoryName)
	if !verifyCategoryName {
		return errors.New("category name already exists")
	}
	category := models.Category{
		CategoryName:   createRequest.CategoryName,
		CategorySlug:   slug.Make(createRequest.CategoryName),
		CategoryStatus: createRequest.CategoryStatus,
	}
	return categoryService.categoryRepo.Create(category)
}

func (categoryService categoryServiceImp) UpdateCategory(categoryID int, updateRequest dto.UpdateCategoryRequest) error {
	category, err := categoryService.categoryRepo.GetOne(categoryID)
	if err != nil {
		return errors.New("category not found")
	}
	if updateRequest.CategoryName != "" && updateRequest.CategoryName != category.CategoryName {
		verifyCategoryName := categoryService.categoryRepo.VerifyCategoryName(updateRequest.CategoryName)
		if !verifyCategoryName {
			return errors.New("category name already exists")
		}
		category.CategoryName = updateRequest.CategoryName
		category.CategorySlug = slug.Make(updateRequest.CategoryName)
	}
	if updateRequest.CategoryStatus != "" {
		category.CategoryStatus = updateRequest.CategoryStatus
	}
	return categoryService.categoryRepo.Update(category.ID, *category)
}

func (categoryService categoryServiceImp) DeleteCategory(categoryID int) error {
	category, err := categoryService.categoryRepo.GetOne(categoryID)
	if err != nil {
		return errors.New("category not found")
	}
	return categoryService.categoryRepo.Delete(category.ID)

}
