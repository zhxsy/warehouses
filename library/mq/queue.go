package mq

/*
所有启动的队列、交换机、routing 配置项
*/
const Qos = 0 // 每次推送次数 0 全部推送

// 交互机
const (
	ExchangeBlock      = "block"
	ExchangeAlarm      = "alarm"
	ExchangeTransfer   = "transfer"
	DelayExchangeOrder = "order"

	MQExchangeRace                        = "race"
	MQDelayExchangeRaceSettlement         = "race_settlement"
	MQExchangeFreeHorseToCoupon           = "exchange_free_horse_to_coupon"
	MQDelayExchangeSellingHorse           = "exchange_horse_selling"
	MQExchangeTicketFeeRefund             = "exhange_ticket_fee_refund"
	MQExchangeGenerateHorse               = "exchange_generate_horse"
	MQDelayExchangeDeathRaceRoom          = "death_race_room"
	MQDelayExchangeDeathRaceRoomLock      = "death_race_room_lock"
	MQDelayExchangeDeathRaceSettle        = "death_race_settle"
	MQDelayExchangeTournamentRaceRoom     = "tournament_race_room"
	MQDelayExchangeTournamentRaceRoomLock = "tournament_race_room_lock"
	MQDelayExchangeTournamentRaceSettle   = "tournament_race_settle"
	MQExchangeJoinRoom                    = "exchange_join_room"
	ExchangeMessage                       = "message"
	MQDelayExchangeScheduleRoom           = "schedule_room"
	MQDelayExchangeScheduleRaceSettlement = "schedule_race_settlement"
	MQExchangePrizePool                   = "prize_pool"
	ExchangeMiniTransferAvax              = "mini_transfer_avax"
	ExchangeMiniTransferKlay              = "mini_transfer_klay"
)

// 路由
const (
	BlockRoutingKey = "%v.%v.%v" // block_logs:hoof:0xddf
)

const (
	DeathRaceRoutingKey               = "death_race_join_room"
	BroadcastStationMessageRoutingKey = "message.station.broadcast"
	SystemStationMessageRoutingKey    = "message.station.system"
	AlarmLogRoutingKey                = "alarm_log"
	OrderTimeoutClose                 = "order.timeout.#"

	MQRaceRoutingKey              = "%v:%v" // raceID:timestamp
	MQRaceSettlementRoutingKey    = "%v:%v" // raceID:timestamp
	MQSellingHorseRoutingKey      = "%v_%v" // horseOwnerAddr:HorseId
	MQRaceDeathRoomRoutingKey     = "%v:%v" // roomId:timestamp
	MQRaceDeathRoomLockRoutingKey = "%v:%v" // roomId:timestamp
	MQRaceDeathRoutingKey         = "%v:%v" // raceId:timestamp
	MQRaceTournamentRoutingKey    = "%v:%v" // raceId:timestamp
	MQPrizePoolRoutingKey         = "%v:%v" // raceId:timestamp

	// 监控链上路由
	HoofLogs                = "%s.hoof.#"
	BurnLogs                = "%s.burn.#"
	EquipLogs               = "%s.equip.#"
	TransferLogs            = "%s.transfer.#"
	HorseLogs               = "%s.horse.#"
	EquipTicketTransferLogs = "%s.equip_ticket_transfer.#"
	MintHoofLogs            = "%s.mint_hoof.#"
	SeHorseLogs             = "%s.se_horse.#"
	// 链上路由
	TransferCheckRoutingKey         = "%s.transfer_check"
	TransferCheckRoutingKeyMiniGame = "%s.transfer_check_minigame"
)

// 队列
const (
	BlockLogs     = "avax"
	BlockKlaytn   = "klaytn"
	BlockEthereum = "ethereum"

	QueueAlarm                  = "alarm_log"                 // 告警日志
	QueueDeathRaceRoom          = "death_race_room"           // 死亡比赛房间
	QueueDeathRaceRoomLock      = "death_race_room_lock"      // 死亡比赛房间锁定
	QueueDeathRaceSettle        = "death_race_settle"         // 死亡比赛结束
	QueueDeathRaceCareer        = "death_race_career"         // 死亡赛生涯
	QueueDeathRaceJoinRoom      = "death_race_join_room"      // 死亡比赛加入
	QueueTournamentRaceRoom     = "tournament_race_room"      // 锦标赛比赛房间
	QueueTournamentRaceRoomLock = "tournament_race_room_lock" // 锦标赛比赛房间锁定
	QueueTournamentRaceSettle   = "tournament_race_settle"    // 锦标赛比赛结束

	QueueOrderTimeout                  = "order_timeout" // 订单超时
	QueueBoxLogsQueueAlarm             = "box_logs"      // 盲盒
	QueueScheduleRoom                  = "schedule_room"
	QueueScheduleFee                   = "schedule_fee"
	QueueScheduleRoomSettle            = "schedule_room_settle"
	QueueScheduleHorseSettle           = "schedule_horse_settle"
	QueueScheduleRaceSettle            = "schedule_race_settle"
	QueueScheduleRaceHorseCareerSettle = "schedule_race_horse_career_settle" // 定时比赛结束
	QueueScheduleRaceAward             = "schedule_race_award"

	QueueBroadcastStationMessage = "send_broadcast_station_message" // 发送公告消息
	QueueSystemStationMessage    = "send_system_station_message"    // 发送系统消息

	MQQueueRunGame                     = "run_game"
	MQQueueHoofSettle                  = "hoof_settle"
	MQQueuePointSettle                 = "point_settle"
	MQQueueHorseStatusSettle           = "horse_status_settle"
	MQQueueRaceStatusSettle            = "race_status_settle"
	MQQueueHorseHashCache              = "horse_hash_cache"
	MQQueueFreeHorseToCouponWhenSettle = "free_horse_to_coupon_settle"
	MQQueueUpdateHorseCareer           = "UpdateHorseCareer"
	MQQueueSellingHorse                = "queue_horse_selling"
	MQQueueTicketFeeRefund             = "ticket_fee_refund"
	MQQueueScheduleFeeRefund           = "schedule_fee_refund"
	MQQueueMintHorseNewContractCheck   = "queue_mint_horse_contract_check"
	MQQueueGenerateHorse               = "queue_generate_horse"
	MQQueueFreeHorseCoupon             = "queue_free_horse_coupon"

	MQQueuePrizePoolChange = "prize_pool_change"

	// 监控链队列
	BoxLogs                   = "%s_box_logs"                   // 盲盒
	RenameHorseBurnLogs       = "%s_rename_horse_burn_logs"     // 改名字
	EquipOrderLogs            = "%s_equip_order_logs"           // 装备
	TicketFeeBurnLogs         = "%s_ticket_fee_burn_logs"       // 门票
	TransferHorseLogs         = "%s_transfer_horse_logs"        // horse 交易
	MintHorseLogs             = "%s_mint_horse_logs"            // mint horse
	SeMintHorseLogs           = "%s_se_mint_horse_logs"         // new mint horse
	SeEquipTicketTransferLogs = "%s_equip_ticket_transfer_logs" //  equip ticket transfer
	MintHooFLogs              = "%s_mint_hoof_logs"             // mint hoof
	SeMintEquipLogs           = "%s_mint_equip_logs"            // mint equip
	ScheduleFee               = "%s_schedule_fee"

	// 链上队列
	QueueTransferCheck    = "%s_transfer_check" // 转账校验
	GenerateNftRoutingKey = "%s_generate_nft"   // horse gear重新mint
)
