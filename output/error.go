package output

import (
	"fmt"
	"github.com/cfx/warehouses/app"
)

// 用户自定义错误，不需要可自行删除
var (
	ErrParams          = NewErr(-1, "params error", true)
	ErrParamsMissing   = NewErr(-2, "params missing", true)
	ErrNet             = NewErr(-3, "Network error", true)
	ErrParamBindFailed = NewErr(-4, "param bindd failed", true)
	ErrUpgrading       = NewErr(-5, "system is upgrading", true)
	ErrBan             = NewErr(-6, "user is ban", true)

	//基础组件相关
	ErrRedisCon = NewErr(-50, "redis error", true)

	ErrWrongSign = NewErr(-10001, "wrong sign", true)

	ErrWrongMonitorKey            = NewErr(-1000, "wrong monitor-key", true)
	ErrWrongToken                 = NewErr(-1001, "wrong token or token is expire", true)
	ErrWrongUser                  = NewErr(-1002, "wrong user or user is need register", true)
	ErrWrongUserWhiteList         = NewErr(-1003, "you are not allowed to attend white list error", true)
	ErrTokenIsUpdateInOtherPlaces = NewErr(-1014, "account is login at different places", true)
	ErrTokenIsExpired             = NewErr(-1015, "token is expired", true)
	ErrTokenIsFalselyUsed         = NewErr(-1016, "token is falsely used", true)
	ErrMissingCTXParams           = NewErr(-1002, "missing context-params", true)
	ErrNoSuchRacing               = NewErr(-1003, "no such racing", true)
	ErrNoPermission               = NewErr(-1004, "Sorry, you don't have permission", true)
	ErrFormValidation             = NewErr(-1005, "Form Validation err", true)
	ErrIdempotentFailed           = NewErr(-1020, "idempotent request, please try again later", true)
	ErrUnmarshalJsonErr           = NewErr(-1021, "json unmarshal failed", true)
	ErrConsumerBodyEmpty          = NewErr(-1021, "Consumer body is empty", true)
	ErrNotEnoughHoof              = NewErr(-2001, "not enough hoof", true)

	// redis 错误
	ErrRedisErr    = NewErr(-5100, "redis err", false)
	ErrRedisGet    = NewErr(-5101, "redis get err", false)
	ErrRedisSet    = NewErr(-5102, "redis set err", false)
	ErrRedisSetEX  = NewErr(-5103, "redis setEX err", false)
	ErrRedisExpire = NewErr(-5104, "redis Expire err", false)
	ErrRedisExists = NewErr(-5105, "redis Exists err", false)
	// mysql
	ErrMysqlErr    = NewErr(-5200, "db find err", false)
	ErrMysqlGet    = NewErr(-5201, "db get err", false)
	ErrMysqlStore  = NewErr(-5202, "db insert err", false)
	ErrMysqlUpdate = NewErr(-5203, "db update err", false)
	ErrMysqlDelete = NewErr(-5204, "db delete err", false)
	ErrMysqlCount  = NewErr(-5205, "db count err", false)
	ErrMysqlFind   = NewErr(-5206, "db find err", false)
	// 邮箱
	ErrAddEmail      = NewErr(-4101, "add email err", true)
	ErrAddEmailIpMax = NewErr(-4102, "IP  exceeded the upper limit", true)
	ErrAddEmailExist = NewErr(-4103, "email exist", true)

	// 白名单
	ErrCheckWhiteOvertime = NewErr(-6101, "check timeout", true)
	ErrSignature          = NewErr(-6102, "generate signature err", true)
	ErrCBCEncrypt         = NewErr(-6201, "AesCBCEncrypt error", false)

	// race
	ErrMatchFailed          = NewErr(-10001, "match failed", true)
	ErrRunRaceFailed        = NewErr(-10002, "run race failed", false)
	ErrWrongRaceFeeOrder    = NewErr(-10003, "wrong race fee order", true)
	ErrWrongUserOfOrder     = NewErr(-10004, "order does not match user", true)
	ErrTicketOrderCancel    = NewErr(-10005, "ticket order cancel", true)
	ErrTicketFeeRefund      = NewErr(-10006, "match failed, ticket fee was refunded!", true)
	ErrCheckHorsePre        = NewErr(-10007, "pre check horse failed, room id needed!", true)
	ErrCheckHorsePreHorseId = NewErr(-10008, "you not has permissions!", true)

	// ticket
	ErrTicketNft               = NewErr(-7100, "Please select the corresponding NFT", true)
	ErrTicketNftApproved       = NewErr(-7101, "Please authorize first", true)
	ErrTicketNftHasReplaced    = NewErr(-7101, "Sorry，NFT has been replaced", true)
	ErrTicketNftReplacedFailed = NewErr(-7102, "Sorry, the exchange failed. Please try again later", true)

	// Validated fail
	ErrFieldNotAllowedEmpty = NewErr(-8000, "This Field should be not empty", true)

	// person center
	ErrUsernameExist   = NewErr(-8100, "username has exist", true)
	ErrUsernameTooLong = NewErr(-8101, "username too long", true)
	ErrEmailTooLong    = NewErr(-8102, "email too long", true)
)

type Error struct {
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
	ShowErr bool   `json:"show_err"`
	error
}

func NewErr(code int, msg string, showErr bool) *Error {
	err := &Error{
		Code:    code,
		Msg:     msg,
		ShowErr: showErr,
	}
	return err
}

func (e *Error) Error() string {
	return fmt.Sprintf("code=%d, msg=%s, show_err=%t", e.Code, e.Msg, e.ShowErr)
}

func (e *Error) SetMsg(msg string) *Error {
	return &Error{
		Code:    e.Code,
		Msg:     e.Msg + ": " + msg,
		ShowErr: e.ShowErr,
		error:   e.error,
	}
}

// 错误记录
func (e *Error) Log(err error) *Error {
	app.Log().WithField("msg", e.Msg).Warn(err)
	return e
}

func CheckErr(err error, name string) bool {
	if err != nil {
		panic(err)
		return false
	}
	return true
}
