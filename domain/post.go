package domain

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/go-playground/validator"
)

type Post struct {
	gorm.Model

	Title string
	Description string
	Published bool

	CustomValidator struct {
		validator *validator.Validate
	}
}

type PostUsecase interface {
	Index(c context.Context) []Post
	Show(c context.Context, id uint) Post
	Store(c context.Context, model *Post) Post
	Update(c context.Context, model *Post) error
	Delete(c context.Context, id uint) error
}

type PostRepository interface {
	Get(c context.Context) []Post
	Find(c context.Context, id uint) Post
	Store(c context.Context, model *Post) Post
	Update(c context.Context, model *Post) error
	Delete(c context.Context, id uint) error
}