package app

func InitMq() {
	var mqURL string
	Config("app").UnmarshalKey("amqp", &mqURL)
	if mqURL == "" {
		panic("mq init failed")
	}
}
