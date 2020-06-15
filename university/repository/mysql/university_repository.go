package mysql

import (
	"context"

  	"github.com/jinzhu/gorm"

	"domain"
)

type UniversityRepository struct {
	DB *gorm.DB
}

func Create(DB *gorm.DB) domain.UniversityRepository {
	return &UniversityRepository{DB}
}

func (repository *UniversityRepository) Get(c context.Context) (res []domain.University) {
	repository.DB.Find(&res)

	return res
}

func (repository *UniversityRepository) Find(c context.Context, id uint) (res domain.University) {
	repository.DB.First(&res, id)

	return res
}

func (repository *UniversityRepository) Store (c context.Context, model *domain.University) (res domain.University) {
	repository.DB.Create(&model)
	res = *model

	return res
}

func (repository *UniversityRepository) Update (c context.Context, model *domain.University) (error error) {
	repository.DB.Model(&model).Update(&model)

	return nil
}

func (repository *UniversityRepository) Delete (c context.Context, id uint) (error error) {
	repository.DB.Delete(domain.University{}, "id = ?", id)

	return nil
}