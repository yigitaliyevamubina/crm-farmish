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

// CreateTreatment ...
// @Summary Create Treatment
// @Description CreateTreatment - Api for crete Animals
// @Tags Treatment
// @Accept json
// @Produce json
// @Param AnimalTypeCreate body models.CreateTreatment true "AnimalTypeCreate"
// @Success 200 {object} models.Treatment
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/treatment [post]
func (h *HandlerV1) CreateTreatment(c *gin.Context) {
	var (
		body        models.CreateTreatment
		jsonMarshal protojson.MarshalOptions
	)
	jsonMarshal.UseProtoNames = true

	err := c.ShouldBindJSON(&body)

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "CreateTreatment") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.treatment.CreateMedicineTreatment(ctx, &entity.Treatment{
		AnimalID:   body.AnimalID,
		MedicineID: body.MedicineID,
	})
	if e.HandleError(c, err, h.log, http.StatusInternalServerError, "CreateTreatment") {
		return
	}

	c.JSON(http.StatusOK, &models.Treatment{
		ID:            res.ID,
		AnimalID:      res.AnimalID,
		MedicineID:    res.MedicineID,
		TreatmentTime: e.Format(res.TreatmentTime),
	})
}

// GetTreatment ...
// @Summary Get Treatment
// @Description GetTreatment - Api for Get Animals
// @Tags Treatment
// @Accept json
// @Produce json
// @Param id query string true "id"
// @Success 200 {object} models.Treatment
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/treatment/get [get]
func (h *HandlerV1) GetTreatment(c *gin.Context) {
	id := c.Query("id")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.treatment.GetMedicineTreatment(ctx, &entity.FieldValueReq{
		Field: "id",
		Value: id,
	})

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "GetTreatment") {
		return
	}

	c.JSON(http.StatusOK, &models.Treatment{
		ID:            res.ID,
		AnimalID:      res.AnimalID,
		MedicineID:    res.MedicineID,
		TreatmentTime: e.Format(res.TreatmentTime),
	})
}

// ListTreatment ...
// @Summary List Treatment
// @Description ListTreatment - Api for List Treatment
// @Tags Treatment
// @Accept json
// @Produce json
// @Param ListReq query models.ListTreatmentReq false "ListReq"
// @Success 200 {object} models.ListTreatment
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/treatment [get]
func (h *HandlerV1) ListTreatment(c *gin.Context) {
	limit := c.Query("limit")
	page := c.Query("page")

	pageInt, limitInt, err := e.ParseQueryParams(page, limit)
	if e.HandleError(c, err, h.log, http.StatusBadRequest, "ListTreatment") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.treatment.ListMedicineTreatment(ctx, &entity.ListTreatmentReq{
		Page:  pageInt,
		Limit: limitInt,
	})

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "ListTreatment") {
		return
	}

	var treatments models.ListTreatment
	for _, treatRes := range res.Treatment {
		var treatment models.Treatment
		treatment.ID = treatRes.ID
		treatment.AnimalID = treatRes.AnimalID
		treatment.MedicineID = treatRes.MedicineID
		treatment.TreatmentTime = e.Format(treatRes.TreatmentTime)
		treatments.Treatment = append(treatments.Treatment, treatment)
	}
	treatments.Count = res.Count
	c.JSON(http.StatusOK, treatments)
}

// ListTreatmentByAnimalID ...
// @Summary List Treatment
// @Description ListTreatmentByAnimalID - Api for List Treatment
// @Tags Treatment
// @Accept json
// @Produce json
// @Param ListReq query models.ListTreatmentReq false "ListReq"
// @Param animal_id query string true "animal_id"
// @Success 200 {object} models.ListTreatment
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/treatment/animal-id [get]
func (h *HandlerV1) ListTreatmentByAnimalID(c *gin.Context) {
	limit := c.Query("limit")
	page := c.Query("page")
	animalID := c.Query("animal_id")

	pageInt, limitInt, err := e.ParseQueryParams(page, limit)
	if e.HandleError(c, err, h.log, http.StatusBadRequest, "ListTreatmentByAnimalID") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.treatment.ListMedicineTreatment(ctx, &entity.ListTreatmentReq{
		Page:  pageInt,
		Limit: limitInt,
		Field: "animal_id",
		Value: animalID,
	})

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "ListTreatmentByAnimalID") {
		return
	}

	var treatments models.ListTreatment
	for _, treatRes := range res.Treatment {
		var treatment models.Treatment
		treatment.ID = treatRes.ID
		treatment.AnimalID = treatRes.AnimalID
		treatment.MedicineID = treatRes.MedicineID
		treatment.TreatmentTime = e.Format(treatRes.TreatmentTime)
		treatments.Treatment = append(treatments.Treatment, treatment)
	}
	treatments.Count = res.Count
	c.JSON(http.StatusOK, treatments)
}

// ListTreatmentByMedicineID ...
// @Summary List Treatment
// @Description ListTreatmentByMedicineID - Api for List Treatment
// @Tags Treatment
// @Accept json
// @Produce json
// @Param ListReq query models.ListTreatmentReq false "ListReq"
// @Param medicine_id query string true "medicine_id"
// @Success 200 {object} models.ListTreatment
// @Failure 400 {object} models.ResponseError
// @Failure 500 {object} models.ResponseError
// @Router /v1/treatment/medicine-id [get]
func (h *HandlerV1) ListTreatmentByMedicineID(c *gin.Context) {
	limit := c.Query("limit")
	page := c.Query("page")
	medicineID := c.Query("medicine_id")

	pageInt, limitInt, err := e.ParseQueryParams(page, limit)
	if e.HandleError(c, err, h.log, http.StatusBadRequest, "ListTreatmentByMedicineID") {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.Context.Timeout))
	defer cancel()

	res, err := h.treatment.ListMedicineTreatment(ctx, &entity.ListTreatmentReq{
		Page:  pageInt,
		Limit: limitInt,
		Field: "medicine_id",
		Value: medicineID,
	})

	if e.HandleError(c, err, h.log, http.StatusBadRequest, "ListTreatmentByMedicineID") {
		return
	}

	var treatments models.ListTreatment
	for _, treatRes := range res.Treatment {
		var treatment models.Treatment
		treatment.ID = treatRes.ID
		treatment.AnimalID = treatRes.AnimalID
		treatment.MedicineID = treatRes.MedicineID
		treatment.TreatmentTime = e.Format(treatRes.TreatmentTime)
		treatments.Treatment = append(treatments.Treatment, treatment)
	}
	treatments.Count = res.Count
	c.JSON(http.StatusOK, treatments)
}
