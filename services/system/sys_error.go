package system

import (
	"com.dreamfsk/blog/repository/models"
	"context"
)

// CreateSysError 创建错误日志记录
func (s *service) CreateSysError(ctx context.Context, sysError *models.SysError) (err error) {
	if s.db == nil {
		return nil
	}
	err = s.db.Create(sysError).Error
	return err
}

// DeleteSysError 删除错误日志记录
func (s *service) DeleteSysError(ctx context.Context, ID string) (err error) {
	err = s.db.Delete(&models.SysError{}, "id = ?", ID).Error
	return err
}

// DeleteSysErrorByIds 批量删除错误日志记录
func (s *service) DeleteSysErrorByIds(ctx context.Context, IDs []string) (err error) {
	err = s.db.Delete(&[]models.SysError{}, "id in ?", IDs).Error
	return err
}

// UpdateSysError 更新错误日志记录
func (s *service) UpdateSysError(ctx context.Context, sysError *models.SysError) (err error) {
	err = s.db.Model(&models.SysError{}).Where("id = ?", sysError.ID).Updates(&sysError).Error
	return err
}

// GetSysError 根据ID获取错误日志记录
func (s *service) GetSysError(ctx context.Context, ID string) (sysError models.SysError, err error) {
	err = s.db.Where("id = ?", ID).First(&sysError).Error
	return
}

// GetSysErrorInfoList 分页获取错误日志记录
func (s *service) GetSysErrorInfoList(ctx context.Context, info *SysErrorSearch) (list []models.SysError, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := s.db.Model(&models.SysError{}).Order("created_at desc")
	var sysErrors []models.SysError
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.CreatedAtRange) == 2 {
		db = db.Where("created_at BETWEEN ? AND ?", info.CreatedAtRange[0], info.CreatedAtRange[1])
	}

	if info.Form != nil && *info.Form != "" {
		db = db.Where("form = ?", *info.Form)
	}
	if info.Info != nil && *info.Info != "" {
		db = db.Where("info LIKE ?", "%"+*info.Info+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&sysErrors).Error
	return sysErrors, total, err
}

// GetSysErrorSolution 异步处理错误
func (s *service) GetSysErrorSolution(ctx context.Context, ID string) (err error) {
	// 立即更新为处理中
	err = s.db.Model(&models.SysError{}).Where("id = ?", ID).Update("status", "处理中").Error
	if err != nil {
		return err
	}
	// 异步协程在一分钟后更新为处理完成
	// todo 接入大模型
	return nil
}
