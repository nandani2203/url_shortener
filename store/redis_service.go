package store

import (
	"context"
	"time"

	"github.com/bits-and-blooms/bloom/v3" // New import
	"github.com/go-redis/redis/v8"
)

type StorageService struct {
	redisClient *redis.Client
	Filter      *bloom.BloomFilter // Add this
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

func InitializeStore() *StorageService {
	// Initialize Bloom Filter for 1M items with 1% false positive rate
	storeService.Filter = bloom.NewWithEstimates(1000000, 0.01)

	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	storeService.redisClient = redisClient
	return storeService
}

func SaveUrlMapping(shortUrl string, originalUrl string) {
	err := storeService.redisClient.Set(ctx, shortUrl, originalUrl, 24*time.Hour).Err()
	if err == nil {
		// ADD TO FILTER: Tell the filter this URL now exists
		storeService.Filter.Add([]byte(shortUrl))
	}
}

func RetrieveInitialUrl(shortUrl string) string {
	return storeService.redisClient.Get(ctx, shortUrl).Val()
}

// Add this helper to access the service globally
func GetStoreService() *StorageService {
	return storeService
}
