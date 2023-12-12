package handlers

import (
	"github.com/alireza-fa/golang-car-shop/api/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HandlerGet godoc
// @Summary Health Check
// @Description Health Check
// #Tags health
// @Accept json
// @Produce json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/health/ [get]
func (h *HealthHandler) HandlerGet(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("working!!!", true, 0))
}
