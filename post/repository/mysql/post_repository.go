package mysql

import (
	"context"

  	"github.com/jinzhu/gorm"

	"domain"
)

type PostRepository struct {
	DB *gorm.DB
}

func Create(DB *gorm.DB) domain.PostRepository {
	return &PostRepository{DB}
}

func (repository *PostRepository) Get(c context.Context) (res []domain.Post) {
	repository.DB.Find(&res)

	return res
}

func (repository *PostRepository) Find(c context.Context, id uint) (res domain.Post) {
	repository.DB.First(&res, id)

	return res
}

func (repository *PostRepository) Store (c context.Context, model *domain.Post) (res domain.Post) {
	repository.DB.Create(&model)
	res = *model

	return res
}

func (repository *PostRepository) Update (c context.Context, model *domain.Post) (error error) {
	repository.DB.Model(&model).Update(&model)

	return nil
}

func (repository *PostRepository) Delete (c context.Context, id uint) (error error) {
	repository.DB.Delete(domain.Post{}, "id = ?", id)

	return nil
}