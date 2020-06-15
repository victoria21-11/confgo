package mysql

import (
	"context"

  	"github.com/jinzhu/gorm"

	"domain"
)

type SectionRepository struct {
	DB *gorm.DB
}

func Create(DB *gorm.DB) domain.SectionRepository {
	return &SectionRepository{DB}
}

func (repository *SectionRepository) Get(c context.Context) (res []domain.Section) {
	repository.DB.Find(&res)

	return res
}

func (repository *SectionRepository) Find(c context.Context, id uint) (res domain.Section) {
	repository.DB.First(&res, id)

	return res
}

func (repository *SectionRepository) Store (c context.Context, model *domain.Section) (res domain.Section) {
	repository.DB.Create(&model)
	res = *model

	return res
}

func (repository *SectionRepository) Update (c context.Context, model *domain.Section) (error error) {
	repository.DB.Model(&model).Update(&model)

	return nil
}

func (repository *SectionRepository) Delete (c context.Context, id uint) (error error) {
	repository.DB.Delete(domain.Section{}, "id = ?", id)

	return nil
}