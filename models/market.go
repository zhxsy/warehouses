package models

type GenerateHorse struct {
	BoxType int32 `json:"box_type"`
	UserId  int64 `json:"user_id"`
	BoxId   int64 `json:"box_id"`
	ChainId int64 `json:"chain_id"`
}

type MqSellingHorseInfo struct {
	OrderId string `json:"order_id"`
}
