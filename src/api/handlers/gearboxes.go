package handlers

import (
	_ "github.com/alireza-fa/golang-car-shop/api/dto"
	_ "github.com/alireza-fa/golang-car-shop/api/helper"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/services"
	"github.com/gin-gonic/gin"
)

type GearboxHandler struct {
	service *services.GearboxService
}

func NewGearboxHandler(cfg *config.Config) *GearboxHandler {
	return &GearboxHandler{
		service: services.NewGearboxService(cfg),
	}
}

// Create Gearbox godoc
// @Summary Create a Gearbox
// @Description Create a Gearbox
// @Tags Gearboxes
// @Accept json
// @produces json
// @Param Request body dto.CreateGearboxRequest true "Create a Gearbox"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.GearboxResponse} "Gearbox response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/gearboxes/ [post]
// @Security AuthBearer
func (h *GearboxHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// Update Gearbox godoc
// @Summary Update a Gearbox
// @Description Update a Gearbox
// @Tags Gearboxes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateGearboxRequest true "Update a Gearbox"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.GearboxResponse} "Gearbox response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/gearboxes/{id} [put]
// @Security AuthBearer
func (h *GearboxHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// Delete Gearbox godoc
// @Summary Delete a Gearbox
// @Description Delete a Gearbox
// @Tags Gearboxes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/gearboxes/{id} [delete]
// @Security AuthBearer
func (h *GearboxHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

// GetById Gearbox godoc
// @Summary Get a Gearbox
// @Description Get a Gearbox
// @Tags Gearboxes
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.GearboxResponse} "Gearbox response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/gearboxes/{id} [get]
// @Security AuthBearer
func (h *GearboxHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetByFilter Gearboxes godoc
// @Summary Get Gearboxes
// @Description Get Gearboxes
// @Tags Gearboxes
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.GearboxResponse]} "Gearbox response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/gearboxes/get-by-filter [post]
// @Security AuthBearer
func (h *GearboxHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
