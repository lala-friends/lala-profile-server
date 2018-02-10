package domain

type Comment struct {
	CommentId int
	Email     string
	Message   string
	ParentId  int64
	ProductId int
}
