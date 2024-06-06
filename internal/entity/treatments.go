package entity

import "time"

type Treatment struct {
	ID            string
	AnimalID      string
	MedicineID    string
	TreatmentTime time.Time
}

type ListTreatment struct {
	Count     int
	Treatment []Treatment
}

type ListTreatmentReq struct {
	Page  uint64
	Limit uint64
	Field string
	Value string
}
