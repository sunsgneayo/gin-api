package common

import (
	"context"
	"dpj-admin-api/config"
	"fmt"
	"log"
	"time"
)

var ctx = context.Background()

const CAPTCHA = "captcha:"

type RedisStore struct {
}

// Set /** set a capt
func (r RedisStore) Set(id string, value string) error {
	key := CAPTCHA + id
	err := config.RedisDb.Set(ctx, key, value, time.Minute*2).Err()
	if err != nil {
		log.Println(err.Error())
	}
	return nil
}

// Get /** get a capt
func (r RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	val, err := config.RedisDb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if clear {
		err := config.RedisDb.Del(ctx, key).Err()
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return val
}

// Verify verify a capt
func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := RedisStore{}.Get(id, clear)
	//fmt.Println("key:"+id+";value:"+v+";answer:"+answer)
	return v == answer
}
