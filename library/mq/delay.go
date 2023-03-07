package mq

// DelayWrapper 通过 DelayWrapper 做一层封装，不引入其他任何依赖
type DelayWrapper interface {
	SetDelayMsg(msg *DelayMsg) error
	GetDelayMsg() ([]*DelayMsg, error)
}

type DelayMsg struct {
	Name      string
	Key       string
	Msg       interface{}
	Delay     int64
	Timestamp int64
}

func (d *DelayMsg) GetExchangeName() string {
	return d.Name
}

func (d *DelayMsg) GetKey() string {
	return d.Key
}

func (d *DelayMsg) GetMsg() interface{} {
	return d.Msg
}

func (d *DelayMsg) GetDelay() int64 {
	return d.Delay
}
