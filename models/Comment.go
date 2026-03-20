package models

import (
	"gorm.io/gorm"
	"log"
	"time"
)

// Comment model

type Comment struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Content   string         `json:"content" gorm:"type:text;not null"`
	UserID    uint           `json:"user_id" gorm:"not null"`
	PostID    uint           `json:"post_id" gorm:"not null"`
	Audit     Audit          `gorm:"embedded"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Post Post `json:"post,omitempty" gorm:"foreignKey:PostID"`
}

func (a *Comment) BeforeCreate(tx *gorm.DB) error {
	user := CurrentOperator(tx)
	a.Audit.CreatedBy = user
	a.Audit.UpdatedBy = user
	return nil
}

func (a *Comment) BeforeUpdate(tx *gorm.DB) error {
	user := CurrentOperator(tx)
	a.Audit.UpdatedBy = user
	return nil
}

func (a *Comment) BeforeDelete(tx *gorm.DB) error {
	user := CurrentOperator(tx)
	a.Audit.DeletedBy = user
	// 如果 PostID
	log.Printf(" cur delete comment before: id: %v , postid: %v", a.ID, a.PostID)
	if a.ID != 0 && a.PostID == 0 {
		var tmp Comment
		if err := tx.Select("post_id").First(&tmp, a.ID).Error; err != nil {
			return err
		}
		a.PostID = tmp.PostID
	}
	return nil
}

func (a *Comment) AfterDelete(tx *gorm.DB) error {
	user := CurrentOperator(tx)
	log.Printf(" cur delete comment after: id: %v , postid: %v", a.ID, a.PostID)

	if a.PostID == 0 {
		return nil
	}

	var count int64
	if err := tx.Model(&Comment{}).Where("post_id = ?", a.PostID).Count(&count).Error; err != nil {
		log.Printf(" afterdelete comment then check post comments count err, id %v,postid: %v,err: %v", a.ID, a.PostID, err)
	}
	if count > 0 {
		return nil
	}
	if err := tx.Model(&Post{}).Where("id = ?", a.PostID).
		Updates(map[string]any{"comment_status": "无评论", "updated_by": user}).Error; err != nil {
		log.Printf(" afterdelete comment then update post commentstatus err, id %v,postid: %v,err: %v", a.ID, a.PostID, err)
	}
	log.Printf(" afterdelete comment then update post commentstatus suceessful, id %v,postid: %v", a.ID, a.PostID)
	return nil
}
