package handlers

import (
	"fmt"
	"github.com/alireza-fa/golang-car-shop/api/dto"
	"github.com/alireza-fa/golang-car-shop/api/helper"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/pkg/logging"
	"github.com/alireza-fa/golang-car-shop/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var logger = logging.NewLogger(config.GetConfig())

type FileHandler struct {
	service *services.FileService
}

func NewFileHandler(cfg *config.Config) *FileHandler {
	return &FileHandler{service: services.NewFileService(cfg)}
}

// Create file godoc
// @Summary Create a file
// @Description Create a file
// @Tags Files
// @Accept x-www-form-urlencoded
// @Produce json
// @Param file formData dto.UploadFileRequest true "Create a file"
// @Param file formData file true "Create a file"
// @Success 201 {object} helper.BaseHttpResponse{result=dto.FileResponse} "File response"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/files [post]
// @Security AuthBearer
func (h *FileHandler) Create(c *gin.Context) {
	upload := dto.UploadFileRequest{}
	err := c.ShouldBind(&upload)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helper.GenerateBaseResponseWithValidationError(nil, false, helper.ValidationError, err))
		return
	}
	req := dto.CreateFileRequest{}
	req.Description = upload.Description
	req.MimeType = upload.File.Header.Get("Content-Type")
	req.Directory = "uploads"
	req.Name, err = saveUploadFile(upload.File, req.Directory)
	if err != nil {
		c.AbortWithStatusJSON(helper.TranslateErrorToStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, helper.InternalError, err))
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
// @Summary Update file
// @Description Update file
// @Tags Files
// @Accept json
// @Produce json
// @Param Request body dto.UpdateFileRequest true "city request"
// @Param id path int true "file id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.FileResponse} "update file"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/files/{id} [patch]
// @Security AuthBearer
func (h *FileHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	req := dto.UpdateFileRequest{}
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
// @Summary Delete file
// @Description Delete file
// @Tags Files
// @Accept json
// @Produce json
// @Param id path int true "file id"
// @Success 204 {object} helper.BaseHttpResponse "delete a file"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/files/{id} [delete]
// @Security AuthBearer
func (h *FileHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	file, err := h.service.GetById(c, id)
	if err != nil {
		logger.Error(logging.IO, logging.RemoveFile, err.Error(), nil)
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, helper.NotFoundError))
		return
	}
	err = os.Remove(fmt.Sprintf("%s/%s", file.Directory, file.Name))
	if err != nil {
		logger.Error(logging.IO, logging.RemoveFile, err.Error(), nil)
		c.AbortWithStatusJSON(http.StatusNotFound,
			helper.GenerateBaseResponse(nil, false, helper.NotFoundError))
		return
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
// @Summary GetById file
// @Description GetById file
// @Tags Files
// @Accept json
// @Produce json
// @Param id path int true "city id"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.FileResponse} "get file by id"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/files/{id} [get]
// @Security AuthBearer
func (h *FileHandler) GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
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
// @Summary GetByFilter files
// @Description GetByFilter files
// @Tags Files
// @Accept json
// @Produce json
// @Param Request body dto.PaginationInputWithFilter true "Request"
// @Success 200 {object} helper.BaseHttpResponse{result=dto.PagedList[dto.FileResponse]} "get files by filter"
// @Failure 400 {object} helper.BaseHttpResponse "Bad request"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/files/get-by-filter [post]
// @Security AuthBearer
func (h *FileHandler) GetByFilter(c *gin.Context) {
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

func saveUploadFile(file *multipart.FileHeader, directory string) (string, error) {
	// test.txt -> 6564464586466.txt (uuid).(type)
	randFileName := uuid.New()
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		return "", err
	}
	fileName := file.Filename
	fileNameArr := strings.Split(fileName, ".")
	fileExt := fileNameArr[len(fileNameArr)-1]
	fileName = fmt.Sprintf("%s.%s", randFileName, fileExt)
	dst := fmt.Sprintf("%s/%s", directory, fileName)

	src, err := file.Open()
	if err != nil {
		return "", nil
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	if err != nil {
		return "", err
	}
	return fileName, nil
}
