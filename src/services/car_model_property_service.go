package services

import (
	"context"
	"github.com/alireza-fa/golang-car-shop/api/dto"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/data/db"
	"github.com/alireza-fa/golang-car-shop/data/models"
	"github.com/alireza-fa/golang-car-shop/pkg/logging"
)

type CarModelPropertyService struct {
	base *BaseService[models.CarModelProperty, dto.CreateCarModelPropertyRequest, dto.UpdateCarModelPropertyRequest, dto.CarModelPropertyResponse]
}

func NewCarModelPropertyService(cfg *config.Config) *CarModelPropertyService {
	return &CarModelPropertyService{
		base: &BaseService[models.CarModelProperty, dto.CreateCarModelPropertyRequest, dto.UpdateCarModelPropertyRequest, dto.CarModelPropertyResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "Property.Category"},
			},
		},
	}
}

// Create carModelProperty
func (s *CarModelPropertyService) Create(ctx context.Context, req *dto.CreateCarModelPropertyRequest) (*dto.CarModelPropertyResponse, error) {
	return s.base.Create(ctx, req)
}

// Update carModelProperty
func (s *CarModelPropertyService) Update(ctx context.Context, req *dto.UpdateCarModelPropertyRequest, id int) (*dto.CarModelPropertyResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete carModelProperty
func (s *CarModelPropertyService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// GetById carModelProperty
func (s *CarModelPropertyService) GetById(ctx context.Context, id int) (*dto.CarModelPropertyResponse, error) {
	return s.base.GetById(ctx, id)
}

// GetByFilter carModelProperty
func (s *CarModelPropertyService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelPropertyResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
