package domain

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/go-playground/validator"
)

type File struct {
	gorm.Model

	User User

	Path string
	MimeType string
	Name string
	OriginalName string
	PublicName string

	CustomValidator struct {
		validator *validator.Validate
	}
}

type FileUsecase interface {
	Index(c context.Context) []File
	Show(c context.Context, id uint) File
	Store(c context.Context, model *File) File
	Update(c context.Context, model *File) error
	Delete(c context.Context, id uint) error
}

type FileRepository interface {
	Get(c context.Context) []File
	Find(c context.Context, id uint) File
	Store(c context.Context, model *File) File
	Update(c context.Context, model *File) error
	Delete(c context.Context, id uint) error
}