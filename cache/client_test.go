package cache_test

import (
	"context"
	"testing"
	"time"

	"github.com/darrenmok/go-c1/cache"
)

func TestClient(t *testing.T) {
	testCases := []struct {
		key string
		val int64
	}{
		{key: "A", val: 1},
		{key: "B", val: 2},
	}
	for _, tc := range testCases {
		t.Run(tc.key, func(t *testing.T) {
			ctx := context.Background()
			client, err := cache.NewClient(ctx, "")
			if err != nil {
				t.Errorf("NewClient error: %v", err)
			}

			err = client.SetInt64(ctx, tc.key, tc.val, 10*time.Second)
			if err != nil {
				t.Errorf("Set key val error: %v", err)
			}

			val, err := client.GetInt64(ctx, tc.key)
			if err != nil {
				t.Errorf("Get key error: %v", err)
			}

			if val != tc.val {
				t.Errorf("Want %d Got %d", tc.val, val)
			}
		})
	}
}

func TestGet_Error(t *testing.T) {
	client, err := cache.NewClient(context.Background(), "")
	if err != nil {
		t.Errorf("NewClient error: %v", err)
	}

	_, err = client.GetInt64(context.Background(), "no_such_key")
	if err == nil {
		t.Errorf("Require error %v", err)
	}
}
