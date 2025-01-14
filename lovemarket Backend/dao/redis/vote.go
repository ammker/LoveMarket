package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"lovemarket/dao/mysql"
	"time"
)

// 推荐阅读
// 基于用户投票的相关算法：http://www.ruanyifeng.com/blog/algorithm/

// 本项目使用简化版的投票分数
// 投一票就加432分   86400/200  --> 200张赞成票可以给你的帖子续一天

/* 投票的几种情况：
   direction=1时，有两种情况：
   	1. 之前没有投过票，现在投赞成票    --> 更新分数和投票记录  差值的绝对值：1  +432
   	2. 之前投反对票，现在改投赞成票    --> 更新分数和投票记录  差值的绝对值：2  +432*2
   direction=0时，有两种情况：
   	1. 之前投过反对票，现在要取消投票  --> 更新分数和投票记录  差值的绝对值：1  +432
	2. 之前投过赞成票，现在要取消投票  --> 更新分数和投票记录  差值的绝对值：1  -432
   direction=-1时，有两种情况：
   	1. 之前没有投过票，现在投反对票    --> 更新分数和投票记录  差值的绝对值：1  -432
   	2. 之前投赞成票，现在改投反对票    --> 更新分数和投票记录  差值的绝对值：2  -432*2

   投票的限制：
   每个贴子自发表之日起一个星期之内允许用户投票，超过一个星期就不允许再投票了。
   	1. 到期之后将redis中保存的赞成票数及反对票数存储到mysql表中
   	2. 到期之后删除那个 KeyPostVotedZSetPF
*/

// 实际生产环境下 context.Background() 按需替换

var (
	ErrVoteRepeated = errors.New("不允许重复投票")
)

func CreatePost(postID, communityID int64) error {
	pipeline := client.TxPipeline()
	// 帖子时间
	pipeline.ZAdd(context.Background(), getRedisKey(KeyPostTimeZSet), &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	})

	// 帖子分数
	pipeline.ZAdd(context.Background(), getRedisKey(KeyPostScoreZSet), &redis.Z{
		Score:  float64(0),
		Member: postID,
	})
	//// 更新：把帖子id加到社区的set
	//cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityID)))
	//pipeline.SAdd(context.Background(), cKey, postID)
	_, err := pipeline.Exec(context.Background())
	return err
}
func VoteStatusPost(userID, postID string) (value int64, err error) {
	// 从 Redis 获取用户对帖子的投票值
	redisKey := getRedisKey(KeyPostVotedZSetPF + postID) // Redis 中的键
	score := client.ZScore(context.Background(), redisKey, userID).Val()

	if score == 0 {
		// 如果 score 为 0，表示用户未投票或该键不存在
		return 0, nil
	}

	// 如果有投票值，将其转为整型返回
	value = int64(score)
	return value, nil
}

func VoteForPost(userID, postID string, value float64) error {
	// 1. 查询用户对该帖子的投票记录
	// 从 Redis 中获取当前用户对帖子的投票值
	ov := client.ZScore(context.Background(), getRedisKey(KeyPostVotedZSetPF+postID), userID).Val()

	// 检查重复投票
	if value == ov {
		return ErrVoteRepeated // 如果投票值相同，提示重复投票
	}

	// 2. 记录用户的投票信息
	// 使用 Redis Pipeline 来确保事务一致性
	pipeline := client.TxPipeline()
	db := mysql.Getdb()
	if value == 0 {
		// 如果投票值为 0，表示取消投票，直接删除记录
		pipeline.ZRem(context.Background(), getRedisKey(KeyPostVotedZSetPF+postID), userID)
		sqlStr := `UPDATE posts SET VoteCount = VoteCount - 1 WHERE PostID = ?`
		_, err := db.Exec(sqlStr, postID)
		if err != nil {
			return fmt.Errorf("failed to update vote count in MySQL: %w", err)
		}
	} else {
		// 否则，记录用户对该帖子的投票值
		pipeline.ZAdd(context.Background(), getRedisKey(KeyPostVotedZSetPF+postID), &redis.Z{
			Score:  value, // 赞成票还是反对票
			Member: userID,
		})
		sqlStr := `UPDATE posts SET VoteCount = VoteCount + 1 WHERE PostID = ?`
		_, err := db.Exec(sqlStr, postID)
		if err != nil {
			return fmt.Errorf("failed to update vote count in MySQL: %w", err)
		}
	}

	// 执行 Redis Pipeline
	_, err := pipeline.Exec(context.Background())
	return err
}
