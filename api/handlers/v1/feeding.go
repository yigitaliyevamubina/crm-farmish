package v1

import (
	"context"
	e "crm-farmish/api/handlers/regtool"
	"crm-farmish/api/models"
	"crm-farmish/internal/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/encoding/protojson"
	"net/http"
	"time"
)

// CreateFeeding ...
// @Summary Create Feeding
// @Description CreateFeeding - Api for crete Feeding
// @Tags Feeding
// @Accept json
// @Produce json
// @Param AnimalTypeCreate body models.CreateFeeding true "AnimalTypeCreate"
// @Success 200 {object} models.Animal
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/feeding [post]
func (h *HandlerV1) CreateFeeding(c *gin.Context) {
	var (
		body        models.CreateFeeding
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "CreateFeeding") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.feeding.CreateFeeding(ctx, &entity.Feeding{
		MaleID:   body.MaleID,
		AnimalID: body.AnimalID,
	})
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "CreateFeeding") {
		return
	}

	c.JSON(http.StatusOK, &models.Feeding{
		ID:          res.ID,
		MaleID:      res.MaleID,
		AnimalID:    res.AnimalID,
		FeedingTime: e.Format(res.FeedingTime),
	})
}

// GetFeeding ...
// @Summary Get Feeding
// @Description GetFeeding - Api for Get Feeding
// @Tags Feeding
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} models.Feeding
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/feeding/get [get]
func (h *HandlerV1) GetFeeding(c *gin.Context) {
	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.feeding.GetFeeding(ctx, &entity.FieldValueReq{
		Field: "id",
		Value: id,
	})

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "GetFeeding") {
		return
	}

	c.JSON(http.StatusOK, &models.Feeding{
		ID:          res.ID,
		MaleID:      res.MaleID,
		AnimalID:    res.AnimalID,
		FeedingTime: e.Format(res.FeedingTime),
	})
}

// ListFeeding ...
// @Summary List Feeding
// @Description ListFeedingBy - API for listing feedings
// @Tags Feeding
// @Accept json
// @Produce json
// @Param limit query string false "limit"
// @Param page query string false "page"
// @Success 200 {object} models.ListFeeding
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/feeding [get]
func (h *HandlerV1) ListFeeding(c *gin.Context) {
	limit := c.Query("limit")
	page := c.Query("page")
	pageInt, limitInt, err := e.ParseQueryParams(page, limit)
	if e.HandleError(c, err, h.log, http.StatusBadRequest, "ListFeeding") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.feeding.ListFeedings(ctx, &entity.ListReq{
		Page:  pageInt,
		Limit: limitInt,
	})

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "ListFeeding") {
		return
	}

	var feedings models.ListFeeding
	for _, feedingRes := range res.Feedings {
		var feeding models.Feeding
		feeding.ID = feedingRes.ID
		feeding.MaleID = feedingRes.MaleID
		feeding.AnimalID = feedingRes.AnimalID
		feeding.FeedingTime = e.Format(feedingRes.FeedingTime)
		feedings.Feedings = append(feedings.Feedings, feeding)
	}
	feedings.Count = res.Count
	c.JSON(http.StatusOK, feedings)
}

// ListFeedingByAnimalID ...
// @Summary List Feeding By AnimalID
// @Description ListFeedingByAnimalID - Api for List Feeding By AnimalID
// @Tags Feeding
// @Accept json
// @Produce json
// @Param ListReq query models.ListFeedingByAnimalIDReq false "ListReq"
// @Success 200 {object} models.ListFeeding
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/feeding/animal-id [get]
func (h *HandlerV1) ListFeedingByAnimalID(c *gin.Context) {
	limit := c.Query("limit")
	page := c.Query("page")
	animalID := c.Query("animal_id")

	pageInt, limitInt, err := e.ParseQueryParams(page, limit)
	if e.HandleError(c, err, h.log, http.StatusBadRequest, "ListFeedingByAnimalID") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.feeding.ListFeedingsByAnimalID(ctx, &entity.ListFeedingByAnimalIDReq{
		Page:  pageInt,
		Limit: limitInt,
		ID:    animalID,
	})

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "ListFeedingByAnimalID") {
		return
	}

	var feedings models.ListFeeding
	for _, feedingRes := range res.Feedings {
		var feeding models.Feeding
		feeding.ID = feedingRes.ID
		feeding.MaleID = feedingRes.MaleID
		feeding.AnimalID = feedingRes.AnimalID
		feeding.FeedingTime = e.Format(feedingRes.FeedingTime)
		feedings.Feedings = append(feedings.Feedings, feeding)
	}
	feedings.Count = res.Count
	c.JSON(http.StatusOK, feedings)
}

// CreateWatering ...
// @Summary Create Watering
// @Description Create Watering - Api for crete Feeding
// @Tags Feeding
// @Accept json
// @Produce json
// @Param AnimalTypeCreate body models.CreateWatering true "AnimalTypeCreate"
// @Success 200 {object} models.Watering
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/watering [post]
func (h *HandlerV1) CreateWatering(c *gin.Context) {
	var (
		body        models.CreateWatering
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "CreateWatering") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.feeding.CreateWatering(ctx, &entity.Watering{
		AnimalID: body.AnimalID,
	})
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "CreateWatering") {
		return
	}

	c.JSON(http.StatusOK, &models.Watering{
		ID:           res.ID,
		AnimalID:     res.AnimalID,
		WateringTime: e.Format(res.WateringTime),
	})
}

// GetWatering ...
// @Summary Get Watering
// @Description GetWatering - Api for Get Watering
// @Tags Feeding
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} models.Watering
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/watering/get [get]
func (h *HandlerV1) GetWatering(c *gin.Context) {
	id := c.Query("id")

	fmt.Println(id)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.feeding.GetWatering(ctx, &entity.FieldValueReq{
		Field: "id",
		Value: id,
	})

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "GetWatering") {
		return
	}

	c.JSON(http.StatusOK, &models.Watering{
		ID:           res.ID,
		AnimalID:     res.AnimalID,
		WateringTime: e.Format(res.WateringTime),
	})
}

// NotFeeding ...
// @Summary List NotFeeding
// @Description ListFeedingBy - API for listing Not Feeding
// @Tags Feeding
// @Accept json
// @Produce json
// @Success 200 {object} models.ListFeeding
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/feeding/not-feeding [get]
func (h *HandlerV1) NotFeeding(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.feeding.NotFeedings(ctx)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "NotFeeding") {
		return
	}

	var feedings models.ListFeeding
	for _, feedingRes := range res.Feedings {
		var feeding models.Feeding
		feeding.ID = feedingRes.ID
		feeding.AnimalID = feedingRes.AnimalID
		feedings.Feedings = append(feedings.Feedings, feeding)
	}
	feedings.Count = res.Count
	c.JSON(http.StatusOK, feedings)
}

// NotWatering ...
// @Summary List NotWatering
// @Description ListFeedingBy - API for listing Not Watering
// @Tags Feeding
// @Accept json
// @Produce json
// @Success 200 {object} models.ListWatering
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/watering/not-watering [get]
func (h *HandlerV1) NotWatering(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.feeding.NotWatered(ctx)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "NotWatering") {
		return
	}

	var waterings models.ListFeeding
	for _, wateringRes := range res.Watering {
		var watering models.Feeding
		watering.ID = wateringRes.ID
		watering.AnimalID = wateringRes.AnimalID
		waterings.Feedings = append(waterings.Feedings, watering)
	}
	waterings.Count = res.Count
	c.JSON(http.StatusOK, waterings)
}
