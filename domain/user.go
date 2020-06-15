package domain

import (
	"time"
	"context"

	"github.com/jinzhu/gorm"
	"github.com/go-playground/validator"
)

type User struct {
	gorm.Model

	Login string
	Password string
	Email string
	Phone string
	FirstName string
	LastName string
	BirthDate time.Time
	VerifiedAt time.Time

	CustomValidator struct {
		validator *validator.Validate
	}
}

type UserUsecase interface {
	Index(c context.Context) []User
	Show(c context.Context, id uint) User
	Store(c context.Context, model *User) User
	Update(c context.Context, model *User) error 
	Delete(c context.Context, id uint) error 
}

type UserRepository interface {
	Get(c context.Context) []User
	Find(c context.Context, id uint) User
	Store(c context.Context, model *User) User
	Update(c context.Context, model *User) error
	Delete(c context.Context, id uint) error 
}