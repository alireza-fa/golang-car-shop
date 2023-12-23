package handlers

import (
	_ "github.com/alireza-fa/golang-car-shop/api/dto"
	_ "github.com/alireza-fa/golang-car-shop/api/helper"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/services"
	"github.com/gin-gonic/gin"
)

type PropertyHandler struct {
	service *services.PropertyService
}

func NewPropertyHandler(cfg *config.Config) *PropertyHandler {
	return &PropertyHandler{
		service: services.NewPropertyService(cfg),
	}
}

// Create Property godoc
// @Summary Create a Property
// @Description Create a Property
// @Tags Properties
// @Accept json
// @produces json
// @Param Request body dto.CreatePropertyRequest true "Create a Property"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.PropertyResponse} "Property response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/properties/ [post]
// @Security AuthBearer
func (h *PropertyHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// Update Property godoc
// @Summary Update a Property
// @Description Update a Property
// @Tags Properties
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdatePropertyRequest true "Update a Property"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PropertyResponse} "Property response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/properties/{id} [put]
// @Security AuthBearer
func (h *PropertyHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// Delete Property godoc
// @Summary Delete a Property
// @Description Delete a Property
// @Tags Properties
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/properties/{id} [delete]
// @Security AuthBearer
func (h *PropertyHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

// GetById Property godoc
// @Summary Get a Property
// @Description Get a Property
// @Tags Properties
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PropertyResponse} "Property response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/properties/{id} [get]
// @Security AuthBearer
func (h *PropertyHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetByFilter Properties godoc
// @Summary Get Properties
// @Description Get Properties
// @Tags Properties
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.PropertyResponse]} "Property response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/properties/get-by-filter [post]
// @Security AuthBearer
func (h *PropertyHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
