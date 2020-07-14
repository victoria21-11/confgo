package domain

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/go-playground/validator"
)

type Section struct {
	gorm.Model

	Conference Conference `json:"conference" validate:"required" gorm:"association_autoupdate:false;association_autocreate:false"`
	Coordinator User `json:"coordinator" validate:"required" gorm:"association_autoupdate:false;association_autocreate:false"`

	Title string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Active bool `json:"active" gorm:"default:false"`
	ConferenceID uint `json:"conference_id" validate:"required"`
	CoordinatorID uint `json:"coordinator_id" validate:"required"`

	CustomValidator struct {
		validator *validator.Validate
	}
}

type SectionUsecase interface {
	Index(c context.Context) []Section
	Show(c context.Context, id uint) Section
	Store(c context.Context, model *Section) (Section, error)
	Update(c context.Context, model *Section) error
	Delete(c context.Context, id uint) error
}

type SectionRepository interface {
	Get(c context.Context) []Section
	Find(c context.Context, id uint) Section
	Store(c context.Context, model *Section) Section
	Update(c context.Context, model *Section) error
	Delete(c context.Context, id uint) error
	FindCoordinator(c context.Context, model *Section) Section
	FindConference(c context.Context, model *Section) Section
}