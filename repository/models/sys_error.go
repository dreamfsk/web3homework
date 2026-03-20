// 自动生成模板SysError
package models

import (
	"gorm.io/gorm"
)

// 错误日志 结构体  SysError
type SysError struct {
	ID        uint           `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt CustomTime     // 创建时间
	UpdatedAt CustomTime     // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                                                                 // 删除时间
	Form      *string        `json:"form" form:"form" gorm:"comment:错误来源;column:form;type:text;" binding:"required"` //错误来源
	Info      *string        `json:"info" form:"info" gorm:"comment:错误内容;column:info;type:text;"`                    //错误内容
	Level     string         `json:"level" form:"level" gorm:"comment:日志等级;column:level;"`
	Solution  *string        `json:"solution" form:"solution" gorm:"comment:解决方案;column:solution;type:text"`               //解决方案
	Status    string         `json:"status" form:"status" gorm:"comment:处理状态;column:status;type:varchar(20);default:未处理;"` //处理状态：未处理/处理中/处理完成
}

// TableName 错误日志 SysError自定义表名 sys_error
func (SysError) TableName() string {
	return "sys_error"
}
