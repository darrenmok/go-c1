package visitor

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"

	"github.com/darrenmok/go-c1/cache"
)

var (
	redisC *cache.Client
)

const (
	key string = "visitor_count"
)

func init() {
	var err error
	redisC, err = cache.NewClient(context.Background(), os.Getenv("REDIS_URL"))
	if err != nil {
		log.Fatalf("%v", err)
	}
}

// Get will display a page showing the current vistor count
func Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	count, err := getVisitorCount(ctx)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(fmt.Sprintf("This is the %d visitor", count)))
	incrementVisitorCount(ctx)
}

func getVisitorCount(ctx context.Context) (int64, error) {
	val, err := redisC.GetInt64(ctx, key)
	if err == redis.Nil {
		redisC.SetInt64(ctx, key, 1, 0)
		return 1, nil
	}
	if err != nil {
		return 0, err
	}

	return val, nil
}

func incrementVisitorCount(ctx context.Context) error {
	curr, err := getVisitorCount(ctx)
	if err != nil {
		return err
	}
	return redisC.SetInt64(ctx, key, curr+1, 0)
}
