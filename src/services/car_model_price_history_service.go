package services

import (
	"context"
	"github.com/alireza-fa/golang-car-shop/api/dto"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/data/db"
	"github.com/alireza-fa/golang-car-shop/data/models"
	"github.com/alireza-fa/golang-car-shop/pkg/logging"
)

type CarModelPriceHistoryService struct {
	base *BaseService[models.CarModelPriceHistory, dto.CreateCarModelPriceHistoryRequest, dto.UpdateCarModelPriceHistoryRequest, dto.CarModelPriceHistoryResponse]
}

func NewCarModelPriceHistoryService(cfg *config.Config) *CarModelPriceHistoryService {
	return &CarModelPriceHistoryService{
		base: &BaseService[models.CarModelPriceHistory, dto.CreateCarModelPriceHistoryRequest, dto.UpdateCarModelPriceHistoryRequest, dto.CarModelPriceHistoryResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// Create carModelPriceHistory
func (s *CarModelPriceHistoryService) Create(ctx context.Context, req *dto.CreateCarModelPriceHistoryRequest) (*dto.CarModelPriceHistoryResponse, error) {
	return s.base.Create(ctx, req)
}

// Update carModelPriceHistory
func (s *CarModelPriceHistoryService) Update(ctx context.Context, req *dto.UpdateCarModelPriceHistoryRequest, id int) (*dto.CarModelPriceHistoryResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete carModelPriceHistory
func (s *CarModelPriceHistoryService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// GetById carModelPriceHistory
func (s *CarModelPriceHistoryService) GetById(ctx context.Context, id int) (*dto.CarModelPriceHistoryResponse, error) {
	return s.base.GetById(ctx, id)
}

// GetByFilter carModelPriceHistory
func (s *CarModelPriceHistoryService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelPriceHistoryResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
