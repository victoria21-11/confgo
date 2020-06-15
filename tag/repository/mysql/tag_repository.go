package mysql

import (
	"context"

  	"github.com/jinzhu/gorm"

	"domain"
)

type TagRepository struct {
	DB *gorm.DB
}

func Create(DB *gorm.DB) domain.TagRepository {
	return &TagRepository{DB}
}

func (repository *TagRepository) Get(c context.Context) (res []domain.Tag) {
	repository.DB.Find(&res)

	return res
}

func (repository *TagRepository) Find(c context.Context, id uint) (res domain.Tag) {
	repository.DB.First(&res, id)

	return res
}

func (repository *TagRepository) Store (c context.Context, model *domain.Tag) (res domain.Tag) {
	repository.DB.Create(&model)
	res = *model

	return res
}

func (repository *TagRepository) Update (c context.Context, model *domain.Tag) (error error) {
	repository.DB.Model(&model).Update(&model)

	return nil
}

func (repository *TagRepository) Delete (c context.Context, id uint) (error error) {
	repository.DB.Delete(domain.Tag{}, "id = ?", id)

	return nil
}