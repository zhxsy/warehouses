package models

type BroadcastMessage struct {
	Title   string  `json:"title" form:"title"`
	Content string  `json:"content" form:"content"`
	UserIds []int64 `json:"user_ids" form:"user_ids"`
}

type SystemMessage struct {
	Title      string `json:"title" form:"title"`
	Content    string `json:"content" form:"content"`
	FromUserId int64  `json:"from_user_id" form:"from_user_id"`
	ToUserId   int64  `json:"to_user_id" form:"to_user_id"`
}
