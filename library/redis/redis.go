package redis

import (
	"context"
	"github.com/cfx/warehouses/app"
	"github.com/cfx/warehouses/library/utils"
	"github.com/cfx/warehouses/output"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"time"
)

func RPush(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	r := app.Redis("local")
	err := r.RPush(ctx, key, val).Err()
	if err != nil {
		return err
	}

	return r.Expire(ctx, key, expiration).Err()
}

func LPop(ctx context.Context, key string) string {
	r := app.Redis("local")
	res := r.LPop(ctx, key).Val()

	return res
}

func TTL(ctx context.Context, key string) float64 {
	res := app.Redis("local").TTL(ctx, key).Val()
	if res.Seconds() <= 0 {
		return 0
	}
	return res.Seconds()
}

func GetStr(ctx context.Context, key string) string {
	return app.Redis("local").Get(ctx, key).Val()
}

func SetVal(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	err := app.Redis("local").Set(ctx, key, val, expiration).Err()
	if err != nil {
		app.Log().WithField("key", key).WithField("val", val).Warn(err)
	}
	return err
}

func SetNx(ctx context.Context, key string, val interface{}, expiration time.Duration) (bool, error) {
	ok, err := app.Redis("local").SetNX(ctx, key, val, expiration).Result()
	if err != nil {
		app.Log().WithField("key", key).WithField("val", val).Warn(err)
	}
	return ok, err
}
func Exists(ctx context.Context, key string) bool {
	return app.Redis("local").Exists(ctx, key).Val() == 1
}

func Del(ctx context.Context, key string) bool {
	if ok := Exists(ctx, key); ok {
		return app.Redis("local").Del(ctx, key).Val() == 1
	}
	return true
}

func Expire(ctx context.Context, key string, expiration time.Duration) error {
	return app.Redis("local").Expire(ctx, key, expiration).Err()
}

func Lock(ctx context.Context, key string, expires time.Duration) (*unLocker, error) {
	token, _ := uuid.NewUUID()
	boolRes := app.Redis("local").SetNX(ctx, key, token.String(), expires)
	if !boolRes.Val() {
		err := boolRes.Err()
		if err == nil {
			err = output.NewErrLog("SetNX failed, key=" + key)
		}
		return nil, err
	}
	holder := &unLocker{key, token.String()}
	return holder, nil
}

// package internal use only
type unLocker struct {
	key   string
	token string
}

const unlockScript = `if redis.call("get",KEYS[1]) == ARGV[1]
then
    return redis.call("del",KEYS[1])
else
    return 0
end`

// 通过 lua 脚本安全解锁
// see https://redis.io/commands/set
func (h *unLocker) Unlock(ctx context.Context) error {
	cmd := app.Redis("local").Eval(ctx, unlockScript, []string{h.key}, h.token)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	val := cmd.Val().(int64)
	if val != 1 {
		return output.NewErrLog("lock expired or not exists or relock by others").Log(nil)
	}
	return nil
}

func GetListLen(ctx context.Context, key string) (int64, error) {
	length, err := app.Redis("local").LLen(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return length, nil
}

func GetListAll(ctx context.Context, key string, start int64, end int64) ([]string, error) {
	values, err := app.Redis("local").LRange(ctx, key, start, end).Result()
	if err != nil {
		return nil, err
	}
	return values, nil
}

// 锁定
func CommonLocked(ctx context.Context, key string, d time.Duration) (bool, error) {
	now := time.Now()
	t := now.Add(d)
	differ := t.Unix() - now.Unix()
	res, err := SetNx(ctx, key, t.Unix(), time.Duration(differ)*time.Second)

	return res, err
}

// 是否锁定
// 0 && nil 表示未锁定
func GetCommonLocked(ctx context.Context, key string) (int64, error) {

	if ok := Exists(ctx, key); !ok {
		return -1, nil
	}
	res := GetStr(ctx, key)
	if res != "" {
		return utils.StringToInt64(res) - utils.NowSec().Unix(), nil
	}

	return 0, nil
}

func SetBit(ctx context.Context, key string, offset int64, expiration time.Duration) error {
	err := app.Redis("local").SetBit(ctx, key, offset, 1).Err()
	if err != nil {
		app.Log().WithField("key", key).WithField("offset", offset).Warn(err)
		return err
	}
	return app.Redis("local").Expire(ctx, key, expiration).Err()
}

func BitCount(ctx context.Context, key string) (int64, error) {

	l, err := app.Redis("local").StrLen(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	sit := redis.BitCount{Start: 0, End: l - 1}

	num, err := app.Redis("local").BitCount(ctx, key, &sit).Result()
	if err != nil {
		return 0, err
	}

	return num, nil
}
