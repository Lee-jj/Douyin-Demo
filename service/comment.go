package service

type CommentResponse struct {
	ID       int64            `json:"id"`
	User     UserInfoResponse `json:"user"`
	Content  string           `json:"content"`
	CreateAt string           `json:"create_date"`
}

func CommentActionService() {

}
