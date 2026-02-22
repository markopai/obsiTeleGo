//go:build redis

package app

import (
	"context"
	"fmt"
	"log/slog"
	"obsiTeleGo/internal/repository"
	"obsiTeleGo/internal/repository/redisRepo"
	"os"

	"github.com/redis/go-redis/v9"
)

func initRepo(log *slog.Logger) (repository.Repo, database, error) {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	repo := redisRepo.New(rdb, log)

	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, nil, fmt.Errorf("ping redis error: %w", err)
	}

	return repo, rdb, nil
}
