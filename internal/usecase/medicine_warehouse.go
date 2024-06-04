package usecase

import (
	"context"
	"crm-farmish/internal/entity"
	"github.com/google/uuid"
)

type MedicineWarehouseUseCase struct {
	Repo MedicineWarehouse
}

func NewMedicineWarehouseUseCase(r MedicineWarehouse) *MedicineWarehouseUseCase {
	return &MedicineWarehouseUseCase{Repo: r}
}

func (r *MedicineWarehouseUseCase) CreateMedicine(ctx context.Context, req *entity.MedicineWarehouseCreate) (*entity.MedicineWarehouse, error) {
	req.ID = uuid.New().String()

	return r.Repo.CreateMedicine(ctx, req)
}

func (r *MedicineWarehouseUseCase) GetMedicine(ctx context.Context, req *entity.FieldValueReq) (*entity.MedicineWarehouse, error) {
	return r.Repo.GetMedicine(ctx, req)
}

func (r *MedicineWarehouseUseCase) ListMedicine(ctx context.Context, req *entity.ListReq) (*entity.ListMedicineWarehouse, error) {
	return r.Repo.ListMedicine(ctx, req)

}

func (r *MedicineWarehouseUseCase) UpdateMedicineType(ctx context.Context, req *entity.UpdateMedicineWarehouseReq) (*entity.MedicineWarehouse, error) {
	return r.Repo.UpdateMedicineType(ctx, req)

}

func (r *MedicineWarehouseUseCase) DeleteMedicine(ctx context.Context, req *entity.FieldValueReq) (*entity.StatusRes, error) {
	return r.Repo.DeleteMedicine(ctx, req)

}
