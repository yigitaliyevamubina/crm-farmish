package models

type Treatment struct {
	ID            string
	AnimalID      string
	MedicineID    string
	TreatmentTime string
}

type CreateTreatment struct {
	AnimalID   string
	MedicineID string
}

type ListTreatment struct {
	Count     int
	Treatment []Treatment
}

type ListTreatmentReq struct {
	Page  uint64
	Limit uint64
}
