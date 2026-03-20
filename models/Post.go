package models

import (
	"gorm.io/gorm"
	"log"
	"time"
)

type Post struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Title         string         `json:"title" gorm:"not null"`
	Content       string         `json:"content" gorm:"type:text;not null"`
	UserID        uint           `json:"user_id" gorm:"not null"`
	Audit         Audit          `gorm:"embedded"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
	CommentStatus string         `json:"comment_status"`
	User          User           `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Comments      []Comment      `json:"comments,omitempty" gorm:"foreignKey:PostID"`
}

func (a *Post) BeforeCreate(tx *gorm.DB) error {
	user := CurrentOperator(tx)
	a.Audit.CreatedBy = user
	a.Audit.UpdatedBy = user
	return nil
}

func (a *Post) BeforeUpdate(tx *gorm.DB) error {
	user := CurrentOperator(tx)
	a.Audit.UpdatedBy = user
	return nil
}

func (a *Post) BeforeDelete(tx *gorm.DB) error {
	user := CurrentOperator(tx)
	a.Audit.DeletedBy = user
	return nil
}

func (a *Post) AfterCreate(tx *gorm.DB) error {
	user := CurrentOperator(tx)

	var count int64
	if err := tx.Model(&Post{}).Where("user_id = ?", a.UserID).Count(&count).Error; err != nil {
		log.Printf(" aftercreate post then count user post count err, userid %v,postid: %v,err: %v", a.UserID, a.ID, err)
	}

	if err := tx.Model(&User{}).Where("id = ?", a.UserID).
		Updates(map[string]any{"post_count": count, "updated_by": user}).Error; err != nil {
		log.Printf(" aftercreate post then update user post count err, userid %v,postid: %v,err: %v", a.UserID, a.ID, err)
	}
	log.Printf(" aftercreate post then update user post count suceessful, userid %v,postid: %v", a.UserID, a.ID)
	return nil
}
