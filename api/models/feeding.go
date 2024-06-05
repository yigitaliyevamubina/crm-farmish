package models

type Feeding struct {
	ID          string `json:"id"`
	MaleID      string `json:"male_id"`
	AnimalID    string `json:"animal_id"`
	FeedingTime string `json:"feeding_time"`
}

type CreateFeeding struct {
	MaleID   string `json:"male_id"`
	AnimalID string `json:"animal_id"`
}

type ListFeeding struct {
	Count    int       `json:"count"`
	Feedings []Feeding `json:"feedings"`
}

type ListFeedingByAnimalIDReq struct {
	Page         uint64 `json:"page"`
	Limit        uint64 `json:"limit"`
	DeleteStatus bool   `json:"delete_status"`
	AnimalID     string `json:"animal_id"`
}
