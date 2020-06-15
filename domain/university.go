package domain

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/go-playground/validator"
)

type University struct {
	gorm.Model

	Title string
	Published bool

	CustomValidator struct {
		validator *validator.Validate
	}
}

type UniversityUsecase interface {
	Index(c context.Context) []University
	Show(c context.Context, id uint) University
	Store(c context.Context, model *University) University
	Update(c context.Context, model *University) error
	Delete(c context.Context, id uint) error
}

type UniversityRepository interface {
	Get(c context.Context) []University
	Find(c context.Context, id uint) University
	Store(c context.Context, model *University) University
	Update(c context.Context, model *University) error
	Delete(c context.Context, id uint) error
}