package async

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var contx = context.Background()

type redisClient interface {
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
}

func GetClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func SetGet(rdb redisClient) string {
	err := rdb.Set(contx, "yo", "1", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(contx, "yo").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	return val
}
