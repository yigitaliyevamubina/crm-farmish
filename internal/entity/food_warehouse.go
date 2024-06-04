package entity

import "time"

type FoodWarehouse struct {
	ID           string
	Name         string
	Quantity     int
	QuantityType string
	AnimalID     string
	AnimalType   string
	GroupFeeding bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type FoodWarehouseCreate struct {
	ID           string
	Name         string
	Quantity     int
	QuantityType string
	AnimalID     string
	AnimalType   string
	GroupFeeding bool
}

type UpdateFoodWarehouseReq struct {
	ID           string
	Name         string
	Quantity     int
	QuantityType string
	AnimalID     string
	AnimalType   string
	GroupFeeding bool
}

type ListFoodWarehouse struct {
	Count          int
	FoodWarehouses []FoodWarehouse
}
