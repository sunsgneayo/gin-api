package service

import (
	"context"
	"dpj-admin-api/config"
	"encoding/json"
	"fmt"
	"github.com/wenlng/go-captcha/captcha"
	"strconv"
	"strings"
	"time"
)

// ctx /** 默认协程上下文
var ctx = context.Background()

// CheckCapByKeyValue /**  使用传值的方式验证行为验证码
func CheckCapByKeyValue(key string, val string) bool {
	cacheData, _ := config.RedisDb.Get(ctx, key).Result()

	src := strings.Split(val, ",")
	var dct map[int]captcha.CharDot
	if err := json.Unmarshal([]byte(cacheData), &dct); err != nil {
		return false
	}

	chkRet := false
	if (len(dct) * 2) == len(src) {
		for i, dot := range dct {
			j := i * 2
			k := i*2 + 1
			sx, _ := strconv.ParseFloat(fmt.Sprintf("%v", src[j]), 64)
			sy, _ := strconv.ParseFloat(fmt.Sprintf("%v", src[k]), 64)
			// 校验点的位置,在原有的区域上添加额外边距进行扩张计算区域,不推荐设置过大的padding
			// 例如：文本的宽和高为30，校验范围x为10-40，y为15-45，此时扩充5像素后校验范围宽和高为40，则校验范围x为5-45，位置y为10-50
			chkRet = captcha.CheckPointDistWithPadding(int64(sx), int64(sy), int64(dot.Dx), int64(dot.Dy), int64(dot.Width), int64(dot.Height), 5)
			if !chkRet {
				break
			}
		}
	}

	if chkRet {
		//删除redis缓存
		err := config.RedisDb.Del(ctx, key)
		if err != nil {
		}
		return true
	}
	return false
}

// CapWriteCache /** 将验证码写入缓存 （redis）
func CapWriteCache(v interface{}, file string) error {

	bt, _ := json.Marshal(v)
	err := config.RedisDb.Set(ctx, file, bt, time.Minute*2).Err()
	if err != nil {
		return err
	}
	return nil
}

// GenerateCaptcha /** 生产行为验证码
func GenerateCaptcha() (map[int]captcha.CharDot, string, string, string, error) {
	capt := captcha.GetCaptcha()

	//限制随机验证字符大小
	capt.SetRangCheckTextLen(captcha.RangeVal{Min: 2, Max: 2})
	// Generate Captcha
	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		return nil, "", "", "", err
	}
	// redis 存储验证码
	err = CapWriteCache(dots, key)
	if err != nil {
		return nil, "", "", "", err
	}
	// 原样返回
	return dots, b64, tb64, key, nil
}
