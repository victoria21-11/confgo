package domain

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/go-playground/validator"
)

type Section struct {
	gorm.Model

	Conference Conference
	Coordinator Coordinator

	Title string
	Description string
	Published bool

	CustomValidator struct {
		validator *validator.Validate
	}
}

type SectionUsecase interface {
	Index(c context.Context) []Section
	Show(c context.Context, id uint) Section
	Store(c context.Context, model *Section) Section
	Update(c context.Context, model *Section) error
	Delete(c context.Context, id uint) error
}

type SectionRepository interface {
	Get(c context.Context) []Section
	Find(c context.Context, id uint) Section
	Store(c context.Context, model *Section) Section
	Update(c context.Context, model *Section) error
	Delete(c context.Context, id uint) error
}