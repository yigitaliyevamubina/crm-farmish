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

// CreateAnimalTypes ...
// @Summary Create Animal Types
// @Description CreateDoctor - Api for crete Animal Types
// @Tags Animal Types
// @Accept json
// @Produce json
// @Param AnimalTypeCreate body models.AnimalTypeCreate true "AnimalTypeCreate"
// @Success 200 {object} models.AnimalType
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/animal-type [post]
func (h *HandlerV1) CreateAnimalTypes(c *gin.Context) {
	var (
		body        models.AnimalTypeCreate
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "CreateAnimalTypes") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.animalType.CreateAnimalTypes(ctx, &entity.AnimalTypeCreate{
		Type:             body.Type,
		FeedingInterval:  body.FeedingInterval,
		WateringInterval: body.WateringInterval,
	})
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "CreateAnimalTypes") {
		return
	}

	c.JSON(http.StatusOK, &models.AnimalType{
		ID:               res.ID,
		Type:             res.Type,
		FeedingInterval:  res.FeedingInterval,
		WateringInterval: res.WateringInterval,
		CreatedAt:        e.Format(res.CreatedAt),
		UpdatedAt:        e.Format(res.UpdatedAt),
		DeletedAt:        e.Format(res.DeletedAt),
	})
}

// GetAnimalTypes ...
// @Summary Get Animal Types
// @Description CreateDoctor - Api for Get Animal Types
// @Tags Animal Types
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} models.AnimalType
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/animal-type/get [get]
func (h *HandlerV1) GetAnimalTypes(c *gin.Context) {
	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.animalType.GetAnimalTypes(ctx, &entity.FieldValueReq{
		Field:        "id",
		Value:        id,
		DeleteStatus: false,
	})

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "GetAnimalTypes") {
		return
	}

	c.JSON(http.StatusOK, &models.AnimalType{
		ID:               res.ID,
		Type:             res.Type,
		FeedingInterval:  res.FeedingInterval,
		WateringInterval: res.WateringInterval,
		CreatedAt:        e.Format(res.CreatedAt),
		UpdatedAt:        e.Format(res.UpdatedAt),
		DeletedAt:        e.Format(res.DeletedAt),
	})
}

// ListAnimalTypes ...
// @Summary Lis tAnimal Types
// @Description CreateDoctor - Api for List Animal Types
// @Tags Animal Types
// @Accept json
// @Produce json
// @Param ListReq query models.ListReq false "ListReq"
// @Success 200 {object} models.ListAnimalTypeRes
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/animal-type [get]
func (h *HandlerV1) ListAnimalTypes(c *gin.Context) {
	limit := c.Query("limit")
	page := c.Query("page")
	orderBy := c.Query("orderBy")

	pageInt, limitInt, err := e.ParseQueryParams(page, limit)
	if e.HandleError(c, err, h.log, http.StatusBadRequest, "ListAnimalTypes") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.animalType.ListAnimalTypes(ctx, &entity.ListReq{
		Page:         pageInt,
		Limit:        limitInt,
		DeleteStatus: false,
		OrderBy:      orderBy,
	})

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "ListAnimalTypes") {
		return
	}

	var animalTypes models.ListAnimalTypeRes
	for _, animalType := range res.AnimalTypes {
		var animaltype models.AnimalType
		animaltype.ID = animalType.ID
		animaltype.Type = animalType.Type
		animaltype.FeedingInterval = animalType.FeedingInterval
		animaltype.WateringInterval = animalType.WateringInterval
		animaltype.CreatedAt = e.Format(animalType.CreatedAt)
		animaltype.UpdatedAt = e.Format(animalType.UpdatedAt)
		animaltype.DeletedAt = e.Format(animalType.DeletedAt)
		animalTypes.AnimalTypes = append(animalTypes.AnimalTypes, animaltype)
	}
	animalTypes.Count = res.Count
	c.JSON(http.StatusOK, animalTypes)
}

// UpdateAnimalTypes ...
// @Summary Update Animal Types
// @Description CreateDoctor - Api for Update Animal Types
// @Tags Animal Types
// @Accept json
// @Produce json
// @Param UpdateAnimalTypeReq body models.UpdateAnimalTypeReq true "UpdateAnimalTypeReq"
// @Success 200 {object} models.AnimalType
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/animal-type [put]
func (h *HandlerV1) UpdateAnimalTypes(c *gin.Context) {
	var (
		body        models.UpdateAnimalTypeReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "UpdateAnimalTypes") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.animalType.UpdateAnimalsType(ctx, &entity.UpdateAnimalTypeReq{
		ID:               body.ID,
		Type:             body.Type,
		FeedingInterval:  body.FeedingInterval,
		WateringInterval: body.WateringInterval,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "UpdateAnimalTypes") {
		return
	}

	c.JSON(http.StatusOK, &models.AnimalType{
		ID:               res.ID,
		Type:             res.Type,
		FeedingInterval:  res.FeedingInterval,
		WateringInterval: res.WateringInterval,
		CreatedAt:        e.Format(res.CreatedAt),
		UpdatedAt:        e.Format(res.UpdatedAt),
		DeletedAt:        e.Format(res.DeletedAt),
	})
}

// DeleteAnimalTypes ...
// @Summary Delete Animal Types
// @Description CreateDoctor - Api for Delete Animal Types
// @Tags Animal Types
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} models.StatusRes
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/animal-type [delete]
func (h *HandlerV1) DeleteAnimalTypes(c *gin.Context) {
	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.animalType.DeleteAnimalTypes(ctx, &entity.FieldValueReq{
		Field:        "id",
		Value:        id,
		DeleteStatus: false,
	})

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "UpdateAnimalTypes") {
		return
	}

	c.JSON(http.StatusOK, &models.StatusRes{Status: res.Status})
}
