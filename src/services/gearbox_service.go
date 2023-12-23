package services

import (
	"context"
	"github.com/alireza-fa/golang-car-shop/api/dto"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/data/db"
	"github.com/alireza-fa/golang-car-shop/data/models"
	"github.com/alireza-fa/golang-car-shop/pkg/logging"
)

type GearboxService struct {
	base *BaseService[models.Gearbox, dto.CreateGearboxRequest, dto.UpdateGearboxRequest, dto.GearboxResponse]
}

func NewGearboxService(cfg *config.Config) *GearboxService {
	return &GearboxService{
		base: &BaseService[models.Gearbox, dto.CreateGearboxRequest, dto.UpdateGearboxRequest, dto.GearboxResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// Create gearbox
func (s *GearboxService) Create(ctx context.Context, req *dto.CreateGearboxRequest) (*dto.GearboxResponse, error) {
	return s.base.Create(ctx, req)
}

// Update gearbox
func (s *GearboxService) Update(ctx context.Context, req *dto.UpdateGearboxRequest, id int) (*dto.GearboxResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete gearbox
func (s *GearboxService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// GetById gearbox
func (s *GearboxService) GetById(ctx context.Context, id int) (*dto.GearboxResponse, error) {
	return s.base.GetById(ctx, id)
}

// GetByFilter gearbox
func (s *GearboxService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.GearboxResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
