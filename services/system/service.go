package system

import (
	"com.dreamfsk/blog/repository/models"
	"context"
	"gorm.io/gorm"
)

type Service interface {
	i()
	CreateSysError(ctx context.Context, sysError *models.SysError) (err error)
	DeleteSysError(ctx context.Context, ID string) (err error)
	DeleteSysErrorByIds(ctx context.Context, IDs []string) (err error)
	UpdateSysError(ctx context.Context, sysError *models.SysError) (err error)
	GetSysError(ctx context.Context, ID string) (sysError models.SysError, err error)
	GetSysErrorInfoList(ctx context.Context, info *SysErrorSearch) (list []models.SysError, total int64, err error)
	GetSysErrorSolution(ctx context.Context, ID string) (err error)
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{db: db}
}
func (s *service) i() {}
