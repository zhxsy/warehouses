package app

//
//var (
//	RpcClient         = &Rpc{}
//	TicketRpcClient   = &TicketRpc{}
//	WalletRpcClient   = &WalletRpc{}
//	T4fRpcClient      = &T4fRpc{}
//	MiniGameRpcClient = &MiniGameRpc{}
//)
//
//type Rpc struct {
//	Api               api.ApiClient
//	MetaderbyConfig   metaderby.ApiClient
//	MetaderbyBlock    metaderby.BlockApiClient
//	MetaderbyMq       metaderby.MqApiClient
//	Horse             metaderby.HorseApiClient
//	Open              metaderby.OpenApiClient
//	Schedule          metaderby.ScheduleApiClient
//	MetaderbyStat     metaderby.StatisticsApiClient
//	MetaderbyRoom     metaderby.RoomRpcClient
//	MetaderbyActivity metaderby.ActivityRpcClient
//	Reconciliation    metaderby.ReconciliationApiClient
//	SendTelegramInfo  metaderby.AlarmApiClient
//	Order             metaderby.OrderApiClient
//	Message           metaderby.MessageRpcClient
//	Market            metaderby.MarketRpcClient
//	Once              sync.Once
//}
//
//type TicketRpc struct {
//	Ticket metaderby.TicketRpcClient
//	Once   sync.Once
//}
//
//type WalletRpc struct {
//	Order  wallet.OrderRpcClient
//	Api    wallet.ApiClient
//	Mpc    wallet.MpcRpcClient
//	Oauth  wallet.OauthRpcClient
//	Assets wallet.AssetsClient
//	Cmd    wallet.WalletCmdRpcClient
//	Once   sync.Once
//}
//
//type T4fRpc struct {
//	Schedule t4f.ScheduleRpcClient
//	Once     sync.Once
//}
//
//type MiniGameRpc struct {
//	MiniGameConfig metaderby.ApiClient
//	MiniGameBlock  metaderby.BlockApiClient
//	MiniGameMq     metaderby.MqApiClient
//	MiniGameUser   metaderby.MiniApiClient
//	Once           sync.Once
//}
//
//func InitRpc(host, port string) {
//	RpcClient.Once.Do(func() {
//		conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
//		if err != nil {
//			log.Fatalf("did not connect: %v", err)
//		}
//		RpcClient = &Rpc{
//			Api:               api.NewApiClient(conn),
//			MetaderbyConfig:   metaderby.NewApiClient(conn),
//			MetaderbyBlock:    metaderby.NewBlockApiClient(conn),
//			MetaderbyMq:       metaderby.NewMqApiClient(conn),
//			Horse:             metaderby.NewHorseApiClient(conn),
//			Open:              metaderby.NewOpenApiClient(conn),
//			Schedule:          metaderby.NewScheduleApiClient(conn),
//			MetaderbyStat:     metaderby.NewStatisticsApiClient(conn),
//			MetaderbyRoom:     metaderby.NewRoomRpcClient(conn),
//			MetaderbyActivity: metaderby.NewActivityRpcClient(conn),
//			Reconciliation:    metaderby.NewReconciliationApiClient(conn),
//			SendTelegramInfo:  metaderby.NewAlarmApiClient(conn),
//			Order:             metaderby.NewOrderApiClient(conn),
//			Message:           metaderby.NewMessageRpcClient(conn),
//			Market:            metaderby.NewMarketRpcClient(conn),
//		}
//	})
//}
//
//func InitTicket(host, port string) {
//	TicketRpcClient.Once.Do(func() {
//		_, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
//		if err != nil {
//			log.Fatalf("did not connect: %v", err)
//		}
//
//		TicketRpcClient = &TicketRpc{}
//	})
//}
//
//func InitWallet(host, port string) {
//	WalletRpcClient.Once.Do(func() {
//		conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
//		if err != nil {
//			log.Fatalf("did not connect: %v", err)
//		}
//
//		WalletRpcClient = &WalletRpc{
//			Order:  wallet.NewOrderRpcClient(conn),
//			Api:    wallet.NewApiClient(conn),
//			Mpc:    wallet.NewMpcRpcClient(conn),
//			Oauth:  wallet.NewOauthRpcClient(conn),
//			Assets: wallet.NewAssetsClient(conn),
//			Cmd:    wallet.NewWalletCmdRpcClient(conn),
//		}
//	})
//}
//
//func InitT4f(host, port string) {
//	T4fRpcClient.Once.Do(func() {
//		conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
//		if err != nil {
//			log.Fatalf("did not connect: %v", err)
//		}
//		T4fRpcClient = &T4fRpc{
//			Schedule: t4f.NewScheduleRpcClient(conn),
//		}
//	})
//}
//
//func InitMiniGameRpc(host, port string) {
//	RpcClient.Once.Do(func() {
//		conn, err := grpc.Dial(fmt.Sprintf("%s:%s", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
//		if err != nil {
//			log.Fatalf("minigameRpc did not connect: %v", err)
//		}
//		MiniGameRpcClient = &MiniGameRpc{
//			MiniGameConfig: metaderby.NewApiClient(conn),
//			MiniGameBlock:  metaderby.NewBlockApiClient(conn),
//			MiniGameMq:     metaderby.NewMqApiClient(conn),
//			MiniGameUser:   metaderby.NewMiniApiClient(conn),
//		}
//	})
//}
