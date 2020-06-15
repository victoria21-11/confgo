package mysql

import (
	"context"

  	"github.com/jinzhu/gorm"

	"domain"
)

type StudentRepository struct {
	DB *gorm.DB
}

func Create(DB *gorm.DB) domain.StudentRepository {
	return &StudentRepository{DB}
}

func (repository *StudentRepository) Get(c context.Context) (res []domain.Student) {
	repository.DB.Find(&res)

	return res
}

func (repository *StudentRepository) Find(c context.Context, id uint) (res domain.Student) {
	repository.DB.First(&res, id)

	return res
}

func (repository *StudentRepository) Store (c context.Context, model *domain.Student) (res domain.Student) {
	repository.DB.Create(&model)
	res = *model

	return res
}

func (repository *StudentRepository) Update (c context.Context, model *domain.Student) (error error) {
	repository.DB.Model(&model).Update(&model)

	return nil
}

func (repository *StudentRepository) Delete (c context.Context, id uint) (error error) {
	repository.DB.Delete(domain.Student{}, "id = ?", id)

	return nil
}