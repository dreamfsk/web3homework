package post

import (
	"com.dreamfsk/blog/repository"
	"com.dreamfsk/blog/repository/user"
	"gorm.io/gorm"
	"log"
)

type PostRepo struct {
	DB    *gorm.DB
	Model *Post
}

func NewPostRepo(db *gorm.DB, p ...*Post) *PostRepo {
	if len(p) > 0 {
		return &PostRepo{db, p[0]}
	} else {
		return &PostRepo{db, new(Post)}
	}
}

func (a *Post) BeforeCreate(tx *gorm.DB) error {
	p := repository.CurrentOperator(tx)
	a.Audit.CreatedBy = p
	a.Audit.UpdatedBy = p
	return nil
}

func (a *Post) BeforeUpdate(tx *gorm.DB) error {
	p := repository.CurrentOperator(tx)
	a.Audit.UpdatedBy = p
	return nil
}

func (a *Post) BeforeDelete(tx *gorm.DB) error {
	p := repository.CurrentOperator(tx)
	a.Audit.DeletedBy = p
	return nil
}

func (a *Post) AfterCreate(tx *gorm.DB) error {
	var count int64
	if err := tx.Model(&Post{}).Where("user_id = ?", a.UserID).Count(&count).Error; err != nil {
		log.Printf(" aftercreate post then count user post count err, userid %v,postid: %v,err: %v", a.UserID, a.ID, err)
	}
	_, err := user.NewRepo(tx, &user.User{ID: a.UserID}).UpdateUserByMap(map[string]any{"post_count": count})
	if err != nil {
		log.Printf(" aftercreate post then update user post count err, userid %v,postid: %v,err: %v", a.UserID, a.ID, err)
	}
	return nil
}

func (repo *PostRepo) UpdatePostByMap(m map[string]any) (pr *Post, er error) {
	p := repo.Model
	// au := repository.CurrentOperator(repo.DB)
	// m["updated_by"] = au
	pp := Post{}
	er = repo.DB.Model(&pp).Where("id = ?", p.ID).Updates(m).Error
	return &pp, er
}

func (repo *PostRepo) CreatePost() (pr *Post, er error) {
	p := repo.Model
	er = repo.DB.Create(p).Error
	return p, er
}

func (repo *PostRepo) PostList() (ps *[]Post, er error) {
	p := repo.Model
	var posts []Post
	er = repo.DB.Model(&Post{}).Where("user_id = ?", p.UserID).Scan(&posts).Error
	return &posts, er
}

func (repo *PostRepo) GetPostById() (ps *Post, er error) {
	p := repo.Model
	var post Post
	er = repo.DB.First(&post, p.ID).Error
	return &post, er
}

func (repo *PostRepo) DeletePost() error {
	p := repo.Model
	return repo.DB.Delete(&Post{}, p.ID).Error
}
