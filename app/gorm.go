package app

import (
	"gorm.io/gorm"
	"sync"
)

var gormManager = CreateGORMClientManager()

func Gorm(name string) *gorm.DB {
	if dbEngine := gormManager.Get(name); dbEngine != nil {
		return dbEngine
	}
	Log().WithField("name", name).Error("gorm_not_init")
	return nil
}

// ----------------------------------------
//  GORM 客户端管理器
// ----------------------------------------

type GORMClientManager struct {
	rw      *sync.RWMutex
	clients map[string]*gorm.DB
}

// CreateGORMClientManager 创建 GORM 客户端管理器实例
func CreateGORMClientManager() *GORMClientManager {
	return &GORMClientManager{rw: &sync.RWMutex{}, clients: make(map[string]*gorm.DB)}
}

// Get 获取给定名称的 GORM 客户端实例（如果客户端不存在则返回 nil）
func (manager *GORMClientManager) Get(name string) *gorm.DB {
	manager.rw.RLock()
	defer manager.rw.RUnlock()
	if client, exists := manager.clients[name]; exists {
		return client
	} else {
		return nil
	}
}

// Set 添加或更新 GORM 客户端实例
func (manager *GORMClientManager) Set(name string, client *gorm.DB) {
	manager.rw.Lock()
	manager.clients[name] = client
	manager.rw.Unlock()
}
