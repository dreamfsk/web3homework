package post

import (
	"com.dreamfsk/blog/repository/post"
)

type PostCreateReq struct {
	Title   string `json:"title" binding:"required,min=3,max=200"`
	Content string `json:"content" binding:"required"`
}
type PostUpdateReq struct {
	Content string `json:"content" binding:"required"`
}

type PostPageReq struct {
	PageNo    int    `json:"pageno" binding:"required"`
	PageSize  int    `json:"content" binding:"required"`
	Condition string `json:"condition"`
}

func (a *service) CreatePost(req *PostCreateReq) (*post.Post, error) {

	return nil, nil
}

func (a *service) GetPostPageList(req *PostPageReq) (*post.Post, error) {

	return nil, nil
}

func (a *service) GetPostById(id uint) (*post.Post, error) {

	return nil, nil
}

func (a *service) UpdatePost(req *PostUpdateReq) (*post.Post, error) {

	return nil, nil
}

func (a *service) DeletePost(id uint) (*post.Post, error) {

	return nil, nil
}
