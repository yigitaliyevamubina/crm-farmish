package models

type FoodWarehouse struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Quantity     int    `json:"quantity"`
	QuantityType string `json:"quantity_type"`
	AnimalID     string `json:"animal_id"`
	AnimalType   string `json:"animal_type"`
	GroupFeeding bool   `json:"group_feeding"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at"`
}

type FoodWarehouseCreate struct {
	Name         string `json:"name"`
	Quantity     int    `json:"quantity"`
	QuantityType string `json:"quantity_type"`
	AnimalID     string `json:"animal_id"`
	AnimalType   string `json:"animal_type"`
	GroupFeeding bool   `json:"group_feeding"`
}

type UpdateFoodWarehouseReq struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Quantity     int    `json:"quantity"`
	QuantityType string `json:"quantity_type"`
	AnimalID     string `json:"animal_id"`
	AnimalType   string `json:"animal_type"`
	GroupFeeding bool   `json:"group_feeding"`
}

type ListFoodWarehouse struct {
	Count          int             `json:"count"`
	FoodWarehouses []FoodWarehouse `json:"food_warehouses"`
}
