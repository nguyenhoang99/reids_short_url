package store

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type StorageService struct {
	RedisClient *redis.Client
}

const CacheDuration = 6 * time.Minute

func InitializeStore() *StorageService {
	storeService := &StorageService{}
	ctx := context.Background()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}
	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	storeService.RedisClient = redisClient
	return storeService
}

func (s *StorageService) SaveUrlMapping(ctx context.Context, shortUrl string, originalUrl string) {
	err := s.RedisClient.Set(ctx, shortUrl, originalUrl, CacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\\n", err, shortUrl, originalUrl))
	}
}

func (s *StorageService) RetrieveInitialUrl(ctx context.Context, shortUrl string) string {
	result, err := s.RedisClient.Get(ctx, shortUrl).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\\n", err, shortUrl))
	}
	return result
}
