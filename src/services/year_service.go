package services

import (
	"context"
	"github.com/alireza-fa/golang-car-shop/api/dto"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/data/db"
	"github.com/alireza-fa/golang-car-shop/data/models"
	"github.com/alireza-fa/golang-car-shop/pkg/logging"
)

type PersianYearService struct {
	base *BaseService[models.PersianYear, dto.CreatePersianYearRequest, dto.UpdatePersianYearRequest, dto.PersianYearResponse]
}

func NewPersianYearService(cfg *config.Config) *PersianYearService {
	return &PersianYearService{
		base: &BaseService[models.PersianYear, dto.CreatePersianYearRequest, dto.UpdatePersianYearRequest, dto.PersianYearResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// Create year
func (s *PersianYearService) Create(ctx context.Context, req *dto.CreatePersianYearRequest) (*dto.PersianYearResponse, error) {
	return s.base.Create(ctx, req)
}

// Update year
func (s *PersianYearService) Update(ctx context.Context, req *dto.UpdatePersianYearRequest, id int) (*dto.PersianYearResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete year
func (s *PersianYearService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// GetById year
func (s *PersianYearService) GetById(ctx context.Context, id int) (*dto.PersianYearResponse, error) {
	return s.base.GetById(ctx, id)
}

// GetByFilter year
func (s *PersianYearService) GetByFilter(ctx context.Context, req *dto.PaginationInputWithFilter) (*dto.PagedList[dto.PersianYearResponse], error) {
	return s.base.GetByFilter(ctx, req)
}
