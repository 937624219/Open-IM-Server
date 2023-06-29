package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"testing"
)

import (
	"context"
	"time"

	"github.com/OpenIMSDK/Open-IM-Server/pkg/common/mw/specialerror"
	"github.com/OpenIMSDK/Open-IM-Server/pkg/errs"
	"github.com/redis/go-redis/v9"
)

func NewRedis() (redis.UniversalClient, error) {
	if len(Config.Redis.Address) == 0 {
		return nil, errors.New("redis address is empty")
	}
	specialerror.AddReplace(redis.Nil, errs.ErrRecordNotFound)
	var rdb redis.UniversalClient
	if len(Config.Redis.Address) > 1 {
		rdb = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    Config.Redis.Address,
			Username: Config.Redis.Username,
			Password: Config.Redis.Password, // no password set
			PoolSize: 50,
		})
	} else {
		rdb = redis.NewClient(&redis.Options{
			Addr:     Config.Redis.Address[0],
			Username: Config.Redis.Username,
			Password: Config.Redis.Password, // no password set
			DB:       0,                     // use default DB
			PoolSize: 100,                   // 连接池大小
		})
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	err := rdb.Ping(ctx).Err()
	if err != nil {
		return nil, fmt.Errorf("redis ping %w", err)
	}
	return rdb, nil
}

func TestName(t *testing.T) {
	name := "./../../../config/config.yaml"
	fmt.Println(filepath.Abs(name))
	data, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	if err := yaml.NewDecoder(bytes.NewBuffer(data)).Decode(&Config); err != nil {
		panic(err)
	}
	jsonData, err := json.Marshal(Config)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
	fmt.Println(string(EncodeConfig()))

}
