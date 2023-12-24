package handlers

import (
	_ "github.com/alireza-fa/golang-car-shop/api/dto"
	_ "github.com/alireza-fa/golang-car-shop/api/helper"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/services"
	"github.com/gin-gonic/gin"
)

type CarModelCommentHandler struct {
	service *services.CarModelCommentService
}

func NewCarModelCommentHandler(cfg *config.Config) *CarModelCommentHandler {
	return &CarModelCommentHandler{
		service: services.NewCarModelCommentService(cfg),
	}
}

// Create CarModelComment godoc
// @Summary Create a CarModelComment
// @Description Create a CarModelComment
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param Request body dto.CreateCarModelCommentRequest true "Create a CarModelComment"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CarModelCommentResponse} "CarModelComment response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-comments/ [post]
// @Security AuthBearer
func (h *CarModelCommentHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// Update CarModelComment godoc
// @Summary Update a CarModelComment
// @Description Update a CarModelComment
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdateCarModelCommentRequest true "Update a CarModelComment"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelCommentResponse} "CarModelComment response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-comments/{id} [put]
// @Security AuthBearer
func (h *CarModelCommentHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// Delete CarModelComment godoc
// @Summary Delete a CarModelComment
// @Description Delete a CarModelComment
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-comments/{id} [delete]
// @Security AuthBearer
func (h *CarModelCommentHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

// GetById CarModelComment godoc
// @Summary Get a CarModelComment
// @Description Get a CarModelComment
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CarModelCommentResponse} "CarModelComment response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/car-model-comments/{id} [get]
// @Security AuthBearer
func (h *CarModelCommentHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetByFilter CarModelComments godoc
// @Summary Get CarModelComments
// @Description Get CarModelComments
// @Tags CarModelComments
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CarModelCommentResponse]} "CarModelComment response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/car-model-comments/get-by-filter [post]
// @Security AuthBearer
func (h *CarModelCommentHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
