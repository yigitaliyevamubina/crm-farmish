package usecase

import (
	"context"
	"crm-farmish/internal/entity"
	"github.com/google/uuid"
)

type AnimalTypeUseCase struct {
	Repo AnimalType
}

func NewAnimalTypeUseCase(r AnimalType) *AnimalTypeUseCase {
	return &AnimalTypeUseCase{Repo: r}
}

func (r *AnimalTypeUseCase) CreateAnimalTypes(ctx context.Context, req *entity.AnimalTypeCreate) (*entity.AnimalType, error) {
	req.ID = uuid.New().String()

	return r.Repo.CreateAnimalTypes(ctx, req)
}

func (r *AnimalTypeUseCase) GetAnimalTypes(ctx context.Context, req *entity.FieldValueReq) (*entity.AnimalType, error) {
	return r.Repo.GetAnimalTypes(ctx, req)
}

func (r *AnimalTypeUseCase) ListAnimalTypes(ctx context.Context, req *entity.ListReq) (*entity.ListAnimalTypeRes, error) {
	return r.Repo.ListAnimalTypes(ctx, req)

}

func (r *AnimalTypeUseCase) UpdateAnimalsType(ctx context.Context, req *entity.UpdateAnimalTypeReq) (*entity.AnimalType, error) {
	return r.Repo.UpdateAnimalsType(ctx, req)

}

func (r *AnimalTypeUseCase) DeleteAnimalTypes(ctx context.Context, req *entity.FieldValueReq) (*entity.StatusRes, error) {
	return r.Repo.DeleteAnimalTypes(ctx, req)
}
