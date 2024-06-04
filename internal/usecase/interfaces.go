package usecase

import (
	"context"
	"crm-farmish/internal/entity"
)

// Interfaces Repo -.
type (
	// AnimalType -.
	AnimalType interface {
		CreateAnimalTypes(ctx context.Context, req *entity.AnimalTypeCreate) (*entity.AnimalType, error)
		GetAnimalTypes(ctx context.Context, req *entity.FieldValueReq) (*entity.AnimalType, error)
		ListAnimalTypes(ctx context.Context, req *entity.ListReq) (*entity.ListAnimalTypeRes, error)
		UpdateAnimalsType(ctx context.Context, req *entity.UpdateAnimalTypeReq) (*entity.AnimalType, error)
		DeleteAnimalTypes(ctx context.Context, req *entity.FieldValueReq) (*entity.StatusRes, error)
	}

	// Animals -.
	Animals interface {
		CreateAnimals(ctx context.Context, req *entity.AnimalsCreate) (*entity.Animals, error)
		GetAnimals(ctx context.Context, req *entity.FieldValueReq) (*entity.Animals, error)
		ListAnimals(ctx context.Context, req *entity.ListReq) (*entity.ListAnimalRes, error)
		UpdateAnimalsType(ctx context.Context, req *entity.UpdateAnimalTypeReq) (*entity.Animals, error)
		DeleteAnimals(ctx context.Context, req *entity.FieldValueReq) (*entity.StatusRes, error)
	}

	// FoodWarehouse -.
	FoodWarehouse interface {
		CreateFoodWarehouse(ctx context.Context, req *entity.FoodWarehouseCreate) (*entity.FoodWarehouse, error)
		GetFoodWarehouse(ctx context.Context, req *entity.FieldValueReq) (*entity.FoodWarehouse, error)
		ListFoodWarehouse(ctx context.Context, req *entity.ListReq) (*entity.ListFoodWarehouse, error)
		UpdateFoodWarehouseType(ctx context.Context, req *entity.UpdateFoodWarehouseReq) (*entity.FoodWarehouse, error)
		DeleteFoodWarehouse(ctx context.Context, req *entity.FieldValueReq) (*entity.StatusRes, error)
	}

	// MedicineWarehouse -.
	MedicineWarehouse interface {
		CreateMedicine(ctx context.Context, req *entity.MedicineWarehouseCreate) (*entity.MedicineWarehouse, error)
		GetMedicine(ctx context.Context, req *entity.FieldValueReq) (*entity.MedicineWarehouse, error)
		ListMedicine(ctx context.Context, req *entity.ListReq) (*entity.ListMedicineWarehouse, error)
		UpdateMedicineType(ctx context.Context, req *entity.UpdateMedicineWarehouseReq) (*entity.MedicineWarehouse, error)
		DeleteMedicine(ctx context.Context, req *entity.FieldValueReq) (*entity.StatusRes, error)
	}
)