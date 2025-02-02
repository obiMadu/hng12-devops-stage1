package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	numbersAPIBaseURL = getEnvOrDefault("NUMBERS_API_BASE_URL", "http://numbersapi.com")
	redisClient      *redis.Client
	ctx             = context.Background()
)

func init() {
	redisURL := os.Getenv("REDIS_URL")
	if redisURL != "" {
		opt, err := redis.ParseURL(redisURL)
		if err != nil {
			fmt.Printf("Failed to parse Redis URL: %v\n", err)
			return
		}
		redisClient = redis.NewClient(opt)
	}
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// GetNumberFact retrieves a math fact about a number from Numbers API
func GetNumberFact(number int) (string, error) {
	cacheKey := fmt.Sprintf("number_fact:%d", number)

	// Try to get from cache first
	if redisClient != nil {
		if cached, err := redisClient.Get(ctx, cacheKey).Result(); err == nil {
			return cached, nil
		}
	}

	// If not in cache, get from API
	url := fmt.Sprintf("%s/%d/math", numbersAPIBaseURL, number)
	
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to get fact: status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	fact := string(body)

	// Cache the result
	if redisClient != nil {
		redisClient.Set(ctx, cacheKey, fact, 24*time.Hour)
	}

	return fact, nil
}
