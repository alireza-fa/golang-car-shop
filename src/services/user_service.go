package services

import (
	"github.com/alireza-fa/golang-car-shop/api/dto"
	"github.com/alireza-fa/golang-car-shop/common"
	"github.com/alireza-fa/golang-car-shop/config"
	"github.com/alireza-fa/golang-car-shop/data/db"
	"github.com/alireza-fa/golang-car-shop/pkg/logging"
	"gorm.io/gorm"
)

type UserService struct {
	logger     logging.Logger
	cfg        *config.Config
	otpService *OtpService
	database   *gorm.DB
}

func NewUserService(cfg *config.Config) *UserService {
	logger := logging.NewLogger(cfg)
	database := db.GetDb()
	return &UserService{
		logger:     logger,
		cfg:        cfg,
		otpService: NewOtpService(cfg),
		database:   database,
	}
}

func (s *UserService) SendOtp(req *dto.GetOtpRequest) error {
	otp := common.GenerateOtp()
	err := s.otpService.SetOtp(req.MobileNumber, otp)
	if err != nil {
		return err
	}
	return nil
}
