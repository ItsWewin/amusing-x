package session

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ItsWewin/superfactory/xerror"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisStore struct {
	Client        *redis.Client
	KeyPrefix     string
	MaxLifeSecond int64
}

func InitRedisStore(addr, password string, db int, prefix string, maxLifeSecond int64) (Store, error) {
	if len(addr) == 0 {
		return nil, xerror.NewError(nil, xerror.Code.BUnexpectedBlankVariable, "addr is blank")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisStore{Client: client, KeyPrefix: prefix, MaxLifeSecond: maxLifeSecond}, nil
}

func (r *RedisStore) SessionInit(ctx context.Context, sid string) (Session, error) {
	if r.Client == nil {
		return nil, xerror.NewError(nil, xerror.Code.BUnexpectedBlankVariable, "client is nil")
	}

	return &RedisSession{
		Sid:   sid,
		Value: nil,
		Store: r,
	}, nil
}

func (r *RedisStore) SessionRead(ctx context.Context, sid string) (Session, error) {
	if r.Client == nil {
		return nil, xerror.NewError(nil, xerror.Code.BUnexpectedBlankVariable, "client is nil")
	}

	key := fmt.Sprintf("%s:%s", r.KeyPrefix, sid)

	cmd := r.Client.Get(ctx, key)
	valueRaw, err := cmd.Result()

	if err != nil {
		return nil, xerror.NewError(nil, xerror.Code.SRedisExecuteErr, "get value failed")
	}

	var value map[string]interface{}
	err = json.Unmarshal([]byte(valueRaw), &value)
	if err != nil {
		return nil, xerror.NewError(err, xerror.Code.SUnexpectedErr, "redis json unmarshal failed")
	}

	return &RedisSession{
		Sid:   sid,
		Value: value,
	}, nil
}

func (r *RedisStore) SessionDestroy(ctx context.Context, sid string) error {
	if r.Client == nil {
		return xerror.NewError(nil, xerror.Code.BUnexpectedBlankVariable, "client is nil")
	}

	key := fmt.Sprintf("%s:%s", r.KeyPrefix, sid)

	_, err := r.Client.Del(ctx, key).Result()
	if err != nil {
		return xerror.NewError(err, xerror.Code.SRedisExecuteErr, err.Error())
	}

	return nil
}

func (r *RedisStore) SessionWrite(ctx context.Context, sid string, value interface{}) error {
	if r.Client == nil {
		return xerror.NewError(nil, xerror.Code.BUnexpectedBlankVariable, "client is nil")
	}

	key := fmt.Sprintf("%s:%s", r.KeyPrefix, sid)

	bt, err := json.Marshal(value)
	if err != nil {
		return xerror.NewError(nil, xerror.Code.SUnexpectedErr, "json unmarshal failed")
	}

	_, err = r.Client.Set(ctx, key, string(bt), time.Duration(r.MaxLifeSecond)*time.Second).Result()
	if err != nil {
		return xerror.NewError(err, xerror.Code.SRedisExecuteErr, "redis set failed")
	}

	return nil
}

func (r *RedisStore) SessionGC(ctx context.Context, maxLifeTime int64) {
	return
}
