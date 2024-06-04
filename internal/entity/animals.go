package entity

import "time"

type Animals struct {
	ID              string
	Type            string
	Name            string
	Gender          string
	Weight          int
	LastFedTime     time.Time
	LastWateredTime time.Time
	Disease         string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       time.Time
}

type AnimalsCreate struct {
	ID              string
	Type            string
	Name            string
	Gender          string
	Weight          int
	LastFedTime     time.Time
	LastWateredTime time.Time
	Disease         string
}

type UpdateAnimalReq struct {
	ID              string
	Name            string
	Gender          string
	Weight          int
	LastFedTime     time.Time
	LastWateredTime time.Time
	Disease         string
}

type ListAnimalRes struct {
	Count   int
	Animals []Animals
}
