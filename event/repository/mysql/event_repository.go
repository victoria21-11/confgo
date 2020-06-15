package mysql

import (
	"context"

  	"github.com/jinzhu/gorm"

	"domain"
)

type EventRepository struct {
	DB *gorm.DB
}

func Create(DB *gorm.DB) domain.EventRepository {
	return &EventRepository{DB}
}

func (repository *EventRepository) Get(c context.Context) (res []domain.Event) {
	repository.DB.Find(&res)

	return res
}

func (repository *EventRepository) Find(c context.Context, id uint) (res domain.Event) {
	repository.DB.First(&res, id)

	return res
}

func (repository *EventRepository) Store (c context.Context, model *domain.Event) (res domain.Event) {
	repository.DB.Create(&model)
	res = *model

	return res
}

func (repository *EventRepository) Update (c context.Context, model *domain.Event) (error error) {
	repository.DB.Model(&model).Update(&model)

	return nil
}

func (repository *EventRepository) Delete (c context.Context, id uint) (error error) {
	repository.DB.Delete(domain.Event{}, "id = ?", id)

	return nil
}