package app

import (
	"context"
	"github.com/go-redis/redis/v8"
	"sync"
)

var redisManager = CreateRedisClientManager()

func InitRedis() {
	configs := make(map[string]*RedisConfig)

	err := Config("redis").Unmarshal(&configs)
	if err != nil {
		Log().WithError(err).Error("redis_init_error")
	}

	for n, c := range configs {
		log := Log().WithField("name", n).WithField("config", c)
		if err := redisManager.InitRedis(n, c); err != nil {
			log.WithError(err).Error("redis_init_error")
			panic(err)
		}
		log.Info("redis_init_success")
	}
}

func Redis(name string) redis.UniversalClient {
	if redisEngine := redisManager.Get(name); redisEngine != nil {
		return redisEngine
	}
	Log().WithField("name", name).Error("redis_not_init")
	return nil
}

// ----------------------------------------
//  Redis 客户端管理器
// ----------------------------------------

type ClientManager struct {
	rw      *sync.RWMutex
	clients map[string]redis.UniversalClient
}

func CreateRedisClientManager() *ClientManager {
	return &ClientManager{rw: &sync.RWMutex{}, clients: make(map[string]redis.UniversalClient)}
}

func (manager *ClientManager) InitRedis(name string, config *RedisConfig) error {
	//构建通用redis连接
	redisEngine := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:        config.Addrs,
		DB:           config.DB,
		Password:     config.Password,
		MaxRetries:   config.MaxRetries,
		PoolSize:     config.PoolSize,
		MinIdleConns: config.MinIdleConns,
		MasterName:   config.MasterName,
		OnConnect: func(ctx context.Context, cn *redis.Conn) error {
			return nil
		},
	})

	//测试连接
	ctx := context.Background()
	err := redisEngine.Ping(ctx).Err()
	if err != nil {
		return err
	}

	//保存到manager
	manager.Set(name, redisEngine)
	return nil
}

func (manager *ClientManager) Get(name string) redis.UniversalClient {
	manager.rw.RLock()
	defer manager.rw.RUnlock()

	if client, exists := manager.clients[name]; exists {
		return client
	} else {
		return nil
	}
}

func (manager *ClientManager) Set(name string, client redis.UniversalClient) {
	manager.rw.Lock()
	defer manager.rw.Unlock()

	manager.clients[name] = client
}

// ----------------------------------------
//  Redis 配置项
// ----------------------------------------

type RedisConfig struct {
	Addrs        []string `json:"addrs"`
	DB           int      `json:"db"`
	Password     string   `json:"password"`
	MaxRetries   int      `json:"max_retries" mapstructure:"max_retries"`
	PoolSize     int      `json:"pool_size" mapstructure:"pool_size"`
	MinIdleConns int      `json:"min_idle_conns" mapstructure:"min_idle_conns"`
	MasterName   string   `json:"master_name" mapstructure:"master_name"`
}
