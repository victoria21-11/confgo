package domain

import (
	"time"
	"context"

	"github.com/jinzhu/gorm"
	"github.com/go-playground/validator"
)

type Conference struct {
	gorm.Model

	Title string
	Description string
	StartDate time.Time
	EndDate time.Time
	Active bool

	CustomValidator struct {
		validator *validator.Validate
	}
}

type ConferenceUsecase interface {
	Index(c context.Context) []Conference
	Show(c context.Context, id uint) Conference
	Store(c context.Context, model *Conference) Conference
	Update(c context.Context, model *Conference) error
	Delete(c context.Context, id uint) error
}

type ConferenceRepository interface {
	Get(c context.Context) []Conference
	Find(c context.Context, id uint) Conference
	Store(c context.Context, model *Conference) Conference
	Update(c context.Context, model *Conference) error
	Delete(c context.Context, id uint) error
}