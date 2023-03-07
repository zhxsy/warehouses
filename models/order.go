package models

type OrderType uint8

const (
	OrderTypeUnknown OrderType = iota
	OrderTypeDecomposition
	OrderTypeIntensify
	OrderTypeRoll
)

type Order struct {
	OrderId string    `json:"order_id"`
	EquipId uint64    `json:"equip_id"`
	Type    OrderType `json:"type"`
}
