package services

import (
	"context"
	"github.com/alireza-fa/golang-car-shop/api/dto"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/data/db"
	"github.com/alireza-fa/golang-car-shop/data/models"
	"github.com/alireza-fa/golang-car-shop/pkg/logging"
)

type CompanyService struct {
	base *BaseService[models.Company, dto.CreateCompanyRequest, dto.UpdateCompanyRequest, dto.CompanyResponse]
}

func NewCompanyService(cfg *config.Config) *CompanyService {
	return &CompanyService{
		base: &BaseService[models.Company, dto.CreateCompanyRequest, dto.UpdateCompanyRequest, dto.CompanyResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{{string: "Country"}},
		},
	}
}

// Create company
func (s *CompanyService) Create(ctx context.Context, req *dto.CreateCompanyRequest) (*dto.CompanyResponse, error) {
	return s.base.Create(ctx, req)
}

// Update company
func (s *CompanyService) Update(ctx context.Context, req *dto.UpdateCompanyRequest, id int) (*dto.CompanyResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete company
func (s *CompanyService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// GetById company
func (s *CompanyService) GetById(ctx context.Context, id int) (*dto.CompanyResponse, error) {
	return s.base.GetById(ctx, id)
}

// GetByFilter company
func (s *CompanyService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CompanyResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
