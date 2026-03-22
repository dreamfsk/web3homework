package repos

import (
	"com.dreamfsk/blog/repository/models"
	"gorm.io/gorm"
)

type PostRepo struct {
	DB    *gorm.DB
	Model *models.Post
}

func NewPostRepo(db *gorm.DB, p ...*models.Post) *PostRepo {
	if len(p) > 0 {
		return &PostRepo{db, p[0]}
	} else {
		return &PostRepo{db, new(models.Post)}
	}
}

func (repo *PostRepo) UpdatePostByMap(m map[string]any) (pr *models.Post, er error) {
	p := repo.Model
	pp := models.Post{}
	er = repo.DB.Model(&pp).Where("id = ?", p.ID).Updates(m).Error
	return &pp, er
}

func (repo *PostRepo) CreatePost() (pr *models.Post, er error) {
	p := repo.Model
	er = repo.DB.Create(p).Error
	return p, er
}

func (repo *PostRepo) PostList() (ps *[]models.Post, er error) {
	p := repo.Model
	var posts []models.Post
	er = repo.DB.Model(&models.Post{}).Where("user_id = ?", p.UserID).Scan(&posts).Error
	return &posts, er
}

func (repo *PostRepo) GetPostById() (ps *models.Post, er error) {
	p := repo.Model
	var post models.Post
	er = repo.DB.First(&post, p.ID).Error
	return &post, er
}

func (repo *PostRepo) DeletePost() error {
	p := repo.Model
	return repo.DB.Delete(&models.Post{}, p.ID).Error
}
