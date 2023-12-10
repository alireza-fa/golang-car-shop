package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) HandlerGet(c *gin.Context) {
	c.JSON(http.StatusOK, "working get")
	return
}

func (h *HealthHandler) HandlerPost(c *gin.Context) {
	c.JSON(http.StatusOK, "working post")
	return
}

func (h *HealthHandler) HandlerById(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, fmt.Sprintf("%s working", id))
	return
}
