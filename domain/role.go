package domain

import(
	"context"

	"github.com/jinzhu/gorm"
	"github.com/go-playground/validator"
)

type Role struct {
	gorm.Model

	Title string `json:"title" validate:"required"`
	Name string `json:"name" validate:"required"`

	CustomValidator struct {
		validator *validator.Validate
	}
}

type RoleUsecase interface {
	Index(c context.Context) []Role
	Show(c context.Context, id uint) Role
	Store(c context.Context, model *Role) Role
	Update(c context.Context, model *Role) error
	Delete(c context.Context, id uint) error
}

type RoleRepository interface {
	Get(c context.Context) []Role
	Find(c context.Context, id uint) Role
	Store(c context.Context, model *Role) Role
	Update(c context.Context, model *Role) error
	Delete(c context.Context, id uint) error
}