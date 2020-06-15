package usecase

import (
	"context"
	"time"

	"domain"
)

type EventUsecase struct {
	repository domain.EventRepository
	contextTimeout time.Duration
}

func Create(repository domain.EventRepository, contextTimeout time.Duration) domain.EventUsecase {
	return &EventUsecase{
		repository,
		contextTimeout,
	}
}

func (usecase *EventUsecase) Index(c context.Context) (res []domain.Event) {
	res = usecase.repository.Get(c)

	return res
}

func (usecase *EventUsecase) Show(c context.Context, id uint) (res domain.Event) {
	res = usecase.repository.Find(c, id)

	return res
}

func (usecase *EventUsecase) Store(c context.Context, model *domain.Event) (res domain.Event) {
	res = usecase.repository.Store(c, model)

	return res
}

func (usecase *EventUsecase) Update(c context.Context, model *domain.Event) (error error) {
	usecase.repository.Update(c, model)

	return nil
}

func (usecase *EventUsecase) Delete(c context.Context, id uint) (error error) {
	usecase.repository.Delete(c, id)

	return nil
}