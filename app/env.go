package app

import (
	"errors"
	"fmt"
	"os"
	"sync"
)

var (
	eMutex  = new(sync.Mutex) // 当前运行时环境写锁
	current = ReadEnv()       // 当前运行时环境（默认为开发环境）

	// 所有可用的运行时环境列表
	environments = []string{
		DevEnvironment,
		StEnvironment,
		PrdEnvironment,
	}
)

func InitEnv(env string) {
	eMutex.Lock()
	defer eMutex.Unlock()
	if env != "" {
		if !isValidEnv(env) {
			panic(errors.New(fmt.Sprintf("init_env_error,env: %s is invalid", env)))
		}

		current = env

		fmt.Printf("init_env_success, env: %s\n", env)
	}

}

// 获取当前运行环境
func GetEnv() string {
	return current
}

//判断当前环境
func IsDev() bool { return current == DevEnvironment }
func IsSt() bool  { return current == StEnvironment }
func IsPrd() bool { return current == PrdEnvironment }

// 确定给定的运行环境是否有效
func isValidEnv(env string) bool {
	for i, j := 0, len(environments); i < j; i++ {
		if env == environments[i] {
			return true
		}
	}
	return false
}

// 配置文件路径
func GetEnvPath(name string) (string, bool) {
	envPath := os.Getenv(name)
	if envPath == "" {
		return "", false
	}
	return envPath, true
}

// 读取环境
func ReadEnv() string {
	env := os.Getenv("META_ENV")
	if env == "" {
		return DevEnvironment
	}

	return env
}
