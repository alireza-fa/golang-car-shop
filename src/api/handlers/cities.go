package handlers

import (
	"github.com/alireza-fa/golang-car-shop/api/dto"
	"github.com/alireza-fa/golang-car-shop/api/helper"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CityHandler struct {
	service *services.CityService
}

func NewCityHandler(cfg *config.Config) *CityHandler {
	return &CityHandler{service: services.NewCityService(cfg)}
}

// Create godoc
// @Summary Create cities
// @Description Create cities
// @Tags Cities
// @Accept json
// @Produce json
// @param Request body dto.CreateCityRequest true "city"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CityResponse} "create a new city"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse
// @Router /v1/cities [post]
// @Security AuthBearer
func (h *CityHandler) Create(c *gin.Context) {
	req := dto.CreateCityRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	res, err := h.service.Create(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, helper.Success))
}

// Update godoc
// @Summary Update city
// @Description Update city
// @Tags Cities
// @Accept json
// @Produce json
// @Param Request body dto.UpdateCityRequest true "city request"
// @Param id path int true "city id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CityResponse} "update city"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/cities/{id} [patch]
// @Security AuthBearer
func (h *CityHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}
	req := dto.UpdateCityRequest{}
	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	res, err := h.service.Update(c, id, &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, helper.Success))
}

// Delete godoc
// @Summary Delete city
// @Description Delete city
// @Tags Cities
// @Accept json
// @Produce json
// @Param id path int true "city id"
// @Success 204 {object} helper.BaseHttpResponse "delete a city"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/cities/{id} [delete]
// @Security AuthBearer
func (h *CityHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	err = h.service.Delete(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}
	c.JSON(http.StatusNoContent, helper.GenerateBaseResponse(nil, true, helper.Success))
}

// GetById godoc
// @Summary GetById cities
// @Description GetById cities
// @Tags Cities
// @Accept json
// @Produce json
// @Param id path int true "city id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CountryResponse} "get city by id"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/cities/{id} [get]
// @Security AuthBearer
func (h *CityHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	res, err := h.service.GetById(c, id)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, helper.Success))
}

// GetByFilter godoc
// @Summary GetByFilter cities
// @Description GetByFilter cities
// @Tags Cities
// @Accept json
// @Produce json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CityResponse]} "get country by id"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/cities/get-by-filter [post]
// @Security AuthBearer
func (h *CityHandler) GetByFilter(c *gin.Context) {
	req := dto.PaginationInputWithFilter{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}

	res, err := h.service.GetByFilter(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, helper.Success))
}
