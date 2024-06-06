package entity

import "time"

type Feeding struct {
	ID          string
	MaleID      string
	AnimalID    string
	FeedingTime time.Time
}

type ListFeeding struct {
	Count    int
	Feedings []Feeding
}

type ListFeedingByAnimalIDReq struct {
	Page  uint64
	Limit uint64
	ID    string
}

type Watering struct {
	ID           string
	AnimalID     string
	WateringTime time.Time
}

type ListWatering struct {
	Count    int
	Watering []Watering
}
