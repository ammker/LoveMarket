package logic

import (
	"go.uber.org/zap"
	"lovemarket/dao/mysql"
	"lovemarket/models"
)

func CreatePost(p *models.Post) (post *models.Post, err error) {

	// 1. 保存到数据库
	post, err = mysql.CreatePost(p)
	if err != nil {
		return post, err
	}
	//err = redis.CreatePost(p.ID, p.CommunityID)
	return
	// 2. 返回
}
func DeletePost(p *models.DeletePost) (err error) {

	// 1. 保存到数据库
	err = mysql.DeletePost(p)
	if err != nil {
		return err
	}
	//err = redis.CreatePost(p.ID, p.CommunityID)
	return
	// 2. 返回
}
func CreateComment(p *models.Comment) (err error) {

	// 1. 保存到数据库
	user, err := mysql.GetUserById(p.UserID)
	if err != nil {
		zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
			zap.Int64("author_id", p.UserID),
			zap.Error(err))
		return
	}
	p.UserName = user.Username
	err = mysql.CreateComment(p)
	if err != nil {
		return err
	}
	//err = redis.CreatePost(p.ID, p.CommunityID)
	return
	// 2. 返回
}

// GetPostById 根据帖子id查询帖子详情数据
func GetPostById(pid int64) (data *models.ApiPostDetail2, err error) {
	// 查询并组合我们接口想用的数据
	post, err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostById(pid) failed",
			zap.Int64("pid", pid),
			zap.Error(err))
		return
	}
	// 根据作者id查询作者信息
	user, err := mysql.GetUserById(post.UserID)
	if err != nil {
		zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
			zap.Int64("author_id", post.UserID),
			zap.Error(err))
		return
	}
	// 根据社区id查询社区详细信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
			zap.Int64("community_id", post.CommunityID),
			zap.Error(err))
		return
	}
	comment, err := mysql.GetCommentListByPostID(pid)
	// 接口数据拼接
	data = &models.ApiPostDetail2{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
		Comments:        comment,
	}
	return
}

// GetPostList 获取帖子列表
func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))

	for _, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.UserID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("author_id", post.UserID),
				zap.Error(err))
			continue
		}

		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

// GetPostList 获取帖子列表
func GetPostListWithSearch(page, size int64, index string) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostListWithSearch(page, size, index)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))

	for _, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.UserID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("author_id", post.UserID),
				zap.Error(err))
			continue
		}

		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}
func GetPostByselfList(page, size, id int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostByselfList(page, size, id)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))

	for _, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.UserID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("author_id", post.UserID),
				zap.Error(err))
			continue
		}

		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}
func GetPostList2(page, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList2(page, size)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))

	for _, post := range posts {
		// 根据作者id查询作者信息
		user, err := mysql.GetUserById(post.UserID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("author_id", post.UserID),
				zap.Error(err))
			continue
		}

		// 根据社区id查询社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
				zap.Int64("community_id", post.CommunityID),
				zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

//func GetPostList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
//	// 2. 去redis查询id列表
//	ids, err := redis.GetPostIDsInOrder(p)
//	if err != nil {
//		return
//	}
//	if len(ids) == 0 {
//		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
//		return
//	}
//	zap.L().Debug("GetPostList2", zap.Any("ids", ids))
//	// 3. 根据id去MySQL数据库查询帖子详细信息
//	// 返回的数据还要按照我给定的id的顺序返回
//	posts, err := mysql.GetPostListByIDs(ids)
//	if err != nil {
//		return
//	}
//	zap.L().Debug("GetPostList2", zap.Any("posts", posts))
//	// 提前查询好每篇帖子的投票数
//	voteData, err := redis.GetPostVoteData(ids)
//	if err != nil {
//		return
//	}
//
//	// 将帖子的作者及分区信息查询出来填充到帖子中
//	for idx, post := range posts {
//		// 根据作者id查询作者信息
//		user, err := mysql.GetUserById(post.UserID)
//		if err != nil {
//			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
//				zap.Int64("author_id", post.UserID),
//				zap.Error(err))
//			continue
//		}
//		// 根据社区id查询社区详细信息
//		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
//		if err != nil {
//			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
//				zap.Int64("community_id", post.CommunityID),
//				zap.Error(err))
//			continue
//		}
//		postDetail := &models.ApiPostDetail{
//			AuthorName:      user.Username,
//			VoteNum:         voteData[idx],
//			Post:            post,
//			CommunityDetail: community,
//		}
//		data = append(data, postDetail)
//	}
//	return
//
//}

//func GetCommunityPostList(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
//	// 2. 去redis查询id列表
//	//ids, err := redis.GetCommunityPostIDsInOrder(p)
//	if err != nil {
//		return
//	}
//	if len(ids) == 0 {
//		zap.L().Warn("redis.GetPostIDsInOrder(p) return 0 data")
//		return
//	}
//	zap.L().Debug("GetCommunityPostIDsInOrder", zap.Any("ids", ids))
//	// 3. 根据id去MySQL数据库查询帖子详细信息
//	// 返回的数据还要按照我给定的id的顺序返回
//	posts, err := mysql.GetPostListByIDs(ids)
//	if err != nil {
//		return
//	}
//	zap.L().Debug("GetPostList2", zap.Any("posts", posts))
//	// 提前查询好每篇帖子的投票数
//	voteData, err := redis.GetPostVoteData(ids)
//	if err != nil {
//		return
//	}
//
//	// 将帖子的作者及分区信息查询出来填充到帖子中
//	for idx, post := range posts {
//		// 根据作者id查询作者信息
//		user, err := mysql.GetUserById(post.UserID)
//		if err != nil {
//			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
//				zap.Int64("author_id", post.UserID),
//				zap.Error(err))
//			continue
//		}
//		// 根据社区id查询社区详细信息
//		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
//		if err != nil {
//			zap.L().Error("mysql.GetUserById(post.AuthorID) failed",
//				zap.Int64("community_id", post.CommunityID),
//				zap.Error(err))
//			continue
//		}
//		postDetail := &models.ApiPostDetail{
//			AuthorName:      user.Username,
//			VoteNum:         voteData[idx],
//			Post:            post,
//			CommunityDetail: community,
//		}
//		data = append(data, postDetail)
//	}
//	return
//}

// GetPostListNew  将两个查询帖子列表逻辑合二为一的函数
//func GetPostListNew(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
//	// 根据请求参数的不同，执行不同的逻辑。
//
//	data, err = GetPostList2(p)
//
//	if err != nil {
//		zap.L().Error("GetPostListNew failed", zap.Error(err))
//		return nil, err
//	}
//	return
//}
