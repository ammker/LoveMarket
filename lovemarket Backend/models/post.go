package models

import "time"

// 内存对齐概念

type Post struct {
	ID          int64     `json:"id,string" db:"postid"`                            // 帖子id
	UserID      int64     `json:"user_id" db:"userid"`                              // 作者id
	CommunityID int64     `json:"community_id" db:"communityid" binding:"required"` // 社区id
	Title       string    `json:"title" db:"title" binding:"required"`              // 帖子标题
	Summary     string    `json:"summary" db:"summary" `
	Content     string    `json:"content" db:"content" binding:"required"` // 帖子内容
	CreateTime  time.Time `json:"create_time" db:"CreateTime"`             // 帖子创建时间
	VoteCount   int64     `json:"vote_count" db:"VoteCount"`
}
type DeletePost struct {
	ID     int64 `json:"postid,string" db:"postid"` // 帖子id
	UserID int64 `json:"user_id" db:"userid"`       // 作者id
}

// ApiPostDetail 帖子详情接口的结构体
type ApiPostDetail struct {
	AuthorName       string             `json:"author_name"` // 作者
	*Post                               // 嵌入帖子结构体
	*CommunityDetail `json:"community"` // 嵌入社区信息
}

// 修改 ApiPostDetail2 结构体
type ApiPostDetail2 struct {
	AuthorName       string             `json:"author_name"` // 作者
	*Post                               // 嵌入帖子结构体
	*CommunityDetail `json:"community"` // 嵌入社区信息
	Comments         []*Comment         `json:"comments"` // 评论切片
}
type VoteStatus struct {
	Status int64 `json:"status"`
}
