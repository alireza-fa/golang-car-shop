package handlers

import (
	"github.com/alireza-fa/golang-car-shop/api/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

type header struct {
	Browser string
	Id      string
}

type personData struct {
	FirstName   string `json:"first_name" binding:"required,alpha,min=3,max=10"`
	LastName    string `json:"last_name" binding:"required,alpha,min=3,max=10"`
	PhoneNumber string `json:"phone_number" binding:"required,mobile,min=11,max=11"`
	Password    string `json:"password" binding:"required,password"`
}

type TestHandler struct {
}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Users(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("users", true, 0))
}

// UserById godoc
// @Summary UserById
// @Description UserById
// @Tags Test
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/test/users/{id} [get]
func (h *TestHandler) UserById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{"detail": "User by Id", "id": id}, true, 0))
}

func (h *TestHandler) UserByUsername(c *gin.Context) {
	id := c.Param("id")
	username := c.Param("username")
	c.JSON(http.StatusOK, gin.H{
		"detail":   "User By Username",
		"username": username,
		"id":       id,
	})
}

func (h *TestHandler) CreateUser(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"detail": "Created User",
	})
}

func (h *TestHandler) HeaderBinder1(c *gin.Context) {
	id := c.GetHeader("id")
	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"result": "BinderHeader1",
	})
}

func (h *TestHandler) HeaderBinder2(c *gin.Context) {
	header := header{}
	_ = c.BindHeader(&header)
	c.JSON(http.StatusOK, gin.H{
		"header": header,
		"result": "BinderHeader2",
	})
}

func (h *TestHandler) QueryBinder1(c *gin.Context) {
	id := c.Query("id")
	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"name":   name,
		"result": "BinderQuery1",
	})
}

func (h *TestHandler) QueryBinder2(c *gin.Context) {
	ids := c.QueryArray("id")
	name := c.Query("name")

	c.JSON(http.StatusOK, gin.H{
		"ids":    ids,
		"name":   name,
		"result": "BinderQuery2",
	})
}

func (h *TestHandler) UriBinder(c *gin.Context) {
	id := c.Param("id")
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"id":     id,
		"name":   name,
		"result": "UriBinder",
	})
}

// BodyBinder godoc
// @Summary BodyBinder
// @Description BodyBinder
// @Tags Test
// @Accept json
// @Produce json
// @Param person body personData true "person data"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 401 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/test/binder/body [post]
// @Security AuthBearer
func (h *TestHandler) BodyBinder(c *gin.Context) {
	p := personData{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{"person": p, "result": "Body binder with ShouldBindJson"}, true, 0))
}

func (h *TestHandler) FormBinder(c *gin.Context) {
	p := personData{}
	err := c.ShouldBind(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"person": p,
		"result": "formBinder",
	})
}

func (h *TestHandler) FileBinder(c *gin.Context) {
	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, "file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"filename": file.Filename,
		"result":   "FileBinder",
	})
}
