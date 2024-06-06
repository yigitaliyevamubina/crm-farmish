package app

import (
	"context"
	"crm-farmish/api"
	repo "crm-farmish/internal/infrastructure/repository/postgresql"
	"crm-farmish/internal/pkg/config"
	"crm-farmish/internal/pkg/logger"
	"crm-farmish/internal/pkg/postgres"
	"crm-farmish/internal/usecase"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type App struct {
	Config            *config.Config
	Logger            *zap.Logger
	DB                *postgres.PostgresDB
	server            *http.Server
	AnimalType        usecase.AnimalType
	Animals           usecase.Animals
	FoodWarehouse     usecase.FoodWarehouse
	MedicineWarehouse usecase.MedicineWarehouse
}

func NewApp(cfg *config.Config) (*App, error) {
	// l init
	l, err := logger.New(cfg.LogLevel, cfg.Environment, cfg.APP+".log")
	if err != nil {
		return nil, err
	}

	// postgres init
	db, err := postgres.New(cfg)
	if err != nil {
		return nil, err
	}

	return &App{
		Config: cfg,
		Logger: l,
		DB:     db,
	}, nil
}

func (a *App) Run() error {
	contextTimeout, err := time.ParseDuration("2s")
	if err != nil {
		return fmt.Errorf("error while parsing context timeout: %v", err)
	}

	// repositories initialization

	animalType := repo.NewAnimalTypesRepo(a.DB)

	animals := repo.NewAnimalsRepo(a.DB)

	foodWarehouse := repo.NewFoodWarehouseRepo(a.DB)

	medicineWarehouse := repo.NewMedicineWarehouseRepo(a.DB)

	feeding := repo.NewFeedingRepo(a.DB)

	treatment := repo.NewTreatmentRepo(a.DB)

	// use case initialization

	animalTypeUseCase := usecase.NewAnimalTypeUseCase(animalType)

	animalsUseCase := usecase.NewAnimalsUseCase(animals)

	foodWarehouseUseCase := usecase.NewFoodWarehouseUseCase(foodWarehouse)

	medicineWarehouseUseCase := usecase.NewMedicineWarehouseUseCase(medicineWarehouse)

	feedingUseCase := usecase.NewFeedingUseCase(feeding)

	treatmentUseCase := usecase.NewTreatmentUseCase(treatment)

	// api init
	handler := api.NewRoute(api.RouteOption{
		ContextTimeout:    contextTimeout,
		Logger:            a.Logger,
		Config:            a.Config,
		AnimalType:        animalTypeUseCase,
		Animals:           animalsUseCase,
		FoodWarehouse:     foodWarehouseUseCase,
		MedicineWarehouse: medicineWarehouseUseCase,
		Feeding:           feedingUseCase,
		Treatment:         treatmentUseCase,
	})

	// server init
	a.server, err = api.NewServer(a.Config, handler)
	if err != nil {
		return fmt.Errorf("error while initializing server: %v", err)
	}

	return a.server.ListenAndServe()
}

func (a *App) Stop() {

	// close database
	a.DB.Close()

	// shutdown server http
	if err := a.server.Shutdown(context.Background()); err != nil {
		a.Logger.Error("shutdown server http ", zap.Error(err))
	}

	// zap logger sync
	a.Logger.Sync()
}
