package mysql

import (
	"context"

  	"github.com/jinzhu/gorm"

	"domain"
)

type CoordinatorRepository struct {
	DB *gorm.DB
}

func Create(DB *gorm.DB) domain.CoordinatorRepository {
	return &CoordinatorRepository{DB}
}

func (repository *CoordinatorRepository) Get(c context.Context) (res []domain.Coordinator) {
	repository.DB.Find(&res)

	return res
}

func (repository *CoordinatorRepository) Find(c context.Context, id uint) (res domain.Coordinator) {
	repository.DB.First(&res, id)

	return res
}

func (repository *CoordinatorRepository) Store (c context.Context, model *domain.Coordinator) (res domain.Coordinator) {
	repository.DB.Create(&model)
	res = *model

	return res
}

func (repository *CoordinatorRepository) Update (c context.Context, model *domain.Coordinator) (error error) {
	repository.DB.Model(&model).Update(&model)

	return nil
}

func (repository *CoordinatorRepository) Delete (c context.Context, id uint) (error error) {
	repository.DB.Delete(domain.Coordinator{}, "id = ?", id)

	return nil
}