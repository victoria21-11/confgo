package usecase

import (
	"context"
	"time"

	"domain"
)

type CoordinatorUsecase struct {
	repository domain.CoordinatorRepository
	contextTimeout time.Duration
}

func Create(repository domain.CoordinatorRepository, contextTimeout time.Duration) domain.CoordinatorUsecase {
	return &CoordinatorUsecase{
		repository,
		contextTimeout,
	}
}

func (usecase *CoordinatorUsecase) Index(c context.Context) (res []domain.Coordinator) {
	res = usecase.repository.Get(c)

	return res
}

func (usecase *CoordinatorUsecase) Show(c context.Context, id uint) (res domain.Coordinator) {
	res = usecase.repository.Find(c, id)

	return res
}

func (usecase *CoordinatorUsecase) Store(c context.Context, model *domain.Coordinator) (res domain.Coordinator) {
	res = usecase.repository.Store(c, model)

	return res
}

func (usecase *CoordinatorUsecase) Update(c context.Context, model *domain.Coordinator) (error error) {
	usecase.repository.Update(c, model)

	return nil
}

func (usecase *CoordinatorUsecase) Delete(c context.Context, id uint) (error error) {
	usecase.repository.Delete(c, id)

	return nil
}