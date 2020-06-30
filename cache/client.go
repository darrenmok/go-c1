package cache

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

// Client struct holds the redis client
type Client struct {
	client *redis.Client
}

// NewClient connects to redis and returns a client to access Redis
func NewClient(ctx context.Context, redisURL string) (*Client, error) {
	if redisURL == "" {
		redisURL = os.Getenv("REDIS_URL")
	}
	rURL, err := url.Parse(redisURL)
	if err != nil {
		return nil, err
	}

	redisPw, _ := rURL.User.Password()

	if rURL.Host == "" {
		return nil, fmt.Errorf("Redis URL incorrect")
	}

	opt := &redis.Options{
		Addr:       rURL.Host,
		Password:   redisPw,
		DB:         0,
		MaxRetries: 3,
	}

	if rURL.Scheme == "rediss" {
		opt.TLSConfig = &tls.Config{
			MinVersion:         0,
			InsecureSkipVerify: true,
		}
	}

	client := redis.NewClient(opt)

	err = client.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
	}, nil
}

// SetInt64 sets a key/int pair which expires in Redis
// 0 expiration means no expiration
func (c *Client) SetInt64(ctx context.Context, key string, value int64, expiration time.Duration) error {
	err := c.client.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

// GetInt64 gets the int by the key
func (c *Client) GetInt64(ctx context.Context, key string) (int64, error) {
	val, err := c.client.Get(ctx, key).Int64()
	if err != nil {
		return 0, err
	}
	return val, nil
}
