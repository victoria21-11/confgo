package mysql

import (
	"context"

  	"github.com/jinzhu/gorm"

	"domain"
)

type RoleRepository struct {
	DB *gorm.DB
}

func Create(DB *gorm.DB) domain.RoleRepository {
	return &RoleRepository{DB}
}

func (repository *RoleRepository) Get(c context.Context) (res []domain.Role) {
	repository.DB.Find(&res)

	return res
}

func (repository *RoleRepository) Find(c context.Context, id uint) (res domain.Role) {
	repository.DB.First(&res, id)

	return res
}

func (repository *RoleRepository) Store (c context.Context, model *domain.Role) (res domain.Role) {
	repository.DB.Create(&model)
	res = *model

	return res
}

func (repository *RoleRepository) Update (c context.Context, model *domain.Role) (err error) {
	repository.DB.Model(&model).Updates(&model)
	
	return nil
}

func (repository *RoleRepository) Delete (c context.Context, id uint) (err error) {
	repository.DB.Delete(domain.Role{}, "id = ?", id)

	return nil
}