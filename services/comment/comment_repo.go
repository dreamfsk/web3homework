package comment

type CommentCreatReq struct {
	UserId  uint   `json:"userId" binding:"required"`
	PostId  uint   `json:"postId" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type CommentPageReq struct {
	PageNo   int  `json:"pageno"`
	PageSize int  `json:"content"`
	PostId   uint `json:"postId" binding:"required"`
}
