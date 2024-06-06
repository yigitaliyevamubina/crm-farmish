package v1

import (
	"crm-farmish/internal/pkg/config"
	"crm-farmish/internal/usecase"
	"time"

	"go.uber.org/zap"
)

type HandlerV1 struct {
	ContextTimeout    time.Duration
	log               *zap.Logger
	cfg               *config.Config
	animalType        usecase.AnimalType
	animals           usecase.Animals
	foodWarehouse     usecase.FoodWarehouse
	medicineWarehouse usecase.MedicineWarehouse
	feeding           usecase.Feeding
	treatment         usecase.Treatment
}

// HandlerV1Config ...
type HandlerV1Config struct {
	ContextTimeout    time.Duration
	Logger            *zap.Logger
	Config            *config.Config
	AnimalType        usecase.AnimalType
	Animals           usecase.Animals
	FoodWarehouse     usecase.FoodWarehouse
	MedicineWarehouse usecase.MedicineWarehouse
	Feeding           usecase.Feeding
	Treatment         usecase.Treatment
}

// New ...
func New(c *HandlerV1Config) *HandlerV1 {
	return &HandlerV1{
		ContextTimeout:    c.ContextTimeout,
		log:               c.Logger,
		cfg:               c.Config,
		animalType:        c.AnimalType,
		animals:           c.Animals,
		foodWarehouse:     c.FoodWarehouse,
		medicineWarehouse: c.MedicineWarehouse,
		feeding:           c.Feeding,
		treatment:         c.Treatment,
	}
}
