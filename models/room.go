package models

type Room struct {
	RoomId uint64 `json:"room_id"`
}

type RaceSettlement struct {
	Body []byte `json:"body"`
}
