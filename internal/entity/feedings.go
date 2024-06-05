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
	Page         uint64
	Limit        uint64
	DeleteStatus bool
	ID           string
}
