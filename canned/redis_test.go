package canned

import (
	"context"
	"testing"

	testcontainers "github.com/testcontainers/testcontainers-go"
)

func TestSetInRedis(t *testing.T) {

	ctx := context.Background()

	c, err := NewRedisContainer(ctx, RedisContainerRequest{
		GenericContainerRequest: testcontainers.GenericContainerRequest{
			Started: true,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer c.Container.Terminate(ctx)

	redisClient, err := c.GetClient(ctx)
	if err != nil {
		t.Fatal(err)
	}

	err = redisClient.Set("key", "value", 0)
	if err != nil {
		t.Fatal(err)
	}
}
