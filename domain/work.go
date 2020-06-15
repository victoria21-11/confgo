package domain

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/go-playground/validator"
)

type Work struct {
	gorm.Model

	User User
	File File

	Title string
	Description string
	Published bool

	CustomValidator struct {
		validator *validator.Validate
	}
}

type WorkUsecase interface {
	Index(c context.Context) []Work
	Show(c context.Context, id uint) Work
	Store(c context.Context, model *Work) Work
	Update(c context.Context, model *Work) error
	Delete(c context.Context, id uint) error
}

type WorkRepository interface {
	Get(c context.Context) []Work
	Find(c context.Context, id uint) Work
	Store(c context.Context, model *Work) Work
	Update(c context.Context, model *Work) error
	Delete(c context.Context, id uint) error
}