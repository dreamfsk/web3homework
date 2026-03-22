package post

type PostCreateReq struct {
	Title   string `json:"title" binding:"required,min=3,max=200"`
	Content string `json:"content" binding:"required"`
	UserId  uint   `json:"userId" binding:"required"`
}
type PostSearchReq struct {
	ID     uint `form:"id" json:"id"`
	UserId uint `form:"userId" json:"userId"`
}
type PostUpdateReq struct {
	ID      uint   `json:"id" binding:"required"`
	UserId  uint   `json:"userId" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type PostPageReq struct {
	PageNo    int    `json:"pageno"`
	PageSize  int    `json:"content"`
	UserId    uint   `json:"userId" binding:"required"`
	Condition string `json:"condition"`
}
