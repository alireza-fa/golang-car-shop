package dto

type CreateCarTypeRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=15"`
}

type UpdateCarTypeRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=15"`
}

type CarTypeResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateGearboxRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=15"`
}

type UpdateGearboxRequest struct {
	Name string `json:"name" binding:"required,alpha,min=3,max=15"`
}

type GearboxResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CreateCarModelRequest struct {
	Name      string `json:"name" binding:"required,min=3,max=15"`
	CompanyId int    `json:"companyId"`
	CarTypeId int    `json:"carTypeId"`
	GearboxId int    `json:"gearboxId"`
}

type UpdateCarModelRequest struct {
	Name      string `json:"name,omitempty"`
	CompanyId int    `json:"companyId,omitempty"`
	CarTypeId int    `json:"carTypeId,omitempty"`
	GearboxId int    `json:"gearboxId,omitempty"`
}

type CarModelResponse struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	CarType CarTypeResponse `json:"carType"`
	Company CompanyResponse `json:"company"`
	Gearbox GearboxResponse `json:"gearbox"`
}
