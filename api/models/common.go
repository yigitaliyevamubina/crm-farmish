package models

type FieldValueReq struct {
	Field        string `json:"field"`
	Value        string `json:"value"`
	DeleteStatus bool   `json:"delete_status"`
}

type ListReq struct {
	Page         uint64 `json:"page"`
	Limit        uint64 `json:"limit"`
	DeleteStatus bool   `json:"delete_status"`
	OrderBy      string `json:"order_by"`
}

type StatusRes struct {
	Status bool `json:"status"`
}
