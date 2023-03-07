package delay

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cfx/warehouses/library/mq"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var controller *Controller

type Controller struct {
	Cli redis.UniversalClient
}

func NewDelayControl(cli redis.UniversalClient) *Controller {
	controller = &Controller{cli}
	return controller
}

const DelayKey = "mq_delay_key"

func (d *Controller) SetDelayMsg(msg *mq.DelayMsg) error {
	b, _ := json.Marshal(msg)
	log.Printf("zset delay: %+v", msg)
	return d.Cli.ZAdd(context.Background(), DelayKey, &redis.Z{
		Score:  float64(time.Now().UnixMilli() + msg.GetDelay() - 1000),
		Member: string(b),
	}).Err()
}

func (d *Controller) GetDelayMsg() ([]*mq.DelayMsg, error) {
	u, err := d.Cli.ZRangeByScore(context.Background(), DelayKey, &redis.ZRangeBy{
		Min:    "-inf",
		Max:    fmt.Sprintf("%d", time.Now().UnixMilli()),
		Offset: 0,
		Count:  10,
	}).Result()
	if err != nil {
		return nil, err
	}
	var msg []*mq.DelayMsg
	for i := range u {
		var m *mq.DelayMsg
		//if v, ok := u[i].Member.(string); ok {
		//	json.Unmarshal([]byte(v), &m)
		//	msg = append(msg, m)
		//}
		json.Unmarshal([]byte(u[i]), &m)
		msg = append(msg, m)
		d.Cli.ZRem(context.Background(), DelayKey, u[i]).Err()
	}
	return msg, nil
}
