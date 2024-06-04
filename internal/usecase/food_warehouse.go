package usecase

import (
	"context"
	"crm-farmish/internal/entity"
	"github.com/google/uuid"
)

type FoodWarehouseUseCase struct {
	Repo FoodWarehouse
}

func NewFoodWarehouseUseCase(r FoodWarehouse) *FoodWarehouseUseCase {
	return &FoodWarehouseUseCase{Repo: r}
}

func (r *FoodWarehouseUseCase) CreateFoodWarehouse(ctx context.Context, req *entity.FoodWarehouseCreate) (*entity.FoodWarehouse, error) {
	req.ID = uuid.New().String()
	return r.Repo.CreateFoodWarehouse(ctx, req)
}

func (r *FoodWarehouseUseCase) GetFoodWarehouse(ctx context.Context, req *entity.FieldValueReq) (*entity.FoodWarehouse, error) {
	return r.Repo.GetFoodWarehouse(ctx, req)
}

func (r *FoodWarehouseUseCase) ListFoodWarehouse(ctx context.Context, req *entity.ListReq) (*entity.ListFoodWarehouse, error) {
	return r.Repo.ListFoodWarehouse(ctx, req)

}

func (r *FoodWarehouseUseCase) UpdateFoodWarehouse(ctx context.Context, req *entity.UpdateFoodWarehouseReq) (*entity.FoodWarehouse, error) {
	return r.Repo.UpdateFoodWarehouse(ctx, req)

}

func (r *FoodWarehouseUseCase) DeleteFoodWarehouse(ctx context.Context, req *entity.FieldValueReq) (*entity.StatusRes, error) {
	return r.Repo.DeleteFoodWarehouse(ctx, req)

}
