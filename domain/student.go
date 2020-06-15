package domain

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/go-playground/validator"
)

type Student struct {
	gorm.Model

	User User
	University University

	Institute string
	Faculty string
	Department string

	CustomValidator struct {
		validator *validator.Validate
	}
}

type StudentUsecase interface {
	Index(c context.Context) []Student
	Show(c context.Context, id uint) Student
	Store(c context.Context, model *Student) Student
	Update(c context.Context, model *Student) error
	Delete(c context.Context, id uint) error
}

type StudentRepository interface {
	Get(c context.Context) []Student
	Find(c context.Context, id uint) Student
	Store(c context.Context, model *Student) Student
	Update(c context.Context, model *Student) error
	Delete(c context.Context, id uint) error
}