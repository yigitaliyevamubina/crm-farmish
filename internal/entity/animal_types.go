package entity

import "time"

type AnimalType struct {
	ID               string
	Type             string
	FeedingInterval  int
	WateringInterval int
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        time.Time
}

type AnimalTypeCreate struct {
	ID               string
	Type             string
	FeedingInterval  int
	WateringInterval int
}

type UpdateAnimalTypeReq struct {
	ID               string
	Type             string
	FeedingInterval  int
	WateringInterval int
}

type ListAnimalTypeRes struct {
	Count       int
	AnimalTypes []AnimalType
}
