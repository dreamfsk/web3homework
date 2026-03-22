package post

import (
	"com.dreamfsk/blog/commons"
	"com.dreamfsk/blog/repository/post"
	"com.dreamfsk/blog/utils"
)

func (a *service) CreatePost(req *PostCreateReq) (*post.Post, error) {
	p := post.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  req.UserId,
	}
	_, err := post.NewPostRepo(a.db, &p).CreatePost()
	if err != nil {
		return nil, utils.ServiceError(commons.PostAddError)
	}
	return &p, nil
}

func (a *service) GetPostPageList(req *PostPageReq) (*[]post.Post, error) {
	p := post.Post{
		UserID: req.UserId,
	}
	posts, err := post.NewPostRepo(a.db, &p).PostList()
	if err != nil {
		return nil, utils.ServiceError(commons.PostListError)
	}
	return posts, nil
}

func (a *service) GetPostById(id uint) (*post.Post, error) {
	p := post.Post{
		ID: id,
	}
	ps, err := post.NewPostRepo(a.db, &p).GetPostById()
	if err != nil {
		return nil, utils.ServiceError(commons.PostNotFoundError)
	}
	return ps, nil
}

func (a *service) UpdatePost(req *PostUpdateReq) (*post.Post, error) {
	p := post.Post{
		ID: req.ID,
	}
	check, err := a.GetPostById(req.ID)
	if err != nil {
		return nil, utils.ServiceError(commons.PostNotFoundError)
	}
	if req.UserId != check.UserID {
		return nil, utils.ServiceError(commons.PostUpdateAuthError)
	}
	_, err = post.NewPostRepo(a.db, &p).UpdatePostByMap(map[string]any{"content": req.Content})
	if err != nil {
		return nil, utils.ServiceError(commons.PostUpdateError)
	}
	res, err := a.GetPostById(req.ID)
	if err != nil {
		return nil, utils.ServiceError(commons.PostNotFoundError)
	}
	return res, nil
}

func (a *service) DeletePost(id uint) (int, error) {
	check, err := a.GetPostById(id)
	if err != nil {
		return 0, utils.ServiceError(commons.PostNotFoundError)
	}
	if id != check.UserID {
		return 0, utils.ServiceError(commons.PostDeleteAuthError)
	}
	err = post.NewPostRepo(a.db, &post.Post{ID: id}).DeletePost()
	if err != nil {
		return 0, utils.ServiceError(commons.PostDeleteError)
	}
	return 1, nil
}
