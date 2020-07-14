package usecase

import (
	"context"
	"time"
	"errors"

	"domain"
)

type SectionUsecase struct {
	repository domain.SectionRepository
	contextTimeout time.Duration
}

func Create(repository domain.SectionRepository, contextTimeout time.Duration) domain.SectionUsecase {
	return &SectionUsecase{
		repository,
		contextTimeout,
	}
}

func (usecase *SectionUsecase) Index(c context.Context) (res []domain.Section) {
	res = usecase.repository.Get(c)

	return res
}

func (usecase *SectionUsecase) Show(c context.Context, id uint) (res domain.Section) {
	res = usecase.repository.Find(c, id)

	return res
}

func (usecase *SectionUsecase) Store(c context.Context, model *domain.Section) (res domain.Section, err error) {
	usecase.repository.FindCoordinator(c, model)
	usecase.repository.FindConference(c, model)

	if model.Coordinator.ID == 0 {
		err = errors.New("Coordinator not found");
	}

	if model.Conference.ID == 0 {
		err = errors.New("Conference not found");
	}

	if err == nil {
		res = usecase.repository.Store(c, model)
	}

	return res, err
}

func (usecase *SectionUsecase) Update(c context.Context, model *domain.Section) (err error) {
	usecase.repository.FindCoordinator(c, model)
	usecase.repository.FindConference(c, model)

	if model.Coordinator.ID == 0 {
		err = errors.New("Coordinator not found");
	}

	if model.Conference.ID == 0 {
		err = errors.New("Conference not found");
	}

	if err == nil {
		usecase.repository.Update(c, model)
	}

	return err
}

func (usecase *SectionUsecase) Delete(c context.Context, id uint) (error error) {
	usecase.repository.Delete(c, id)

	return nil
}