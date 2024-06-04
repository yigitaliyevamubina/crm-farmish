package models

type MedicineWarehouse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Quantity     int    `json:"quantity"`
	QuantityType string `json:"quantity_type"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at"`
}

type MedicineWarehouseCreate struct {
	Name         string `json:"name"`
	Quantity     int    `json:"quantity"`
	QuantityType string `json:"quantity_type"`
}

type UpdateMedicineWarehouseReq struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Quantity     int    `json:"quantity"`
	QuantityType string `json:"quantity_type"`
}

type ListMedicineWarehouse struct {
	Count              int                 `json:"count"`
	MedicineWarehouses []MedicineWarehouse `json:"medicine_warehouses"`
}
