package repo

import (
	"crm-farmish/internal/infrastructure/repository"
	"crm-farmish/internal/pkg/postgres"
)

type IFarmishStorage interface {
	AnimalType() repository.AnimalType
	Animals() repository.Animals
	FoodWarehouse() repository.FoodWarehouse
	MedicineWarehouse() repository.MedicineWarehouse
}

type FarmishStoragePg struct {
	animalType        repository.AnimalType
	animals           repository.Animals
	foodWarehouse     repository.FoodWarehouse
	medicineWarehouse repository.MedicineWarehouse
}

func NewStorage(db *postgres.PostgresDB) IFarmishStorage {
	return &FarmishStoragePg{
		animalType:        NewAnimalTypesRepo(db),
		animals:           NewAnimalsRepo(db),
		foodWarehouse:     NewFoodWarehouseRepo(db),
		medicineWarehouse: NewMedicineWarehouseRepo(db),
	}
}

func (s *FarmishStoragePg) AnimalType() repository.AnimalType {
	return s.animalType
}

func (s *FarmishStoragePg) Animals() repository.Animals {
	return s.animals
}

func (s *FarmishStoragePg) FoodWarehouse() repository.FoodWarehouse {
	return s.foodWarehouse
}

func (s *FarmishStoragePg) MedicineWarehouse() repository.MedicineWarehouse {
	return s.medicineWarehouse
}
