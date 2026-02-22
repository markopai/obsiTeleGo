package redisRepo

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RedisRepo struct {
	db  *redis.Client
	Log *slog.Logger
}

func New(client *redis.Client, log *slog.Logger) *RedisRepo {
	return &RedisRepo{
		db:  client,
		Log: log,
	}
}

func (r *RedisRepo) GetThreadName(ctx context.Context, threadID int64) (string, error) {
	val, err := r.db.Get(ctx, strconv.FormatInt(threadID, 10)).Result()

	if err == redis.Nil {
		return "", fmt.Errorf("no key redis error: %w", err)
	} else if err != nil {
		return "", fmt.Errorf("get redis error: %w", err)
	}

	return val, nil
}
func (r *RedisRepo) NewThread(ctx context.Context, threadID int64, text string) error {
	err := r.db.Set(ctx, strconv.FormatInt(threadID, 10), text, 0).Err()

	if err != nil {

		return fmt.Errorf("set redis error: %w", err)
	}

	return nil
}
