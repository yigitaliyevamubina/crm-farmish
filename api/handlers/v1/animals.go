package v1

import (
	"context"
	e "crm-farmish/api/handlers/regtool"
	"crm-farmish/api/models"
	"crm-farmish/internal/entity"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"time"
)

// CreateAnimals ...
// @Summary Create Animal s
// @Description CreateDoctor - Api for crete Animals
// @Tags Animals
// @Accept json
// @Produce json
// @Param AnimalTypeCreate body models.AnimalsCreate true "AnimalTypeCreate"
// @Success 200 {object} models.Animal
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/animals [post]
func (h *HandlerV1) CreateAnimals(c *gin.Context) {
	var (
		body        models.AnimalsCreate
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "CreateAnimals") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	lastFedTime, err := time.Parse("2006-01-02 15:04:05", body.LastFedTime)
	if e.HandleError(c, err, h.log, http.StatusBadRequest, "CreateAnimals") {
		return
	}

	lastWateredTime, err := time.Parse("2006-01-02 15:04:05", body.LastWateredTime)
	if e.HandleError(c, err, h.log, http.StatusBadRequest, "CreateAnimals") {
		return
	}

	res, err := h.animals.CreateAnimals(ctx, &entity.AnimalsCreate{
		Type:            body.AnimalTypeID,
		Name:            body.Name,
		Gender:          body.Gender,
		Weight:          body.Weight,
		LastFedTime:     lastFedTime,
		LastWateredTime: lastWateredTime,
		Disease:         body.Disease,
	})
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "CreateAnimals") {
		return
	}

	c.JSON(http.StatusOK, &models.Animal{
		ID:              res.ID,
		AnimalTypeID:    res.Type,
		Name:            res.Name,
		Gender:          res.Gender,
		Weight:          res.Weight,
		LastFedTime:     e.Format(res.LastWateredTime),
		LastWateredTime: e.Format(res.LastFedTime),
		Disease:         res.Disease,
		CreatedAt:       e.Format(res.CreatedAt),
		UpdatedAt:       e.Format(res.UpdatedAt),
		DeletedAt:       e.Format(res.DeletedAt),
	})
}

// GetAnimals ...
// @Summary Get Animal s
// @Description CreateDoctor - Api for Get Animals
// @Tags Animals
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} models.Animal
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/animals/get [get]
func (h *HandlerV1) GetAnimals(c *gin.Context) {
	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.animals.GetAnimals(ctx, &entity.FieldValueReq{
		Field:        "id",
		Value:        id,
		DeleteStatus: false,
	})

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "GetAnimals") {
		return
	}

	c.JSON(http.StatusOK, &models.Animal{
		ID:              res.ID,
		AnimalTypeID:    res.Type,
		Name:            res.Name,
		Gender:          res.Gender,
		Weight:          res.Weight,
		LastFedTime:     e.Format(res.LastWateredTime),
		LastWateredTime: e.Format(res.LastFedTime),
		Disease:         res.Disease,
		CreatedAt:       e.Format(res.CreatedAt),
		UpdatedAt:       e.Format(res.UpdatedAt),
		DeletedAt:       e.Format(res.DeletedAt),
	})
}

// ListAnimals ...
// @Summary Lis tAnimals
// @Description CreateDoctor - Api for List Animals
// @Tags Animals
// @Accept json
// @Produce json
// @Param ListReq query models.ListReq false "ListReq"
// @Success 200 {object} models.ListAnimalsRes
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/animals [get]
func (h *HandlerV1) ListAnimals(c *gin.Context) {
	limit := c.Query("limit")
	page := c.Query("page")
	orderBy := c.Query("orderBy")

	pageInt, limitInt, err := e.ParseQueryParams(page, limit)
	if e.HandleError(c, err, h.log, http.StatusBadRequest, "ListAnimals") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.animals.ListAnimals(ctx, &entity.ListReq{
		Page:         pageInt,
		Limit:        limitInt,
		DeleteStatus: false,
		OrderBy:      orderBy,
	})

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "ListAnimals") {
		return
	}

	var animals models.ListAnimalsRes
	for _, animal := range res.Animals {
		var animalRes models.Animal
		animalRes.ID = animal.ID
		animalRes.AnimalTypeID = animal.Type
		animalRes.Name = animal.Name
		animalRes.Gender = animal.Gender
		animalRes.Weight = animal.Weight
		animalRes.LastWateredTime = e.Format(animal.LastWateredTime)
		animalRes.LastFedTime = e.Format(animal.LastFedTime)
		animalRes.Disease = animal.Disease
		animalRes.CreatedAt = e.Format(animal.CreatedAt)
		animalRes.UpdatedAt = e.Format(animal.UpdatedAt)
		animalRes.DeletedAt = e.Format(animal.DeletedAt)
		animals.Animals = append(animals.Animals, animalRes)
	}
	animals.Count = res.Count
	c.JSON(http.StatusOK, animals)
}

// UpdateAnimals ...
// @Summary Update Animals
// @Description CreateDoctor - Api for Update Animals
// @Tags Animals
// @Accept json
// @Produce json
// @Param UpdateAnimalTypeReq body models.UpdateAnimalReq true "UpdateAnimalTypeReq"
// @Success 200 {object} models.Animal
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/animals [put]
func (h *HandlerV1) UpdateAnimals(c *gin.Context) {
	var (
		body        models.UpdateAnimalReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "UpdateAnimals") {
		return
	}

	lastFedTime, err := time.Parse("2006-01-02 15:04:05", body.LastFedTime)
	if e.HandleError(c, err, h.log, http.StatusBadRequest, "CreateAnimals") {
		return
	}

	lastWateredTime, err := time.Parse("2006-01-02 15:04:05", body.LastWateredTime)
	if e.HandleError(c, err, h.log, http.StatusBadRequest, "CreateAnimals") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.animals.UpdateAnimals(ctx, &entity.UpdateAnimalReq{
		ID:              body.ID,
		Name:            body.Name,
		Gender:          body.Gender,
		Weight:          body.Weight,
		LastFedTime:     lastFedTime,
		LastWateredTime: lastWateredTime,
		Disease:         body.Disease,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "UpdateAnimals") {
		return
	}

	c.JSON(http.StatusOK, &models.Animal{
		ID:              res.ID,
		AnimalTypeID:    res.Type,
		Name:            res.Name,
		Gender:          res.Gender,
		Weight:          res.Weight,
		LastFedTime:     e.Format(res.LastWateredTime),
		LastWateredTime: e.Format(res.LastFedTime),
		Disease:         res.Disease,
		CreatedAt:       e.Format(res.CreatedAt),
		UpdatedAt:       e.Format(res.UpdatedAt),
		DeletedAt:       e.Format(res.DeletedAt),
	})
}

// DeleteAnimals ...
// @Summary Delete Animals
// @Description CreateDoctor - Api for Delete Animals
// @Tags Animals
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} models.StatusRes
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/animals [delete]
func (h *HandlerV1) DeleteAnimals(c *gin.Context) {
	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.animals.DeleteAnimals(ctx, &entity.FieldValueReq{
		Field:        "id",
		Value:        id,
		DeleteStatus: false,
	})

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "UpdateAnimals") {
		return
	}

	c.JSON(http.StatusOK, &models.StatusRes{Status: res.Status})
}
