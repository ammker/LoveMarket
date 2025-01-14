package mysql

import (
	"fmt"
	"lovemarket/models"
	"strings"

	"github.com/jmoiron/sqlx"
)

// CreatePost 创建帖子
// CreatePost 创建帖子
func CreatePost(p *models.Post) (post *models.Post, err error) {
	sqlStr := `insert into posts(
		title, summary, content, userid, communityid)
	values ( ?, LEFT(?, 30), ?, ?, ?)
	`

	// 执行插入并获取结果
	result, err := db.Exec(sqlStr, p.Title, p.Content, p.Content, p.UserID, p.CommunityID)
	if err != nil {
		return post, err
	}

	// 获取插入后自动生成的 ID
	id, err := result.LastInsertId()
	if err != nil {
		return post, err
	}

	p.ID = id // 设置到结构体中，便于后续使用
	post = p
	return post, nil
}

// DeletePost 更新帖子状态为逻辑删除
func DeletePost(p *models.DeletePost) (err error) {
	// SQL语句：将帖子状态更新为 0（逻辑删除）
	sqlStr := `UPDATE posts SET status = 0 WHERE postid = ? AND userid = ?`

	// 执行更新操作
	_, err = db.Exec(sqlStr, p.ID, p.UserID)
	if err != nil {
		return err // 如果更新操作失败，返回错误
	}

	return nil // 如果更新成功，返回nil
}

func CreateComment(c *models.Comment) (err error) {
	// SQL 查询语句，插入一条新的评论记录到 "comment" 表
	sqlStr := `INSERT INTO comment (
        pid, userid, username, content
    ) VALUES (?, ?, ?, ?)`

	// 执行查询，使用 Comment 结构体中的字段
	result, err := db.Exec(sqlStr, c.PID, c.UserID, c.UserName, c.Content)
	if err != nil {
		return err // 如果查询执行失败，返回错误
	}

	// 获取新插入记录的自动生成 ID
	id, err := result.LastInsertId()
	if err != nil {
		return err // 如果获取 LastInsertId 失败，返回错误
	}

	c.CommentID = int(id) // 将生成的 CommentID 设置到 Comment 结构体中
	return nil            // 返回 nil，表示成功
}

func VoteForPost(postID string, direction int8) error {
	// 确保 Direction 为合法值（1 表示投票 +1，-1 表示投票 -1）
	if direction == 1 && direction != -1 {
		return fmt.Errorf("invalid vote direction: %d", direction)
	}

	// 更新 MySQL 中帖子投票数
	sqlStr := `UPDATE posts SET VoteCount = VoteCount + ? WHERE PostID = ?`
	_, err := db.Exec(sqlStr, direction, postID)
	if err != nil {
		return fmt.Errorf("failed to update vote count: %w", err)
	}

	return nil
}

// GetPostById 根据id查询单个贴子数据
func GetPostById(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select
	postid, title,summary,content, userid, communityid, CreateTime,VoteCount
	from posts
	where postid = ?
	`
	err = db.Get(post, sqlStr, pid)
	return
}

// GetPostList 查询帖子列表函数
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select 
	postid, title,summary, content, userid, communityid, CreateTime,VoteCount
	from posts
	WHERE status = 1
	ORDER BY CreateTime
	DESC
	limit ?,?
	`
	posts = make([]*models.Post, 0, 2) // 不要写成make([]*models.Post, 2)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}

func GetPostListWithSearch(page, size int64, index string) (posts []*models.Post, err error) {
	// 定义基础 SQL 查询
	sqlStr := `select 
	postid, title, summary, content, userid, communityid, CreateTime, VoteCount
	from posts
	WHERE status = 1 AND content LIKE ?
	ORDER BY CreateTime DESC
	limit ?, ?`

	// 创建存储结果的切片
	posts = make([]*models.Post, 0, size)

	// 设置模糊搜索的参数值
	searchIndex := "%" + index + "%"

	// 执行查询，分页参数 (page-1)*size 表示偏移量
	err = db.Select(&posts, sqlStr, searchIndex, (page-1)*size, size)
	return
}

// GetPostList 查询帖子列表函数
func GetPostByselfList(page, size, userid int64) (posts []*models.Post, err error) {
	sqlStr := `select 
	postid,title,summary, content, userid, communityid, CreateTime,VoteCount
	from posts
	WHERE status = 1 and userid = ?
	ORDER BY CreateTime
	DESC
	limit ?,?
	`
	posts = make([]*models.Post, 0, 2) // 不要写成make([]*models.Post, 2)
	err = db.Select(&posts, sqlStr, userid, (page-1)*size, size)
	return
}
func GetPostList2(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select 
	postid, title,summary, content, userid, communityid, CreateTime,VoteCount
	from posts
	WHERE status = 1
	ORDER BY VoteCount
	DESC
	limit ?,?
	`
	posts = make([]*models.Post, 0, 2) // 不要写成make([]*models.Post, 2)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}

// GetPostListByIDs 根据给定的id列表查询帖子数据
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select postid, title, content, userid, communityid,summary, createdAt
	from posts
	where postid in (?)
	order by FIND_IN_SET(postid, ?)
	`
	// https: //www.liwenzhou.com/posts/Go/sqlx/
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...) // !!!!!!
	return
}
