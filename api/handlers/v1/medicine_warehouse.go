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

// CreateMedicineWarehouse ...
// @Summary Create Medicine Warehouse
// @Description CreateDoctor - Api for crete MedicineWarehouse
// @Tags Medicine Warehouse
// @Accept json
// @Produce json
// @Param AnimalTypeCreate body models.MedicineWarehouseCreate true "AnimalTypeCreate"
// @Success 200 {object} models.MedicineWarehouse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/medicine [post]
func (h *HandlerV1) CreateMedicineWarehouse(c *gin.Context) {
	var (
		body        models.MedicineWarehouseCreate
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "CreateMedicineWarehouse") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.medicineWarehouse.CreateMedicine(ctx, &entity.MedicineWarehouseCreate{
		Name:         body.Name,
		Quantity:     body.Quantity,
		QuantityType: body.QuantityType,
	})
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "CreateMedicineWarehouse") {
		return
	}

	c.JSON(http.StatusOK, &models.MedicineWarehouse{
		ID:           res.ID,
		Name:         res.Name,
		Quantity:     res.Quantity,
		QuantityType: res.QuantityType,
		CreatedAt:    e.Format(res.CreatedAt),
		UpdatedAt:    e.Format(res.UpdatedAt),
		DeletedAt:    e.Format(res.DeletedAt),
	})
}

// GetMedicineWarehouse ...
// @Summary Get Medicine Warehouse
// @Description CreateDoctor - Api for Get Medicine Warehouse
// @Tags Medicine Warehouse
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} models.MedicineWarehouse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/medicine/get [get]
func (h *HandlerV1) GetMedicineWarehouse(c *gin.Context) {
	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.medicineWarehouse.GetMedicine(ctx, &entity.FieldValueReq{
		Field:        "id",
		Value:        id,
		DeleteStatus: false,
	})

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "GetMedicineWarehouse") {
		return
	}

	c.JSON(http.StatusOK, &models.MedicineWarehouse{
		ID:           res.ID,
		Name:         res.Name,
		Quantity:     res.Quantity,
		QuantityType: res.QuantityType,
		CreatedAt:    e.Format(res.CreatedAt),
		UpdatedAt:    e.Format(res.UpdatedAt),
		DeletedAt:    e.Format(res.DeletedAt),
	})
}

// ListMedicineWarehouse ...
// @Summary List Medicine Warehouse
// @Description CreateDoctor - Api for List Medicine Warehouse
// @Tags Medicine Warehouse
// @Accept json
// @Produce json
// @Param ListReq query models.ListReq false "ListReq"
// @Success 200 {object} models.ListMedicineWarehouse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/medicine [get]
func (h *HandlerV1) ListMedicineWarehouse(c *gin.Context) {
	limit := c.Query("limit")
	page := c.Query("page")
	orderBy := c.Query("orderBy")

	pageInt, limitInt, err := e.ParseQueryParams(page, limit)
	if e.HandleError(c, err, h.log, http.StatusBadRequest, "ListMedicineWarehouse") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.medicineWarehouse.ListMedicine(ctx, &entity.ListReq{
		Page:         pageInt,
		Limit:        limitInt,
		DeleteStatus: false,
		OrderBy:      orderBy,
	})

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "ListMedicineWarehouse") {
		return
	}

	var medicines models.ListMedicineWarehouse
	for _, medicine := range res.MedicineWarehouses {
		var med models.MedicineWarehouse
		med.ID = medicine.ID
		med.Name = medicine.Name
		med.Quantity = medicine.Quantity
		med.QuantityType = medicine.QuantityType
		med.CreatedAt = e.Format(medicine.CreatedAt)
		med.UpdatedAt = e.Format(medicine.UpdatedAt)
		med.DeletedAt = e.Format(medicine.DeletedAt)
		medicines.MedicineWarehouses = append(medicines.MedicineWarehouses, med)
	}
	medicines.Count = res.Count
	c.JSON(http.StatusOK, medicines)
}

// UpdateMedicineWarehouse ...
// @Summary Update Medicine Warehouse
// @Description CreateDoctor - Api for Update Medicine Warehouse
// @Tags Medicine Warehouse
// @Accept json
// @Produce json
// @Param UpdateAnimalTypeReq body models.UpdateMedicineWarehouseReq true "UpdateAnimalTypeReq"
// @Success 200 {object} models.MedicineWarehouse
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/medicine [put]
func (h *HandlerV1) UpdateMedicineWarehouse(c *gin.Context) {
	var (
		body        models.UpdateAnimalReq
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "UpdateMedicineWarehouse") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.medicineWarehouse.UpdateMedicine(ctx, &entity.UpdateMedicineWarehouseReq{
		ID:           body.ID,
		Name:         body.Name,
		Quantity:     body.Weight,
		QuantityType: body.Disease,
	})

	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "UpdateMedicineWarehouse") {
		return
	}

	c.JSON(http.StatusOK, &models.MedicineWarehouse{
		ID:           res.ID,
		Name:         res.Name,
		Quantity:     res.Quantity,
		QuantityType: res.QuantityType,
		CreatedAt:    e.Format(res.CreatedAt),
		UpdatedAt:    e.Format(res.UpdatedAt),
		DeletedAt:    e.Format(res.DeletedAt),
	})
}

// DeleteMedicineWarehouse ...
// @Summary Delete Medicine Warehouse
// @Description CreateDoctor - Api for Delete Medicine Warehouse
// @Tags Medicine Warehouse
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} models.StatusRes
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/medicine [delete]
func (h *HandlerV1) DeleteMedicineWarehouse(c *gin.Context) {
	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.medicineWarehouse.DeleteMedicine(ctx, &entity.FieldValueReq{
		Field:        "id",
		Value:        id,
		DeleteStatus: false,
	})

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "UpdateMedicineWarehouse") {
		return
	}

	c.JSON(http.StatusOK, &models.StatusRes{Status: res.Status})
}
