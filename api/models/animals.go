package models

type Animal struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	Name            string `json:"name"`
	Gender          string `json:"gender"`
	Weight          int    `json:"weight"`
	LastFedTime     string `json:"last_fed_time"`
	LastWateredTime string `json:"last_watered_time"`
	Disease         string `json:"disease"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
	DeletedAt       string `json:"deleted_at"`
}

type AnimalsCreate struct {
	Type            string `json:"type"`
	Name            string `json:"name"`
	Gender          string `json:"gender"`
	Weight          int    `json:"weight"`
	LastFedTime     string `json:"last_fed_time"`
	LastWateredTime string `json:"last_watered_time"`
	Disease         string `json:"disease"`
}

type UpdateAnimalReq struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Gender          string `json:"gender"`
	Weight          int    `json:"weight"`
	LastFedTime     string `json:"last_fed_time"`
	LastWateredTime string `json:"last_watered_time"`
	Disease         string `json:"disease"`
}

type ListAnimalsRes struct {
	Count   int      `json:"count"`
	Animals []Animal `json:"animals"`
}
