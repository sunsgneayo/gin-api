package captcha

import (
	"context"
	"dpj-admin-api/config"
	"fmt"
	"github.com/mojocn/base64Captcha"
	"strings"
	"time"
)

func DefaultRedisStore() *RedisStore {
	return &RedisStore{
		Expiration: time.Second * 180,
		PreKey:     "CAPTCHA:",
		Context:    context.TODO(),
	}
}

type RedisStore struct {
	Expiration time.Duration
	PreKey     string
	Context    context.Context
}

func (rs *RedisStore) UseWithCtx(ctx context.Context) base64Captcha.Store {
	rs.Context = ctx
	return rs
}

func (rs *RedisStore) Set(id string, value string) error {

	fmt.Println(rs.Context)

	err := config.RedisDb.Set(rs.Context, rs.PreKey+id, value, rs.Expiration).Err()
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}

func (rs *RedisStore) Get(key string, clear bool) string {
	val, err := config.RedisDb.Get(rs.Context, key).Result()

	if err != nil {
		fmt.Println("获取失败", err)
		return ""
	}
	if clear {
		err := config.RedisDb.Del(rs.Context, key).Err()
		if err != nil {
			fmt.Println("删除失败")
			return ""
		}
	}
	return val
}

func (rs *RedisStore) Verify(id, answer string, clear bool) bool {
	key := rs.PreKey + id
	//去掉换行
	v := rs.Get(strings.Trim(key, "\n"), clear)
	fmt.Println(v)
	return v == answer
}
