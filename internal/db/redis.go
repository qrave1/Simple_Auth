package db

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"rchir7/internal/config"
	"time"
)

type RedisDb struct {
	db *redis.Client
}

func NewRedis(c *config.Config) *RedisDb {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", c.RdbAddress, c.RdbPort),
		Password: "",
		DB:       0,
	})

	//rdb.Set(context.Background(), "1", "11111111111111111", 60)
	//s, err := rdb.Get(context.Background(), "1").Result()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(s)

	return &RedisDb{rdb}
}

func (r *RedisDb) Insert(email, token string) error {
	ctx := context.Background()
	return r.db.Set(ctx, email, token, 24*time.Hour).Err()
}

func (r *RedisDb) Read(email string) (string, error) {
	ctx := context.Background()
	result, err := r.db.Get(ctx, email).Result()
	if err != nil {
		return "", err
	}

	return result, nil
}
