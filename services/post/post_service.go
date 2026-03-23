package post

import (
	"com.dreamfsk/blog/repository/models"
	"com.dreamfsk/blog/repository/repos"
)

func (a *service) CreatePost(req *PostCreateReq) (*models.Post, error) {
	p := models.Post{
		Title:   req.Title,
		Content: req.Content,
		UserID:  req.UserId,
	}
	_, err := repos.NewPostRepo(a.db, &p).CreatePost()
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (a *service) GetPostPageList(req *PostPageReq) (*[]models.Post, error) {
	p := models.Post{
		UserID: req.UserId,
	}
	posts, err := repos.NewPostRepo(a.db, &p).PostList()
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func (a *service) GetPostById(id uint) (*models.Post, error) {
	p := models.Post{
		ID: id,
	}
	ps, err := repos.NewPostRepo(a.db, &p).GetPostById()
	if err != nil {
		return nil, err
	}
	return ps, nil
}

func (a *service) UpdatePost(req *PostUpdateReq) (*models.Post, error) {
	p := models.Post{
		ID: req.ID,
	}
	check, err := a.GetPostById(req.ID)
	if err != nil {
		return nil, err
	}
	if req.UserId != check.UserID {
		return nil, err
	}
	_, err = repos.NewPostRepo(a.db, &p).UpdatePostByMap(map[string]any{"content": req.Content})
	if err != nil {
		return nil, err
	}
	res, err := a.GetPostById(req.ID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a *service) DeletePost(id uint) (int, error) {
	check, err := a.GetPostById(id)
	if err != nil {
		return 0, err
	}
	if id != check.UserID {
		return 0, err
	}
	err = repos.NewPostRepo(a.db, &models.Post{ID: id}).DeletePost()
	if err != nil {
		return 0, err
	}
	return 1, nil
}
