package services

import (
	"context"
	"github.com/alireza-fa/golang-car-shop/api/dto"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/data/db"
	"github.com/alireza-fa/golang-car-shop/data/models"
	"github.com/alireza-fa/golang-car-shop/pkg/logging"
)

type PropertyService struct {
	base *BaseService[models.Property, dto.CreatePropertyRequest, dto.UpdatePropertyRequest, dto.PropertyResponse]
}

func NewPropertyService(cfg *config.Config) *PropertyService {
	return &PropertyService{
		base: &BaseService[models.Property, dto.CreatePropertyRequest, dto.UpdatePropertyRequest, dto.PropertyResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Category"},
				{string: "Category.Properties"},
			},
		},
	}
}

// Create property
func (s *PropertyService) Create(ctx context.Context, req *dto.CreatePropertyRequest) (*dto.PropertyResponse, error) {
	return s.base.Create(ctx, req)
}

// Update property
func (s *PropertyService) Update(ctx context.Context, req *dto.UpdatePropertyRequest, id int) (*dto.PropertyResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete property
func (s *PropertyService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// GetById property
func (s *PropertyService) GetById(ctx context.Context, id int) (*dto.PropertyResponse, error) {
	return s.base.GetById(ctx, id)
}

// GetByFilter property
func (s *PropertyService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.PropertyResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
