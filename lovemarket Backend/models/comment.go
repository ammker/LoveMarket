package models

// 定义评论结构体
type Comment struct {
	CommentID  int    `json:"comment_id" db:"commentid"` // 评论ID
	UserID     int64  `json:"user_id" db:"userid"`       // 用户ID
	UserName   string `json:"user_name" db:"username"`
	PID        int64  `json:"postid" db:"pid"`
	Content    string `json:"content" db:"content"`        // 评论内容
	CreateTime string `json:"create_time" db:"createtime"` // 评论创建时间
	Status     int    `json:"status"`                      // 评论状态 (1=有效, 0=删除)
}
