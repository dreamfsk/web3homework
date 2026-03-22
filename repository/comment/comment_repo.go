package comment

import (
	"com.dreamfsk/blog/repository"
	"com.dreamfsk/blog/repository/post"
	"gorm.io/gorm"
	"log"
)

type CommentRepo struct {
	DB    *gorm.DB
	Model *Comment
}

func NewCommentRepo(db *gorm.DB, c ...*Comment) *CommentRepo {
	if len(c) == 0 {
		return &CommentRepo{db, new(Comment)}
	} else {
		return &CommentRepo{db, c[0]}
	}
}

func (a *Comment) BeforeCreate(tx *gorm.DB) error {
	user := repository.CurrentOperator(tx)
	a.Audit.CreatedBy = user
	a.Audit.UpdatedBy = user
	return nil
}

func (a *Comment) BeforeUpdate(tx *gorm.DB) error {
	user := repository.CurrentOperator(tx)
	a.Audit.UpdatedBy = user
	return nil
}

func (a *Comment) BeforeDelete(tx *gorm.DB) error {
	user := repository.CurrentOperator(tx)
	a.Audit.DeletedBy = user
	// 如果 PostID
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

	_, err := post.NewPostRepo(tx, &post.Post{ID: a.PostID}).UpdatePostByMap(map[string]any{"comment_status": "无评论"})
	if err != nil {
		log.Printf(" afterdelete comment then update post commentstatus err, id %v,postid: %v,err: %v", a.ID, a.PostID, err)
	}
	return nil
}

func (repo *CommentRepo) CreateComment() (cr *Comment, err error) {
	c := repo.Model
	err = repo.DB.Create(c).Error
	return c, err
}
func (repo *CommentRepo) ListComment() (crs *[]Comment, err error) {
	c := repo.Model
	var comments []Comment
	err = repo.DB.Model(&Comment{}).Where("post_id = ?", c.PostID).Scan(&comments).Error
	return &comments, err
}
