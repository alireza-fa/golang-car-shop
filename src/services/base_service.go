package services

import (
	"context"
	"database/sql"
	"github.com/alireza-fa/golang-car-shop/common"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/constants"
	"github.com/alireza-fa/golang-car-shop/data/db"
	"github.com/alireza-fa/golang-car-shop/pkg/logging"
	"github.com/alireza-fa/golang-car-shop/pkg/service_errors"
	"gorm.io/gorm"
	"time"
)

type BaseService[T any, Tc any, Tu any, Tr any] struct {
	Database *gorm.DB
	Logger   logging.Logger
}

func NewBaseService[T any, Tc any, Tu any, Tr any](cfg *config.Config) *BaseService[T, Tc, Tu, Tr] {
	return &BaseService[T, Tc, Tu, Tr]{
		Database: db.GetDb(),
		Logger:   logging.NewLogger(cfg),
	}
}

func (s *BaseService[T, Tc, Tu, Tr]) Create(ctx context.Context, req *Tc) (*Tr, error) {
	model, _ := common.TypeConverter[T](req)
	tx := s.Database.WithContext(ctx).Begin()
	err := tx.
		Create(model).Error
	if err != nil {
		tx.Rollback()
		s.Logger.Error(logging.Postgres, logging.Insert, err.Error(), nil)
		return nil, err
	}
	tx.Commit()
	return common.TypeConverter[Tr](model)
}

func (s *BaseService[T, Tc, Tu, Tr]) Update(ctx context.Context, id int, req *Tu) (*Tr, error) {
	updateMap, _ := common.TypeConverter[map[string]interface{}](req)
	(*updateMap)["modified_by"] = &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true}
	(*updateMap)["modified_at"] = sql.NullTime{Time: time.Now().UTC(), Valid: true}
	model := new(T)
	tx := s.Database.WithContext(ctx).Begin()
	if err := tx.
		Model(model).
		Where("id = ? and deleted_by is null", id).
		Updates(*updateMap).Error; err != nil {
		tx.Rollback()
		s.Logger.Error(logging.Postgres, logging.Update, err.Error(), nil)
		return nil, err
	}
	tx.Commit()
	return s.GetById(ctx, id)
}

func (s *BaseService[T, Tc, Tu, Tr]) Delete(ctx context.Context, id int) error {
	tx := s.Database.WithContext(ctx).Begin()

	model := new(T)

	deleteMap := map[string]interface{}{
		"deleted_by": &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true},
		"deleted_at": sql.NullTime{Time: time.Now().UTC(), Valid: true},
	}
	deleteMap["modified_by"] = &sql.NullInt64{Int64: int64(ctx.Value(constants.UserIdKey).(float64)), Valid: true}
	deleteMap["modified_at"] = sql.NullTime{Time: time.Now().UTC(), Valid: true}

	if ctx.Value(constants.UserIdKey) == nil {
		return &service_errors.ServiceError{EndUserMessage: service_errors.PermissionDenied}
	}
	if cnt := tx.
		Model(model).
		Where("id = ? and deleted_by is null", id).
		Updates(deleteMap).RowsAffected; cnt == 0 {
		tx.Rollback()
		s.Logger.Error(logging.Postgres, logging.Update, service_errors.RecordNotFound, nil)
		return &service_errors.ServiceError{EndUserMessage: service_errors.RecordNotFound}
	}
	tx.Commit()
	return nil
}

func (s *BaseService[T, Tc, Tu, Tr]) GetById(ctx context.Context, id int) (*Tr, error) {
	model := new(T)
	err := s.Database.
		Where("id = ? and deleted_by is null", id).
		First(model).
		Error
	if err != nil {
		return nil, err
	}
	return common.TypeConverter[Tr](model)
}
