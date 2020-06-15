package domain

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/go-playground/validator"
)

type Coordinator struct {
	gorm.Model

	User User

	CustomValidator struct {
		validator *validator.Validate
	}
}

type CoordinatorUsecase interface {
	Index(c context.Context) []Coordinator
	Show(c context.Context, id uint) Coordinator
	Store(c context.Context, model *Coordinator) Coordinator
	Update(c context.Context, model *Coordinator) error
	Delete(c context.Context, id uint) error
}

type CoordinatorRepository interface {
	Get(c context.Context) []Coordinator
	Find(c context.Context, id uint) Coordinator
	Store(c context.Context, model *Coordinator) Coordinator
	Update(c context.Context, model *Coordinator) error
	Delete(c context.Context, id uint) error
}