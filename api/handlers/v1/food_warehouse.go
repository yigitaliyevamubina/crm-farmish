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

// CreateFoodWarehouse ...
// @Summary Create Food Warehouse
// @Description CreateFoodWarehouse - Api for crete FoodWarehouse
// @Tags Food Warehouse
// @Accept json
// @Produce json
// @Param FoodWarehouseCreate body models.FoodWarehouseCreate true "FoodWarehouseCreate"
// @Success 200 {object} models.FoodWarehouse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/food [post]
func (h *HandlerV1) CreateFoodWarehouse(c *gin.Context) {
	var (
		body        models.FoodWarehouseCreate
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "CreateFoodWarehouse") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.foodWarehouse.CreateFoodWarehouse(ctx, &entity.FoodWarehouseCreate{
		Name:         body.Name,
		Quantity:     body.Quantity,
		QuantityType: body.QuantityType,
		AnimalID:     body.AnimalID,
		AnimalType:   body.AnimalTypeID,
		GroupFeeding: body.GroupFeeding,
	})
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "CreateFoodWarehouse") {
		return
	}

	c.JSON(http.StatusOK, &models.FoodWarehouse{
		ID:           res.ID,
		Name:         res.Name,
		Quantity:     res.Quantity,
		QuantityType: res.QuantityType,
		AnimalID:     res.AnimalID,
		AnimalTypeID: res.AnimalType,
		GroupFeeding: res.GroupFeeding,
		CreatedAt:    e.Format(res.CreatedAt),
		UpdatedAt:    e.Format(res.UpdatedAt),
		DeletedAt:    e.Format(res.DeletedAt),
	})
}

// GetFoodWarehouse ...
// @Summary Get Food Warehouse
// @Description GetFoodWarehouse - Api for Get FoodWarehouse
// @Tags Food Warehouse
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} models.FoodWarehouse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/food/get [get]
func (h *HandlerV1) GetFoodWarehouse(c *gin.Context) {
	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.foodWarehouse.GetFoodWarehouse(ctx, &entity.FieldValueReq{
		Field:        "id",
		Value:        id,
		DeleteStatus: false,
	})

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "GetFoodWarehouse") {
		return
	}

	c.JSON(http.StatusOK, &models.FoodWarehouse{
		ID:           res.ID,
		Name:         res.Name,
		Quantity:     res.Quantity,
		QuantityType: res.QuantityType,
		AnimalID:     res.AnimalID,
		AnimalTypeID: res.AnimalType,
		GroupFeeding: res.GroupFeeding,
		CreatedAt:    e.Format(res.CreatedAt),
		UpdatedAt:    e.Format(res.UpdatedAt),
		DeletedAt:    e.Format(res.DeletedAt),
	})
}

// ListFoodWarehouse ...
// @Summary List Food Warehouse
// @Description ListFoodWarehouse - Api for List FoodWarehouse
// @Tags Food Warehouse
// @Accept json
// @Produce json
// @Param ListReq query models.ListReq false "ListReq"
// @Success 200 {object} models.ListFoodWarehouse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/food [get]
func (h *HandlerV1) ListFoodWarehouse(c *gin.Context) {
	limit := c.Query("limit")
	page := c.Query("page")
	orderBy := c.Query("orderBy")

	pageInt, limitInt, err := e.ParseQueryParams(page, limit)
	if e.HandleError(c, err, h.log, http.StatusBadRequest, "ListFoodWarehouse") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.foodWarehouse.ListFoodWarehouse(ctx, &entity.ListReq{
		Page:         pageInt,
		Limit:        limitInt,
		DeleteStatus: false,
		OrderBy:      orderBy,
	})

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "ListFoodWarehouse") {
		return
	}

	var foods models.ListFoodWarehouse
	for _, food := range res.FoodWarehouses {
		var foodRes models.FoodWarehouse
		foodRes.ID = food.ID
		foodRes.Name = food.Name
		foodRes.Quantity = food.Quantity
		foodRes.QuantityType = food.QuantityType
		foodRes.AnimalID = food.AnimalID
		foodRes.AnimalTypeID = food.AnimalType
		foodRes.GroupFeeding = food.GroupFeeding
		foodRes.CreatedAt = e.Format(food.CreatedAt)
		foodRes.UpdatedAt = e.Format(food.UpdatedAt)
		foodRes.DeletedAt = e.Format(food.DeletedAt)
		foods.FoodWarehouses = append(foods.FoodWarehouses, foodRes)
	}
	foods.Count = res.Count
	c.JSON(http.StatusOK, foods)
}

// UpdateFoodWarehouse ...
// @Summary Update FoodWarehouse
// @Description UpdateFoodWarehouse - Api for Update FoodWarehouse
// @Tags Food Warehouse
// @Accept json
// @Produce json
// @Param UpdateAnimalTypeReq body models.UpdateFoodWarehouseReq true "UpdateAnimalTypeReq"
// @Success 200 {object} models.FoodWarehouse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/food [put]
func (h *HandlerV1) UpdateFoodWarehouse(c *gin.Context) {
	var (
		body        models.UpdateFoodWarehouseReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "UpdateFoodWarehouse") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.foodWarehouse.UpdateFoodWarehouse(ctx, &entity.UpdateFoodWarehouseReq{
		ID:           body.ID,
		Name:         body.Name,
		Quantity:     body.Quantity,
		QuantityType: body.QuantityType,
		AnimalID:     body.AnimalID,
		AnimalType:   body.AnimalType,
		GroupFeeding: body.GroupFeeding,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "UpdateFoodWarehouse") {
		return
	}

	c.JSON(http.StatusOK, &models.FoodWarehouse{
		ID:           res.ID,
		Name:         res.Name,
		Quantity:     res.Quantity,
		QuantityType: res.QuantityType,
		AnimalID:     res.AnimalID,
		AnimalTypeID: res.AnimalType,
		GroupFeeding: res.GroupFeeding,
		CreatedAt:    e.Format(res.CreatedAt),
		UpdatedAt:    e.Format(res.UpdatedAt),
		DeletedAt:    e.Format(res.DeletedAt),
	})
}

// DeleteFoodWarehouse ...
// @Summary Delete FoodWarehouse
// @Description DeleteFoodWarehouse - Api for Delete FoodWarehouse
// @Tags Food Warehouse
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} models.StatusRes
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/food [delete]
func (h *HandlerV1) DeleteFoodWarehouse(c *gin.Context) {
	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.foodWarehouse.DeleteFoodWarehouse(ctx, &entity.FieldValueReq{
		Field:        "id",
		Value:        id,
		DeleteStatus: false,
	})

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "UpdateFoodWarehouse") {
		return
	}

	c.JSON(http.StatusOK, &models.StatusRes{Status: res.Status})
}
