package models

type BlockLogsEthResult struct {
	Address     string   `json:"address"`
	Topics      []string `json:"topics"`
	Data        string   `json:"data"`
	BlockNumber uint64   `json:"blockNumber"`
	TxHash      string   `json:"tx_hash"`
	TxIndex     uint     `json:"tx_index"`
	BlockHash   string   `json:"blockHash"`
	Index       uint     `json:"index"`
	Removed     bool     `json:"Removed"`
}

type BlockLogsQuickResult struct {
	Address     string   `json:"address"`
	Topics      []string `json:"topics"`
	Data        string   `json:"data"`
	BlockNumber uint64   `json:"block_number"`
	TxHash      string   `json:"tx_hash"`
	TxIndex     uint     `json:"tx_index"`
	BlockHash   string   `json:"blockHash"`
	Index       uint     `json:"index"`
	Removed     bool     `json:"removed"`
	ChainId     uint64   `json:"chain_id"`
}
