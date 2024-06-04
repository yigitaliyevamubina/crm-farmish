package usecase

import (
	"context"
	"crm-farmish/internal/entity"
	"github.com/google/uuid"
)

type AnimalsUseCase struct {
	Repo Animals
}

func NewAnimalsUseCase(r Animals) *AnimalsUseCase {
	return &AnimalsUseCase{Repo: r}
}

func (r *AnimalsUseCase) CreateAnimals(ctx context.Context, req *entity.AnimalsCreate) (*entity.Animals, error) {
	req.ID = uuid.New().String()

	return r.Repo.CreateAnimals(ctx, req)
}

func (r *AnimalsUseCase) GetAnimals(ctx context.Context, req *entity.FieldValueReq) (*entity.Animals, error) {
	return r.Repo.GetAnimals(ctx, req)
}

func (r *AnimalsUseCase) ListAnimals(ctx context.Context, req *entity.ListReq) (*entity.ListAnimalRes, error) {
	return r.Repo.ListAnimals(ctx, req)
}

func (r *AnimalsUseCase) UpdateAnimalsType(ctx context.Context, req *entity.UpdateAnimalTypeReq) (*entity.Animals, error) {
	return r.Repo.UpdateAnimalsType(ctx, req)
}

func (r *AnimalsUseCase) DeleteAnimals(ctx context.Context, req *entity.FieldValueReq) (*entity.StatusRes, error) {
	return r.Repo.DeleteAnimals(ctx, req)
}
