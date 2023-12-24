package services

import (
	"context"
	"github.com/alireza-fa/golang-car-shop/api/dto"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/data/db"
	"github.com/alireza-fa/golang-car-shop/data/models"
	"github.com/alireza-fa/golang-car-shop/pkg/logging"
)

type CarModelYearService struct {
	base *BaseService[models.CarModelYear, dto.CreateCarModelYearRequest, dto.UpdateCarModelYearRequest, dto.CarModelYearResponse]
}

func NewCarModelYearService(cfg *config.Config) *CarModelYearService {
	return &CarModelYearService{
		base: &BaseService[models.CarModelYear, dto.CreateCarModelYearRequest, dto.UpdateCarModelYearRequest, dto.CarModelYearResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
			Preloads: []preload{
				{string: "PersianYear"},
			},
		},
	}
}

// Create carModelYear
func (s *CarModelYearService) Create(ctx context.Context, req *dto.CreateCarModelYearRequest) (*dto.CarModelYearResponse, error) {
	return s.base.Create(ctx, req)
}

// Update carModelYear
func (s *CarModelYearService) Update(ctx context.Context, req *dto.UpdateCarModelYearRequest, id int) (*dto.CarModelYearResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete carModelYear
func (s *CarModelYearService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// GetById carModelYear
func (s *CarModelYearService) GetById(ctx context.Context, id int) (*dto.CarModelYearResponse, error) {
	return s.base.GetById(ctx, id)
}

// GetByFilter carModelYear
func (s *CarModelYearService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.CarModelYearResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
