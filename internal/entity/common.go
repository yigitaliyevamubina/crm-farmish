package entity

type FieldValueReq struct {
	Field        string
	Value        string
	DeleteStatus bool
}

type ListReq struct {
	Page         uint64
	Limit        uint64
	DeleteStatus bool
	OrderBy      string
}

type StatusRes struct {
	Status bool
}
