package handlers

import (
	_ "github.com/alireza-fa/golang-car-shop/api/dto"
	_ "github.com/alireza-fa/golang-car-shop/api/helper"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/services"
	"github.com/gin-gonic/gin"
)

type CarModelImageHandler struct {
	service *services.CarModelImageService
}

func NewCarModelImageHandler(cfg *config.Config) *CarModelImageHandler {
	return &CarModelImageHandler{
		service: services.NewCarModelImageService(cfg),
	}
}

// Create CarModelImage godoc
// @Summary Create a CarModelImage
// @Description Create a CarModelImage
// @Tags CarModelImages
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelImageRequest true "Create a CarModelImage"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelImageResponse} "CarModelImage response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-images/ [post]
// @Security AuthBearer
func (h *CarModelImageHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// Update CarModelImage godoc
// @Summary Update a CarModelImage
// @Description Update a CarModelImage
// @Tags CarModelImages
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelImageRequest true "Update a CarModelImage"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelImageResponse} "CarModelImage response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-images/{id} [put]
// @Security AuthBearer
func (h *CarModelImageHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// Delete CarModelImage godoc
// @Summary Delete a CarModelImage
// @Description Delete a CarModelImage
// @Tags CarModelImages
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-images/{id} [delete]
// @Security AuthBearer
func (h *CarModelImageHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

// GetById CarModelImage godoc
// @Summary Get a CarModelImage
// @Description Get a CarModelImage
// @Tags CarModelImages
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelImageResponse} "CarModelImage response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-images/{id} [get]
// @Security AuthBearer
func (h *CarModelImageHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetByFilter CarModelImages godoc
// @Summary Get CarModelImages
// @Description Get CarModelImages
// @Tags CarModelImages
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelImageResponse]} "CarModelImage response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-images/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelImageHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
