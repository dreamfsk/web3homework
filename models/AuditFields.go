package models

import (
	"gorm.io/gorm"
	"homework03/utils"
)

type Audit struct {
	CreatedBy string
	UpdatedBy string
	DeletedBy string
}

func CurrentOperator(tx *gorm.DB) string {
	if tx != nil && tx.Statement != nil && tx.Statement.Context != nil {
		if v, ok := tx.Statement.Context.Value(utils.CTX_KEY_OP).(string); ok && v != "" {
			return v
		}
	}
	return "system"
}
