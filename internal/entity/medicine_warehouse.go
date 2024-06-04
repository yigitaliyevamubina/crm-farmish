package entity

import "time"

type MedicineWarehouse struct {
	ID           string
	Name         string
	Quantity     int
	QuantityType string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    time.Time
}

type MedicineWarehouseCreate struct {
	ID           string
	Name         string
	Quantity     int
	QuantityType string
}

type UpdateMedicineWarehouseReq struct {
	ID           string
	Name         string
	Quantity     int
	QuantityType string
}

type ListMedicineWarehouse struct {
	Count              int
	MedicineWarehouses []MedicineWarehouse
}
