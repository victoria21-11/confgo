package usecase

import (
	"context"
	"time"

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

func (usecase *SectionUsecase) Store(c context.Context, model *domain.Section) (res domain.Section) {
	res = usecase.repository.Store(c, model)

	return res
}

func (usecase *SectionUsecase) Update(c context.Context, model *domain.Section) (error error) {
	usecase.repository.Update(c, model)

	return nil
}

func (usecase *SectionUsecase) Delete(c context.Context, id uint) (error error) {
	usecase.repository.Delete(c, id)

	return nil
}