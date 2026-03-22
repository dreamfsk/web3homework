package repos

import (
	"com.dreamfsk/blog/repository/models"
	"gorm.io/gorm"
)

type CommentRepo struct {
	DB    *gorm.DB
	Model *models.Comment
}

func NewCommentRepo(db *gorm.DB, c ...*models.Comment) *CommentRepo {
	if len(c) == 0 {
		return &CommentRepo{db, new(models.Comment)}
	} else {
		return &CommentRepo{db, c[0]}
	}
}

func (repo *CommentRepo) CreateComment() (cr *models.Comment, err error) {
	c := repo.Model
	err = repo.DB.Create(c).Error
	return c, err
}
func (repo *CommentRepo) ListComment() (crs *[]models.Comment, err error) {
	c := repo.Model
	var comments []models.Comment
	err = repo.DB.Model(&models.Comment{}).Where("post_id = ?", c.PostID).Scan(&comments).Error
	return &comments, err
}
