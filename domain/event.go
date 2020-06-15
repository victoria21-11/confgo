package domain

import (
	"time"
	"context"

	"github.com/jinzhu/gorm"
	"github.com/go-playground/validator"
)

type Event struct {
	gorm.Model

	Title string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Published bool `json:"published" validate:"required"`
	StartDate time.Time `json:"start_date" validate:"datetime"`
	EndDate time.Time `json:"end_date" validate:"datetime"`

	CustomValidator struct {
		validator *validator.Validate
	}
}

type EventUsecase interface {
	Index(c context.Context) []Event
	Show(c context.Context, id uint) Event
	Store(c context.Context, model *Event) Event
	Update(c context.Context, model *Event) error
	Delete(c context.Context, id uint) error
}

type EventRepository interface {
	Get(c context.Context) []Event
	Find(c context.Context, id uint) Event
	Store(c context.Context, model *Event) Event
	Update(c context.Context, model *Event) error
	Delete(c context.Context, id uint) error
}