package usecase

import (
	"context"
	"time"

	"domain"
)

type RoleUsecase struct {
	repository domain.RoleRepository
	contextTimeout time.Duration
}

func Create(repository domain.RoleRepository, contextTimeout time.Duration) domain.RoleUsecase {
	return &RoleUsecase{
		repository,
		contextTimeout,
	}
}

func (usecase *RoleUsecase) Index(c context.Context) (res []domain.Role) {
	res = usecase.repository.Get(c)

	return res
}

func (usecase *RoleUsecase) Show(c context.Context, id uint) (res domain.Role) {
	res = usecase.repository.Find(c, id)

	return res
}

func (usecase *RoleUsecase) Store(c context.Context, model *domain.Role) (res domain.Role) {
	res = usecase.repository.Store(c, model)

	return res
}

func (usecase *RoleUsecase) Update(c context.Context, model *domain.Role) (error error) {
	usecase.repository.Update(c, model)

	return nil
}

func (usecase *RoleUsecase) Delete(c context.Context, id uint) (error error) {
	usecase.repository.Delete(c, id)

	return nil
}