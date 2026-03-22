package models

import (
	"gorm.io/gorm"
	"log"
)

type Post struct {
	ID            uint       `json:"id" gorm:"primaryKey"`
	Title         string     `json:"title" gorm:"not null"`
	Content       string     `json:"content" gorm:"type:text;not null"`
	UserID        uint       `json:"userId" gorm:"not null"`
	Audit         Audit      `gorm:"embedded"`
	CreatedAt     CustomTime `json:"createdAt"`
	UpdatedAt     CustomTime `json:"updatedAt"`
	DeletedAt     CustomTime `json:"-" gorm:"index"`
	CommentStatus string     `json:"commentStatus"`
	User          User       `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Comments      []Comment  `json:"comments,omitempty" gorm:"foreignKey:PostID"`
}

func (a *Post) BeforeCreate(tx *gorm.DB) error {
	p := CurrentOperator(tx)
	a.Audit.CreatedBy = p
	a.Audit.UpdatedBy = p
	return nil
}

func (a *Post) BeforeUpdate(tx *gorm.DB) error {
	p := CurrentOperator(tx)
	a.Audit.UpdatedBy = p
	return nil
}

func (a *Post) BeforeDelete(tx *gorm.DB) error {
	p := CurrentOperator(tx)
	a.Audit.DeletedBy = p
	return nil
}

func (a *Post) AfterCreate(tx *gorm.DB) error {
	u := CurrentOperator(tx)
	var count int64
	if err := tx.Model(&Post{}).Where("user_id = ?", a.UserID).Count(&count).Error; err != nil {
		log.Printf(" aftercreate post then count user post count err, userid %v,postid: %v,err: %v", a.UserID, a.ID, err)
	}
	err := tx.Model(&User{}).Where("id = ?", a.UserID).Updates(map[string]any{"post_count": count, "updated_by": u}).Error
	if err != nil {
		log.Printf(" aftercreate post then update user post count err, userid %v,postid: %v,err: %v", a.UserID, a.ID, err)
	}
	return nil
}
