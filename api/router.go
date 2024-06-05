package api

import (
	_ "crm-farmish/api/docs"
	"crm-farmish/internal/usecase"
	"time"

	v1 "crm-farmish/api/handlers/v1"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	"crm-farmish/internal/pkg/config"
)

type RouteOption struct {
	ContextTimeout    time.Duration
	Logger            *zap.Logger
	Config            *config.Config
	AnimalType        usecase.AnimalType
	Animals           usecase.Animals
	FoodWarehouse     usecase.FoodWarehouse
	MedicineWarehouse usecase.MedicineWarehouse
	Feeding           usecase.Feeding
}

// NewRoute
// @title CRM Farmish
// @version 1.7
// @host localhost:9050
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func NewRoute(option RouteOption) *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	HandlerV1 := v1.New(&v1.HandlerV1Config{
		Config:            option.Config,
		Logger:            option.Logger,
		ContextTimeout:    option.ContextTimeout,
		AnimalType:        option.AnimalType,
		Animals:           option.Animals,
		FoodWarehouse:     option.FoodWarehouse,
		MedicineWarehouse: option.MedicineWarehouse,
		Feeding:           option.Feeding,

		//BrokerProducer: option.BrokerProducer,
	})

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	corsConfig.AllowBrowserExtensions = true
	corsConfig.AllowMethods = []string{"*"}
	router.Use(cors.New(corsConfig))

	api := router.Group("/v1")

	animalType := api.Group("/animal-type")
	animalType.POST("", HandlerV1.CreateAnimalTypes)
	animalType.GET("/get", HandlerV1.GetAnimalTypes)
	animalType.GET("", HandlerV1.ListAnimalTypes)
	animalType.PUT("", HandlerV1.UpdateAnimalTypes)
	animalType.DELETE("", HandlerV1.DeleteAnimalTypes)

	animals := api.Group("/animals")
	animals.POST("", HandlerV1.CreateAnimals)
	animals.GET("/get", HandlerV1.GetAnimals)
	animals.GET("", HandlerV1.ListAnimals)
	animals.PUT("", HandlerV1.UpdateAnimals)
	animals.DELETE("", HandlerV1.DeleteAnimals)

	foodWarehouse := api.Group("/food")
	foodWarehouse.POST("", HandlerV1.CreateFoodWarehouse)
	foodWarehouse.GET("/get", HandlerV1.GetFoodWarehouse)
	foodWarehouse.GET("", HandlerV1.ListFoodWarehouse)
	foodWarehouse.PUT("", HandlerV1.UpdateFoodWarehouse)
	foodWarehouse.DELETE("", HandlerV1.DeleteFoodWarehouse)

	medicine := api.Group("/medicine")
	medicine.POST("", HandlerV1.CreateMedicineWarehouse)
	medicine.GET("/get", HandlerV1.GetMedicineWarehouse)
	medicine.GET("", HandlerV1.ListMedicineWarehouse)
	medicine.PUT("", HandlerV1.UpdateMedicineWarehouse)
	medicine.DELETE("", HandlerV1.DeleteMedicineWarehouse)

	feeding := api.Group("/feeding")
	feeding.POST("", HandlerV1.CreateFeeding)
	feeding.GET("/get", HandlerV1.GetFeeding)
	feeding.GET("", HandlerV1.ListFeeding)
	feeding.GET("/animal-id", HandlerV1.ListFeedingByAnimalID)

	url := ginSwagger.URL("swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}
