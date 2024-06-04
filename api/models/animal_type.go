package models

type AnimalType struct {
	ID               string `json:"id"`
	Type             string `json:"type"`
	FeedingInterval  int    `json:"feeding_interval"`
	WateringInterval int    `json:"watering_interval"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
	DeletedAt        string `json:"deleted_at"`
}

type AnimalTypeCreate struct {
	Type             string `json:"type"`
	FeedingInterval  int    `json:"feeding_interval"`
	WateringInterval int    `json:"watering_interval"`
}

type UpdateAnimalTypeReq struct {
	ID               string `json:"id"`
	Type             string `json:"type"`
	FeedingInterval  int    `json:"feeding_interval"`
	WateringInterval int    `json:"watering_interval"`
}

type ListAnimalTypeRes struct {
	Count       int          `json:"count"`
	AnimalTypes []AnimalType `json:"animal_types"`
}
