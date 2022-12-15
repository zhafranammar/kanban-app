package repository

import (
	"a21hc3NpZ25tZW50/entity"
	"context"
	"errors"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error)
	StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error)
	StoreManyCategory(ctx context.Context, categories []entity.Category) error
	GetCategoryByID(ctx context.Context, id int) (entity.Category, error)
	UpdateCategory(ctx context.Context, category *entity.Category) error
	DeleteCategory(ctx context.Context, id int) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetCategoriesByUserId(ctx context.Context, id int) ([]entity.Category, error) {
	categories := []entity.Category{}
	err := r.db.Where("user_id = ?",id).Find(&categories).Error
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			return categories,nil
		}
		return categories, err
	}

	return categories, nil // TODO: replace this
}

func (r *categoryRepository) StoreCategory(ctx context.Context, category *entity.Category) (categoryId int, err error) {
	res := r.db.Create(&category).Error
	if res != nil {
		return 0,res
	}
	return category.ID, nil // TODO: replace this
}

func (r *categoryRepository) StoreManyCategory(ctx context.Context, categories []entity.Category) error {
	err := r.db.CreateInBatches(categories, 5).Error
	return err // TODO: replace this
}

func (r *categoryRepository) GetCategoryByID(ctx context.Context, id int) (entity.Category, error) {
	categories := entity.Category{}
	err := r.db.Where("id = ?",id).First(categories).Error
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			return categories,nil
		}
		return categories, err
	}
	return categories, nil // TODO: replace this
}

func (r *categoryRepository) UpdateCategory(ctx context.Context, category *entity.Category) error {
	err :=  r.db.Where("id = ?", category.ID).Updates(&category).Error
	return err // TODO: replace this
}

func (r *categoryRepository) DeleteCategory(ctx context.Context, id int) error {
	categories := entity.Category{}
	err :=  r.db.Where("id = ?", id).Delete(&categories).Error
	return err // TODO: replace this
}
