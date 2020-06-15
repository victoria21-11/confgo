package domain

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/go-playground/validator"
)

type Participation struct {
	gorm.Model

	Conference int
	Section int

	Published bool
	Verified bool
	Precence bool
	ScientificDirector string

	CustomValidator struct {
		validator *validator.Validate
	}
}

type ParticipationUsecase interface {
	Index(c context.Context) []Participation
	Show(c context.Context, id uint) Participation
	Store(c context.Context, model *Participation) Participation
	Update(c context.Context, model *Participation) error
	Delete(c context.Context, id uint) error
}

type ParticipationRepository interface {
	Get(c context.Context) []Participation
	Find(c context.Context, id uint) Participation
	Store(c context.Context, model *Participation) Participation
	Update(c context.Context, model *Participation) error
	Delete(c context.Context, id uint) error
}