package models

import (
	"math/big"
	"strings"
)

type BlockHandle interface {
	ParseData(data string) *big.Int
}

// 区块
type BlockNumber struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

// 日志
type BlockLogs struct {
	Status  string            `json:"status"`
	Message string            `json:"message"`
	Result  []BlockLogsResult `json:"result"`
}

type BlockLogsResult struct {
	Address          string   `json:"address"`
	Topics           []string `json:"topics"`
	Data             string   `json:"data"`
	BlockNumber      string   `json:"blockNumber"`
	TimeStamp        string   `json:"timeStamp"`
	GasPrice         string   `json:"gasPrice"`
	GasUsed          string   `json:"gasUsed"`
	LogIndex         string   `json:"logIndex"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
}

type CheckTransaction struct {
	Status  string               `json:"status"`
	Message string               `json:"message"`
	Result  CheckTransactionInfo `json:"result"`
}

type CheckTransactionInfo struct {
	Status string `json:"status"`
}

// 解析data 数据
func (b *BlockLogsResult) ParseData(data string) *big.Int {
	order16 := "0x" + strings.TrimLeft(data, "0")
	orderStr, _ := new(big.Int).SetString(order16, 0)
	return orderStr
}

type KlaytnGetTransactionReceiptLogs struct {
	Jsonrpc string                     `json:"jsonrpc"`
	Id      int64                      `json:"id"`
	Result  KlaytnCheckTransactionInfo `json:"result"`
}

type KlaytnCheckTransactionInfo struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	To               string `json:"to"`
	GasPrice         string `json:"gasPrice"`
	GasUsed          string `json:"gasUsed"`
	TransactionHash  string `json:"transactionHash"`
	TransactionIndex string `json:"transactionIndex"`
	Status           string `json:"status"`
}
