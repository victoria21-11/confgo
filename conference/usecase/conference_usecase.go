package usecase

import (
	"context"
	"time"

	"domain"
)

type ConferenceUsecase struct {
	repository domain.ConferenceRepository
	contextTimeout time.Duration
}

func Create(repository domain.ConferenceRepository, contextTimeout time.Duration) domain.ConferenceUsecase {
	return &ConferenceUsecase{
		repository,
		contextTimeout,
	}
}

func (usecase *ConferenceUsecase) Index(c context.Context) (res []domain.Conference) {
	res = usecase.repository.Get(c)

	return res
}

func (usecase *ConferenceUsecase) Show(c context.Context, id uint) (res domain.Conference) {
	res = usecase.repository.Find(c, id)

	return res
}

func (usecase *ConferenceUsecase) Store(c context.Context, model *domain.Conference) (res domain.Conference) {
	res = usecase.repository.Store(c, model)

	return res
}

func (usecase *ConferenceUsecase) Update(c context.Context, model *domain.Conference) (error error) {
	usecase.repository.Update(c, model)

	return nil
}

func (usecase *ConferenceUsecase) Delete(c context.Context, id uint) (error error) {
	usecase.repository.Delete(c, id)

	return nil
}