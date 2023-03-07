package app

import (
	"errors"
	"fmt"
)

var (
	chain = ""
	// 所有可用的运行时环境列表
	chainList = []string{
		ChainAvax,
		ChainMAvax,
		ChainKlaytn,
		ChainEthereum,
		ChainBnb,
		ChainMBnb,
		ChainMatic,
		ChainOptimism,
		ChainArbitrum,
		ChainFantom,
	}
)

// 初始化链
func InitChain(c string) {
	eMutex.Lock()
	defer eMutex.Unlock()
	if c != "" {
		if !isValidChain(c) {
			panic(errors.New(fmt.Sprintf("init_chain_error,env: %s is invalid", c)))
		}
		chain = c
	}
	if chain == "" {
		panic("chain not is empty")
	}
}

func GetChain() string {
	return chain
}

// 确定选择的链是否有效
func isValidChain(chain string) bool {
	for i, j := 0, len(chainList); i < j; i++ {
		if chain == chainList[i] {
			return true
		}
	}
	return false
}

// 队列按照链区分
func GetMq(k string) string {
	return fmt.Sprintf(k, GetChain())
}
