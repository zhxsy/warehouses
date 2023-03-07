package mq

import (
	"github.com/cfx/warehouses/library/send"
)

// 发送警告log
func AlarmLog(msg string, err string) {
	if GetConnection() == nil {
		return
	}
	GetConnection().Publish(
		ExchangeAlarm,
		QueueAlarm,
		&send.AlarmLogInfo{
			Msg: msg,
			Err: err,
		},
	)
}
