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

type CountryHandler struct {
	service *services.CountryService
}

func NewCountryHandler(cfg *config.Config) *CountryHandler {
	return &CountryHandler{
		service: services.NewCountryService(cfg),
	}
}

// Create godoc
// @Summary Create country
// @Description Create country
// @Tags Countries
// @Accept json
// @Produce json
// @Param Request body dto.CreateUpdateCountryRequest true "country"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CountryResponse} "create a country"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/countries [post]
// @Security AuthBearer
func (h *CountryHandler) Create(c *gin.Context) {
	req := dto.CreateUpdateCountryRequest{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}

	res, err := h.service.Create(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}

	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(res, true, 0))
}

// Update godoc
// @Summary Update country
// @Description Update country
// @Tags Countries
// @Accept json
// @Produce json
// @Param Request body dto.CreateUpdateCountryRequest true "country request"
// @Param id path int true "country id"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.CountryResponse} "update a country"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/countries/{id} [patch]
// @Security AuthBearer
func (h *CountryHandler) Update(c *gin.Context) {
	countryId, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil && countryId == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	req := dto.CreateUpdateCountryRequest{}
	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}

	res, err := h.service.Update(c, &req, countryId)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))
}

// Delete godoc
// @Summary Delete country
// @Description Delete country
// @Tags Countries
// @Accept json
// @Produce json
// @Param id path int true "country id"
// @Success 204 "delete a country"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/countries/{id} [delete]
// @Security AuthBearer
func (h *CountryHandler) Delete(c *gin.Context) {
	countryId, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil && countryId == 0 {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	err = h.service.Delete(c, countryId)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}

	c.JSON(http.StatusNoContent, helper.GenerateBaseResponse(nil, true, 0))
}

// GetById godoc
// @Summary GetById country
// @Description GetById country
// @Tags Countries
// @Accept json
// @Produce json
// @Param id path int true "country id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.CountryResponse} "get country by id"
// @Success 204 "delete a country"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/countries/{id} [get]
// @Security AuthBearer
func (h *CountryHandler) GetById(c *gin.Context) {
	countryId, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	res, err := h.service.GetById(c, countryId)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))
}

// GetByFilter godoc
// @Summary GetByFilter country
// @Description GetByFilter country
// @Tags Countries
// @Accept json
// @Produce json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.CountryResponse]} "get country by id"
// @Success 400 {object} helper.BaseHttpResponse "Bad request"
// @Router /v1/countries/get-by-filter [post]
// @Security AuthBearer
func (h *CountryHandler) GetByFilter(c *gin.Context) {
	req := dto.PaginationInputWithFilter{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	res, err := h.service.GetByFilter(c, &req)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(res, true, 0))
}
