package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
)

func getIDsFormKey(key string, page, size int64) ([]string, error) {
	start := (page - 1) * size
	end := start + size - 1
	// 3. ZREVRANGE 按分数从大到小的顺序查询指定数量的元素
	return client.ZRevRange(context.Background(), key, start, end).Result()
}

//func GetPostVoteData2(id string) (data int64) {
//	//data = make([]int64, 0, len(ids))
//	//for _, id := range ids {
//	//	key := getRedisKey(KeyPostVotedZSetPF + id)
//	//	// 查找key中分数是1的元素的数量->统计每篇帖子的赞成票的数量
//	//	v := client.ZCount(key, "1", "1").Val()
//	//	data = append(data, v)
//	//}
//	// 使用pipeline一次发送多条命令,减少RTT
//
//	key := getRedisKey(KeyPostVotedZSetPF + id)
//	data = client.ZCount(context.Background(), key, "1", "1").Val()
//
//	return
//}

// GetPostVoteData 根据ids查询每篇帖子的投赞成票的数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	//data = make([]int64, 0, len(ids))
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVotedZSetPF + id)
	//	// 查找key中分数是1的元素的数量->统计每篇帖子的赞成票的数量
	//	v := client.ZCount(key, "1", "1").Val()
	//	data = append(data, v)
	//}
	// 使用pipeline一次发送多条命令,减少RTT
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSetPF + id)
		pipeline.ZCount(context.Background(), key, "1", "1")
	}
	cmders, err := pipeline.Exec(context.Background())
	if err != nil {
		return nil, err
	}
	data = make([]int64, 0, len(cmders))
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}
