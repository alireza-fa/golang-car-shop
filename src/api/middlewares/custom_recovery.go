package middlewares

import (
	"github.com/alireza-fa/golang-car-shop/api/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler(c *gin.Context, err any) {
	if err, ok := err.(error); ok {
		httpResponse := helper.GenerateBaseResponseWithError(nil, false, -1, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
		return
	}
	httpResponse := helper.GenerateBaseResponseWithAnyError(nil, false, -1, err)
	c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
}
