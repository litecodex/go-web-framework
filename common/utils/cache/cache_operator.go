package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	ObjectUtil "github.com/litecodex/go-web-framework/common/utils/object"
	"time"
)

type StringRedisOperator struct {
	RedisClient *redis.Client
}

func NewStringRedisOperator(redisClient *redis.Client) *StringRedisOperator {
	return &StringRedisOperator{
		RedisClient: redisClient,
	}
}

func (operator *StringRedisOperator) Get(redisKey string) (string, error) {
	result := operator.RedisClient.Get(context.TODO(), redisKey)
	return result.Val(), result.Err()
}

func (operator *StringRedisOperator) MustGet(redisKey string) string {
	result, err := operator.Get(redisKey)
	if err != nil {
		panic(err)
	}
	return result
}

func (operator *StringRedisOperator) Set(redisKey string, value interface{}, expiration time.Duration) error {
	return operator.RedisClient.Set(context.TODO(), redisKey, value, expiration).Err()
}

func (operator *StringRedisOperator) Delete(redisKey string) error {
	return operator.RedisClient.Del(context.TODO(), redisKey).Err()
}

func (operator *StringRedisOperator) MustDelete(redisKey string) {
	err := operator.Delete(redisKey)
	if err != nil {
		panic(err)
	}
}

const limitationLuaScript string = " local key = KEYS[1] " +
	" local ttl = ARGV[2] " +
	" local maxCount = ARGV[1] " +
	" local reqCounts = redis.call('get', key) " +
	" if (not reqCounts or tonumber(reqCounts) < tonumber(maxCount)) then " +
	"	 reqCounts = redis.call('incr', key) " +
	"	 if tonumber(reqCounts) == 1 then " +
	"		 redis.call('expire', key, tonumber(ttl)) " +
	"	 end " +
	"	 return 1 " +
	" end " +
	" if tonumber(redis.call('ttl', key)) <= 0 then " +
	"	 local res = redis.call('set', key, 1, 'ex', tonumber(ttl)) " +
	"	 if res.ok ~= \"OK\" then " +
	"		 return 2 " +
	"	 end " +
	"	 return 1 " +
	" end " +
	" return 2 "

func (operator *StringRedisOperator) IsOverLimit(key string, count uint, ttl int64) bool {
	redisClient := operator.RedisClient
	result, err := redisClient.Eval(context.TODO(), limitationLuaScript, []string{key}, count, ttl).Result()
	if err != nil {
		fmt.Println("执行计数器脚本出错：\n", err)
		return true
	}

	if ObjectUtil.ToIntValue(result) != 1 {
		// 超过达到限流值了。
		return true
	}

	// 未超过限流值，一切正常
	return false
}
