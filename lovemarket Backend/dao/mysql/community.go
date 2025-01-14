package mysql

import (
	"database/sql"
	"lovemarket/models"

	"go.uber.org/zap"
)

func GetCommunityList() (communityList []*models.Community, err error) {
	sqlStr := "select communityId, communityName from community"
	if err := db.Select(&communityList, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}
func GetCommentListByPostID(pid int64) (commentList []*models.Comment, err error) {
	// 定义 SQL 查询语句
	sqlStr := "SELECT pid, userid,username,content, createtime FROM comment WHERE pid = ?"

	// 执行查询并绑定结果到 commentList
	err = db.Select(&commentList, sqlStr, pid)
	if err != nil {
		// 检查是否是没有行的错误
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no comment in db for the given post id", zap.Int64("pid", pid))
			err = nil // 重置错误为 nil
			return
		}
		// 如果是其他错误，直接返回
		zap.L().Error("failed to fetch comments", zap.Error(err))
		return
	}

	return
}

// GetCommunityDetailByID 根据ID查询社区详情
func GetCommunityDetailByID(id int64) (community *models.CommunityDetail, err error) {
	sqlStr := `select 
			communityId, communityName, introduction, createTime
			from community 
			where communityId = ?
	`

	// 初始化 community
	community = &models.CommunityDetail{}

	// 执行查询
	if err = db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID // 自定义的无效ID错误
		}
	}
	return community, err
}
