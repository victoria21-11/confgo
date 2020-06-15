package usecase

import (
	"context"
	"time"

	"domain"
)

type PostUsecase struct {
	repository domain.PostRepository
	contextTimeout time.Duration
}

func Create(repository domain.PostRepository, contextTimeout time.Duration) domain.PostUsecase {
	return &PostUsecase{
		repository,
		contextTimeout,
	}
}

func (usecase *PostUsecase) Index(c context.Context) (res []domain.Post) {
	res = usecase.repository.Get(c)

	return res
}

func (usecase *PostUsecase) Show(c context.Context, id uint) (res domain.Post) {
	res = usecase.repository.Find(c, id)

	return res
}

func (usecase *PostUsecase) Store(c context.Context, model *domain.Post) (res domain.Post) {
	res = usecase.repository.Store(c, model)

	return res
}

func (usecase *PostUsecase) Update(c context.Context, model *domain.Post) (error error) {
	usecase.repository.Update(c, model)

	return nil
}

func (usecase *PostUsecase) Delete(c context.Context, id uint) (error error) {
	usecase.repository.Delete(c, id)

	return nil
}