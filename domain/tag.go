package domain

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/go-playground/validator"
)

type Tag struct {
	gorm.Model

	Title string
	Verified bool

	CustomValidator struct {
		validator *validator.Validate
	}
}

type TagUsecase interface {
	Index(c context.Context) []Tag
	Show(c context.Context, id uint) Tag
	Store(c context.Context, model *Tag) Tag
	Update(c context.Context, model *Tag) error
	Delete(c context.Context, id uint) error
}

type TagRepository interface {
	Get(c context.Context) []Tag
	Find(c context.Context, id uint) Tag
	Store(c context.Context, model *Tag) Tag
	Update(c context.Context, model *Tag) error
	Delete(c context.Context, id uint) error
}