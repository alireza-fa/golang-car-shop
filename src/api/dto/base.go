package dto

import "mime/multipart"

type CreateUpdateCountryRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=20"`
}

type CountryResponse struct {
	Id     int            `json:"id"`
	Name   string         `json:"name"`
	Cities []CityResponse `json:"cities"`
}

type CreateCityRequest struct {
	Name      string `json:"name" binding:"alpha,min=3,max=20"`
	CountryId int    `json:"countryId" binding:"required"`
}

type UpdateCityRequest struct {
	Name      string `json:"name" binding:"alpha,min=3,max=20"`
	CountryId int    `json:"countryId,omitempty"`
}

type CityResponse struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Country CountryResponse `json:"country,omitempty"`
}

// Files

type FileFormRequest struct {
	File *multipart.FileHeader `json:"file" form:"file" binding:"required" swaggerignore:"true"`
}

type UploadFileRequest struct {
	FileFormRequest
	Description string `json:"description" form:"description" binding:"required"`
}

type CreateFileRequest struct {
	Name        string `json:"name"`
	Directory   string `json:"directory"`
	Description string `json:"description"`
	MimeType    string `json:"mimeType"`
}

type UpdateFileRequest struct {
	Description string `json:"description"`
}

type FileResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Directory   string `json:"directory"`
	Description string `json:"description"`
	MimeType    string `json:"mimeType"`
}
