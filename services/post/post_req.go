package post

type PostCreateReq struct {
	Title   string `json:"title" binding:"required,min=3,max=200"`
	Content string `json:"content" binding:"required"`
	UserId  uint   `json:"userId" binding:"required"`
}
type PostUpdateReq struct {
	ID      uint   `json:"id" binding:"required"`
	UserId  uint   `json:"userId" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type PostPageReq struct {
	PageNo    int    `json:"pageno" binding:"required"`
	PageSize  int    `json:"content" binding:"required"`
	UserId    uint   `json:"userId" binding:"required"`
	Condition string `json:"condition"`
}
