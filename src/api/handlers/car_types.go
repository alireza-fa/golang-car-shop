package handlers

import (
	_ "github.com/alireza-fa/golang-car-shop/api/dto"
	_ "github.com/alireza-fa/golang-car-shop/api/helper"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/services"
	"github.com/gin-gonic/gin"
)

type CarTypeHandler struct {
	service *services.CarTypeService
}

func NewCarTypeHandler(cfg *config.Config) *CarTypeHandler {
	return &CarTypeHandler{
		service: services.NewCarTypeService(cfg),
	}
}

// Create CarType godoc
// @Summary Create a CarType
// @Description Create a CarType
// @Tags CarTypes
// @Accept json
// @produces json
// @Param Request body dto.CreateCarTypeRequest true "Create a CarType"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarTypeResponse} "CarType response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-types/ [post]
// @Security AuthBearer
func (h *CarTypeHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// Update CarType godoc
// @Summary Update a CarType
// @Description Update a CarType
// @Tags CarTypes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarTypeRequest true "Update a CarType"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarTypeResponse} "CarType response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-types/{id} [put]
// @Security AuthBearer
func (h *CarTypeHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// Delete CarType godoc
// @Summary Delete a CarType
// @Description Delete a CarType
// @Tags CarTypes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-types/{id} [delete]
// @Security AuthBearer
func (h *CarTypeHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

// GetById CarType godoc
// @Summary Get a CarType
// @Description Get a CarType
// @Tags CarTypes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarTypeResponse} "CarType response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-types/{id} [get]
// @Security AuthBearer
func (h *CarTypeHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetByFilter CarTypes godoc
// @Summary Get CarTypes
// @Description Get CarTypes
// @Tags CarTypes
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarTypeResponse]} "CarType response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-types/get-by-filter [post]
// @Security AuthBearer
func (h *CarTypeHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
