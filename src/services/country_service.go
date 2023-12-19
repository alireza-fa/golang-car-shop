package services

import (
	"context"
	"database/sql"
	"github.com/alireza-fa/golang-car-shop/api/dto"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/constants"
	"github.com/alireza-fa/golang-car-shop/data/db"
	"github.com/alireza-fa/golang-car-shop/data/models"
	"github.com/alireza-fa/golang-car-shop/pkg/logging"
	"github.com/alireza-fa/golang-car-shop/pkg/service_errors"
	"gorm.io/gorm"
	"time"
)

type CountryService struct {
	database *gorm.DB
	logger   logging.Logger
}

func NewCountryService(cfg *config.Config) *CountryService {
	return &CountryService{
		database: db.GetDb(),
		logger:   logging.NewLogger(cfg),
	}
}

// Create a country
func (s *CountryService) Create(ctx context.Context, req *dto.CreateUpdateCountryRequest) (*dto.CountryResponse, error) {
	country := models.Country{Name: req.Name}
	userId := ctx.Value(constants.UserIdKey)
	if userId == nil {
		return nil, &service_errors.ServiceError{EndUserMessage: service_errors.PermissionDenied}
	}

	country.CreatedAt = time.Now().UTC()
	country.CreatedBy = int(userId.(float64))

	tx := s.database.WithContext(ctx).Begin()
	err := tx.Create(country).Error
	if err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Insert, "Create", nil)
		return nil, err
	}
	tx.Commit()
	// TODO: get by id
	return nil, nil
}

// Update country
func (s *CountryService) Update(ctx context.Context, req *dto.CreateUpdateCountryRequest, countryId int) (*dto.CountryResponse, error) {
	userId := ctx.Value(constants.UserIdKey)
	if userId == nil {
		return nil, &service_errors.ServiceError{EndUserMessage: service_errors.PermissionDenied}
	}

	countryUpdate := map[string]interface{}{
		"modified_at": sql.NullTime{Valid: true, Time: time.Now().UTC()},
		"modified_by": &sql.NullInt64{Valid: true, Int64: int64(userId.(float64))},
		"name":        req.Name,
	}

	tx := s.database.Where(ctx).Begin()
	if err := tx.
		Model(&models.Country{}).
		Where("id = ? AND deleted_by is null", countryId).
		Updates(countryUpdate).
		Error; err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Update, "Update", nil)
		return nil, err
	}
	tx.Commit()
	// TODO: get by id
	return nil, nil
}

// Delete country
func (s *CountryService) Delete(ctx context.Context, countryId int) error {
	tx := s.database.WithContext(ctx).Begin()

	userId := ctx.Value(constants.UserIdKey)
	if userId == nil {
		tx.Rollback()
		return &service_errors.ServiceError{EndUserMessage: service_errors.PermissionDenied}
	}

	deleteCountry := map[string]interface{}{
		"deleted_at": sql.NullTime{Time: time.Now().UTC(), Valid: true},
		"deleted_by": &sql.NullInt64{Int64: int64(userId.(float64))},
	}

	if err := tx.
		Model(&models.Country{}).
		Where("id = ? AND deleted_by is null", countryId).
		Updates(deleteCountry).
		Error; err != nil {
		tx.Rollback()
		s.logger.Error(logging.Postgres, logging.Delete, "Delete", nil)
		return err
	}
	tx.Commit()
	return nil
}

// GetById country
func (s *CountryService) GetById(ctx context.Context, countryId int) (*dto.CountryResponse, error) {
	var country dto.CountryResponse

	if err := s.database.
		Model(&models.Country{}).
		Where("id = ? and deleted_by is null", countryId).
		First(&country).
		Error; err != nil {
		return nil, err
	}
	return &country, nil
}
