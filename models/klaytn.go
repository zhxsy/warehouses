package models

import (
	"math/big"
	"strings"
)

// 区块
type BlockKlaytn struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      int64  `json:"id"`
	Result  string `json:"result"`
}

// 日志
type BlockLogsKlay struct {
	Jsonrpc string                `json:"jsonrpc"`
	Id      int64                 `json:"id"`
	Result  []BlockLogsKlayResult `json:"result"`
}

type BlockLogsKlayResult struct {
	Address          string   `json:"address"`
	Topics           []string `json:"topics"`
	Data             string   `json:"data"`
	BlockNumber      string   `json:"blockNumber"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
	BlockHash        string   `json:"blockHash"`
	LogIndex         string   `json:"LogIndex"`
	Removed          bool     `json:"Removed"`
}

// 解析data 数据
func (b *BlockLogsKlayResult) ParseData(data string) *big.Int {
	order16 := "0x" + strings.TrimLeft(data, "0")
	orderStr, _ := new(big.Int).SetString(order16, 0)
	return orderStr
}
