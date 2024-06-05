package suit_tests

import (
	"context"
	"crm-farmish/internal/entity"
	repo "crm-farmish/internal/infrastructure/repository/postgresql"
	"crm-farmish/internal/pkg/config"
	db "crm-farmish/internal/pkg/postgres"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"log"
	"testing"
	"time"
)

type MedicinWareHouseTestSuite struct {
	suite.Suite
	CleanUpFunc func()
	Repository  *repo.MedicineWarehouse
}

func (s *MedicinWareHouseTestSuite) SetupTest() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Println(err)
		return
	}
	pgPool, err := db.New(cfg)
	if err != nil {
		log.Fatal(err)
		return
	}
	s.Repository = repo.NewMedicineWarehouseRepo(pgPool)
	s.CleanUpFunc = pgPool.Close
}

func (s *MedicinWareHouseTestSuite) TestAnimalTypesCrud() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(2))
	defer cancel()

	medicine := &entity.MedicineWarehouseCreate{
		ID:           uuid.NewString(),
		Name:         "Name",
		Quantity:     1,
		QuantityType: "Quantity Type",
	}
	resp, err := s.Repository.CreateMedicine(ctx, medicine)
	s.Suite.NoError(err)
	s.Suite.NotNil(resp)
	s.Suite.Equal(resp.ID, medicine.ID)
	s.Suite.Equal(resp.Name, medicine.Name)
	s.Suite.Equal(resp.Quantity, medicine.Quantity)
	s.Suite.Equal(resp.QuantityType, medicine.QuantityType)

	getMedicine, err := s.Repository.GetMedicine(ctx, &entity.FieldValueReq{
		Field:        "id",
		Value:        medicine.ID,
		DeleteStatus: false,
	})
	s.Suite.NoError(err)
	s.Suite.NotNil(getMedicine)
	s.Suite.Equal(getMedicine.ID, medicine.ID)
	s.Suite.Equal(getMedicine.Name, medicine.Name)
	s.Suite.Equal(getMedicine.Quantity, medicine.Quantity)
	s.Suite.Equal(getMedicine.QuantityType, medicine.QuantityType)

	respAll, err := s.Repository.ListMedicine(ctx, &entity.ListReq{
		Page:         1,
		Limit:        10,
		OrderBy:      "",
		DeleteStatus: false,
	})
	s.Suite.NoError(err)
	s.Suite.NotNil(respAll)

	newUpName := "Update Name"
	newUpQuantity := 3
	newUpQuantityType := "Update Quantity Type"

	updatedMedecine, err := s.Repository.UpdateMedicine(ctx, &entity.UpdateMedicineWarehouseReq{
		ID:           medicine.ID,
		Name:         newUpName,
		Quantity:     newUpQuantity,
		QuantityType: newUpQuantityType,
	})
	s.Suite.NoError(err)
	s.Suite.NotNil(updatedMedecine)
	s.Suite.NotNil(updatedMedecine.UpdatedAt)
	s.Suite.Equal(updatedMedecine.ID, medicine.ID)
	s.Suite.Equal(newUpName, updatedMedecine.Name)
	s.Suite.Equal(newUpQuantity, updatedMedecine.Quantity)
	s.Suite.Equal(newUpQuantityType, updatedMedecine.QuantityType)

	deleteDep, err := s.Repository.DeleteMedicine(ctx, &entity.FieldValueReq{
		Field:        "id",
		Value:        medicine.ID,
		DeleteStatus: true,
	})
	s.Suite.NotNil(deleteDep)
	s.Suite.NoError(err)
}

func (s *MedicinWareHouseTestSuite) TearDownTest() {
	s.CleanUpFunc()
}

func TestMedicinWareHouseTestSuite(t *testing.T) {
	suite.Run(t, new(MedicinWareHouseTestSuite))
}
