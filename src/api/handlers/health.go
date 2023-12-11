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

func (h *HealthHandler) HandlerGet(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("working!!!", true, 0))
}
