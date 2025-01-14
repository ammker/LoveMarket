package models

import "time"

type Community struct {
	ID   int64  `json:"id" db:"communityId"`
	Name string `json:"name" db:"communityName"`
}

type CommunityDetail struct {
	ID           int64     `json:"id" db:"communityId"`
	Name         string    `json:"name" db:"communityName"`
	Introduction string    `json:"introduction,omitempty" db:"introduction"`
	CreateTime   time.Time `json:"create_time" db:"createTime"`
}
