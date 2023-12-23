package handlers

import (
	_ "github.com/alireza-fa/golang-car-shop/api/dto"
	_ "github.com/alireza-fa/golang-car-shop/api/helper"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/services"
	"github.com/gin-gonic/gin"
)

type PropertyCategoryHandler struct {
	service *services.PropertyCategoryService
}

func NewPropertyCategoryHandler(cfg *config.Config) *PropertyCategoryHandler {
	return &PropertyCategoryHandler{service: services.NewPropertyCategoryService(cfg)}
}

// Create PropertyCategory godoc
// @Summary Create a PropertyCategory
// @Description Create a PropertyCategory
// @Tags PropertyCategories
// @Access json
// @Produce json
// @Param Request body dto.CreatePropertyCategoryRequest true "Create a PropertyCategory"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.PropertyCategoryResponse} "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Bad Request"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/property-categories [get]
// @Security AuthBearer
func (h *PropertyCategoryHandler) Create(c *gin.Context) {
	Create(c, h.service.Create)
}

// Update PropertyCategory godoc
// @Summary Update a PropertyCategory
// @Description Update a PropertyCategory
// @Tags PropertyCategories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Param Request body dto.UpdatePropertyCategoryRequest true "Update a PropertyCategory"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PropertyCategoryResponse} "PropertyCategory response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/property-categories/{id} [put]
// @Security AuthBearer
func (h *PropertyCategoryHandler) Update(c *gin.Context) {
	Update(c, h.service.Update)
}

// Delete PropertyCategory godoc
// @Summary Delete a PropertyCategory
// @Description Delete a PropertyCategory
// @Tags PropertyCategories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse "response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/property-categories/{id} [delete]
// @Security AuthBearer
func (h *PropertyCategoryHandler) Delete(c *gin.Context) {
	Delete(c, h.service.Delete)
}

// GetById PropertyCategory godoc
// @Summary Get a PropertyCategory
// @Description Get a PropertyCategory
// @Tags PropertyCategories
// @Accept json
// @produces json
// @Param id path int true "Id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PropertyCategoryResponse} "PropertyCategory response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 404 {object} helper.BaseHttpResponse "Not found"
// @Router /v1/property-categories/{id} [get]
// @Security AuthBearer
func (h *PropertyCategoryHandler) GetById(c *gin.Context) {
	GetById(c, h.service.GetById)
}

// GetByFilter PropertyCategories godoc
// @Summary Get PropertyCategories
// @Description Get PropertyCategories
// @Tags PropertyCategories
// @Accept json
// @produces json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.PropertyCategoryResponse]} "PropertyCategory response"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/property-categories/get-by-filter [post]
// @Security AuthBearer
func (h *PropertyCategoryHandler) GetByFilter(c *gin.Context) {
	GetByFilter(c, h.service.GetByFilter)
}
