package mysql

import (
	"context"

  	"github.com/jinzhu/gorm"

	"domain"
)

type ConferenceRepository struct {
	DB *gorm.DB
}

func Create(DB *gorm.DB) domain.ConferenceRepository {
	return &ConferenceRepository{DB}
}

func (repository *ConferenceRepository) Get(c context.Context) (res []domain.Conference) {
	repository.DB.Find(&res)

	return res
}

func (repository *ConferenceRepository) Find(c context.Context, id uint) (res domain.Conference) {
	repository.DB.First(&res, id)

	return res
}

func (repository *ConferenceRepository) Store (c context.Context, model *domain.Conference) (res domain.Conference) {
	repository.DB.Create(&model)
	res = *model

	return res
}

func (repository *ConferenceRepository) Update (c context.Context, model *domain.Conference) (error error) {
	repository.DB.Model(&model).Update(&model)

	return nil
}

func (repository *ConferenceRepository) Delete (c context.Context, id uint) (error error) {
	repository.DB.Delete(domain.Conference{}, "id = ?", id)

	return nil
}