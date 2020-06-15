package usecase

import (
	"context"
	"time"

	"domain"
)

type UniversityUsecase struct {
	repository domain.UniversityRepository
	contextTimeout time.Duration
}

func Create(repository domain.UniversityRepository, contextTimeout time.Duration) domain.UniversityUsecase {
	return &UniversityUsecase{
		repository,
		contextTimeout,
	}
}

func (usecase *UniversityUsecase) Index(c context.Context) (res []domain.University) {
	res = usecase.repository.Get(c)

	return res
}

func (usecase *UniversityUsecase) Show(c context.Context, id uint) (res domain.University) {
	res = usecase.repository.Find(c, id)

	return res
}

func (usecase *UniversityUsecase) Store(c context.Context, model *domain.University) (res domain.University) {
	res = usecase.repository.Store(c, model)

	return res
}

func (usecase *UniversityUsecase) Update(c context.Context, model *domain.University) (error error) {
	usecase.repository.Update(c, model)

	return nil
}

func (usecase *UniversityUsecase) Delete(c context.Context, id uint) (error error) {
	usecase.repository.Delete(c, id)

	return nil
}