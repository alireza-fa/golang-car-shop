package services

import (
	"context"
	"github.com/alireza-fa/golang-car-shop/api/dto"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/data/db"
	"github.com/alireza-fa/golang-car-shop/data/models"
	"github.com/alireza-fa/golang-car-shop/pkg/logging"
)

//type CountryService struct {
//	database *gorm.DB
//	logger   logging.Logger
//}

type CountryService struct {
	base *BaseService[models.Country, dto.CreateUpdateCountryRequest, dto.CreateUpdateCountryRequest, dto.CountryResponse]
}

func NewCountryService(cfg *config.Config) *CountryService {
	return &CountryService{
		base: &BaseService[models.Country, dto.CreateUpdateCountryRequest, dto.CreateUpdateCountryRequest, dto.CountryResponse]{
			Database: db.GetDb(),
			Logger:   logging.NewLogger(cfg),
		},
	}
}

// Create country
func (s *CountryService) Create(ctx context.Context, req *dto.CreateUpdateCountryRequest) (*dto.CountryResponse, error) {
	return s.base.Create(ctx, req)
}

// Update country
func (s *CountryService) Update(ctx context.Context, req *dto.CreateUpdateCountryRequest, id int) (*dto.CountryResponse, error) {
	return s.base.Update(ctx, id, req)
}

// Delete country
func (s *CountryService) Delete(ctx context.Context, id int) error {
	return s.base.Delete(ctx, id)
}

// GetById country
func (s *CountryService) GetById(ctx context.Context, id int) (*dto.CountryResponse, error) {
	return s.base.GetById(ctx, id)
}

//// Create a country
//func (s *CountryService) Create(ctx context.Context, req *dto.CreateUpdateCountryRequest) (*dto.CountryResponse, error) {
//	country := models.Country{Name: req.Name}
//	userId := ctx.Value(constants.UserIdKey)
//	if userId == nil {
//		return nil, &service_errors.ServiceError{EndUserMessage: service_errors.PermissionDenied}
//	}
//
//	country.CreatedAt = time.Now().UTC()
//	country.CreatedBy = int(userId.(float64))
//
//	tx := s.database.WithContext(ctx).Begin()
//	err := tx.Create(&country).Error
//	if err != nil {
//		tx.Rollback()
//		s.logger.Error(logging.Postgres, logging.Insert, "Create", nil)
//		return nil, err
//	}
//	tx.Commit()
//
//	return s.GetById(ctx, country.Id)
//}
//
//// Update country
//func (s *CountryService) Update(ctx context.Context, req *dto.CreateUpdateCountryRequest, countryId int) (*dto.CountryResponse, error) {
//	userId := ctx.Value(constants.UserIdKey)
//	if userId == nil {
//		return nil, &service_errors.ServiceError{EndUserMessage: service_errors.PermissionDenied}
//	}
//
//	updateMap := map[string]interface{}{
//		"Name":        req.Name,
//		"modified_by": &sql.NullInt64{Int64: int64(userId.(float64)), Valid: true},
//		"modified_at": sql.NullTime{Time: time.Now().UTC(), Valid: true},
//	}
//
//	tx := s.database.WithContext(ctx).Begin()
//	if err := tx.
//		Model(&models.Country{}).
//		Where("id = ? AND deleted_by is null", countryId).
//		Updates(updateMap).
//		Error; err != nil {
//		tx.Rollback()
//		s.logger.Error(logging.Postgres, logging.Update, "Update", nil)
//		return nil, err
//	}
//	tx.Commit()
//
//	return s.GetById(ctx, countryId)
//}
//
//// Delete country
//func (s *CountryService) Delete(ctx context.Context, countryId int) error {
//	tx := s.database.WithContext(ctx).Begin()
//
//	userId := ctx.Value(constants.UserIdKey)
//	if userId == nil {
//		tx.Rollback()
//		return &service_errors.ServiceError{EndUserMessage: service_errors.PermissionDenied}
//	}
//
//	deleteCountry := map[string]interface{}{
//		"deleted_at": sql.NullTime{Time: time.Now().UTC(), Valid: true},
//		"deleted_by": &sql.NullInt64{Int64: int64(userId.(float64)), Valid: true},
//	}
//
//	if err := tx.
//		Model(&models.Country{}).
//		Where("id = ? AND deleted_by is null", countryId).
//		Updates(deleteCountry).
//		Error; err != nil {
//		tx.Rollback()
//		s.logger.Error(logging.Postgres, logging.Delete, "Delete", nil)
//		return err
//	}
//	tx.Commit()
//	return nil
//}
//
//// GetById country
//func (s *CountryService) GetById(ctx context.Context, countryId int) (*dto.CountryResponse, error) {
//	var country dto.CountryResponse
//
//	if err := s.database.
//		Model(&models.Country{}).
//		Where("id = ? and deleted_by is null", countryId).
//		First(&country).
//		Error; err != nil {
//		return nil, err
//	}
//	return &country, nil
//}
