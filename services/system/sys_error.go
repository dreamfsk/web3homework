package system

import (
	"com.dreamfsk/blog/config"
	"com.dreamfsk/blog/repository/models"
)

type SysErrorService struct{}

// CreateSysError 创建错误日志记录
func (sysErrorService *SysErrorService) CreateSysError(sysError *models.SysError) (err error) {
	if config.DB == nil {
		return nil
	}
	err = config.DB.Create(sysError).Error
	return err
}

// DeleteSysError 删除错误日志记录
func (sysErrorService *SysErrorService) DeleteSysError(ID string) (err error) {
	err = config.DB.Delete(&models.SysError{}, "id = ?", ID).Error
	return err
}

// DeleteSysErrorByIds 批量删除错误日志记录
func (sysErrorService *SysErrorService) DeleteSysErrorByIds(IDs []string) (err error) {
	err = config.DB.Delete(&[]models.SysError{}, "id in ?", IDs).Error
	return err
}

// UpdateSysError 更新错误日志记录
func (sysErrorService *SysErrorService) UpdateSysError(sysError models.SysError) (err error) {
	err = config.DB.Model(&models.SysError{}).Where("id = ?", sysError.ID).Updates(&sysError).Error
	return err
}

// GetSysError 根据ID获取错误日志记录
func (sysErrorService *SysErrorService) GetSysError(ID string) (sysError models.SysError, err error) {
	err = config.DB.Where("id = ?", ID).First(&sysError).Error
	return
}

// GetSysErrorInfoList 分页获取错误日志记录
func (sysErrorService *SysErrorService) GetSysErrorInfoList(info SysErrorSearch) (list []models.SysError, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := config.DB.Model(&models.SysError{}).Order("created_at desc")
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
func (sysErrorService *SysErrorService) GetSysErrorSolution(ID string) (err error) {
	// 立即更新为处理中
	err = config.DB.Model(&models.SysError{}).Where("id = ?", ID).Update("status", "处理中").Error
	if err != nil {
		return err
	}
	// 异步协程在一分钟后更新为处理完成
	// todo 接入大模型
	return nil
}
