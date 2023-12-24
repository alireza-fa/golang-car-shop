package handlers

import (
	_ "github.com/alireza-fa/golang-car-shop/api/dto"
	_ "github.com/alireza-fa/golang-car-shop/api/helper"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/services"
	"github.com/gin-gonic/gin"
)

type CarModelYearHandler struct {
	service *services.CarModelYearService
}

func NewCarModelYearHandler(cfg *config.Config) *CarModelYearHandler {
	return &CarModelYearHandler{
		service: services.NewCarModelYearService(cfg),
	}
}

// Create CarModelYear godoc
// @Summary Create a CarModelYear
// @Description Create a CarModelYear
// @Tags CarModelYears
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelYearRequest true "Create a CarModelYear"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelYearResponse} "CarModelYear response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-years/ [post]
// @Security AuthBearer
func (h *CarModelYearHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// Update CarModelYear godoc
// @Summary Update a CarModelYear
// @Description Update a CarModelYear
// @Tags CarModelYears
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelYearRequest true "Update a CarModelYear"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelYearResponse} "CarModelYear response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-years/{id} [put]
// @Security AuthBearer
func (h *CarModelYearHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// Delete CarModelYear godoc
// @Summary Delete a CarModelYear
// @Description Delete a CarModelYear
// @Tags CarModelYears
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-years/{id} [delete]
// @Security AuthBearer
func (h *CarModelYearHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

// GetById CarModelYear godoc
// @Summary Get a CarModelYear
// @Description Get a CarModelYear
// @Tags CarModelYears
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelYearResponse} "CarModelYear response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-years/{id} [get]
// @Security AuthBearer
func (h *CarModelYearHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetByFilter CarModelYears godoc
// @Summary Get CarModelYears
// @Description Get CarModelYears
// @Tags CarModelYears
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelYearResponse]} "CarModelYear response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-years/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelYearHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
