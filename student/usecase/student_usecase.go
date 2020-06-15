package usecase

import (
	"context"
	"time"

	"domain"
)

type StudentUsecase struct {
	repository domain.StudentRepository
	contextTimeout time.Duration
}

func Create(repository domain.StudentRepository, contextTimeout time.Duration) domain.StudentUsecase {
	return &StudentUsecase{
		repository,
		contextTimeout,
	}
}

func (usecase *StudentUsecase) Index(c context.Context) (res []domain.Student) {
	res = usecase.repository.Get(c)

	return res
}

func (usecase *StudentUsecase) Show(c context.Context, id uint) (res domain.Student) {
	res = usecase.repository.Find(c, id)

	return res
}

func (usecase *StudentUsecase) Store(c context.Context, model *domain.Student) (res domain.Student) {
	res = usecase.repository.Store(c, model)

	return res
}

func (usecase *StudentUsecase) Update(c context.Context, model *domain.Student) (error error) {
	usecase.repository.Update(c, model)

	return nil
}

func (usecase *StudentUsecase) Delete(c context.Context, id uint) (error error) {
	usecase.repository.Delete(c, id)

	return nil
}