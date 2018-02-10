package domain

type Comment struct {
	CommentId int    `json:"commentId"`
	Email     string `json:"email"`
	Message   string `json:"message"`
	ParentId  int64  `json:"parentId"`
	ProductId int    `json:"productId"`
	RegDt     string `json:"regDt"`
	ModDt     string `json:"modDt"`
}