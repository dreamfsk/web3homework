package repository

import (
	"com.dreamfsk/blog/commons"
	"gorm.io/gorm"
)

type Audit struct {
	CreatedBy string
	UpdatedBy string
	DeletedBy string
}

func CurrentOperator(tx *gorm.DB) string {
	if tx != nil && tx.Statement != nil && tx.Statement.Context != nil {
		if v, ok := tx.Statement.Context.Value(commons.CtxKey).(string); ok && v != "" {
			return v
		}
	}
	return "system"
}
