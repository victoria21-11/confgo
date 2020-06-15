package usecase

import (
	"context"
	"time"

	"domain"
)

type TagUsecase struct {
	repository domain.TagRepository
	contextTimeout time.Duration
}

func Create(repository domain.TagRepository, contextTimeout time.Duration) domain.TagUsecase {
	return &TagUsecase{
		repository,
		contextTimeout,
	}
}

func (usecase *TagUsecase) Index(c context.Context) (res []domain.Tag) {
	res = usecase.repository.Get(c)

	return res
}

func (usecase *TagUsecase) Show(c context.Context, id uint) (res domain.Tag) {
	res = usecase.repository.Find(c, id)

	return res
}

func (usecase *TagUsecase) Store(c context.Context, model *domain.Tag) (res domain.Tag) {
	res = usecase.repository.Store(c, model)

	return res
}

func (usecase *TagUsecase) Update(c context.Context, model *domain.Tag) (error error) {
	usecase.repository.Update(c, model)

	return nil
}

func (usecase *TagUsecase) Delete(c context.Context, id uint) (error error) {
	usecase.repository.Delete(c, id)

	return nil
}