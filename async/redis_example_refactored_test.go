package async

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

func TestSetGet(t *testing.T) {
	rdb := GetClient()
	val := SetGet(rdb)
	if val != "1" {
		t.Errorf("val should be 1")
	}
}

type mockClient struct{}

func (m mockClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return redis.NewStatusCmd(ctx, key, "123", expiration)
}

func (m mockClient) Get(ctx context.Context, key string) *redis.StringCmd {
	r := redis.NewStringCmd(ctx, "123")
	r.SetVal("123")
	return r
}

func TestSetGetWithMock(t *testing.T) {
	rdb := mockClient{}
	val := SetGet(rdb)
	if val != "123" {
		t.Errorf("val should be 123 but got %v", val)
	}
}
