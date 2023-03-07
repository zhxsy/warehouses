package send

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"log"
	"sync"
	"time"
)

type TgService struct {
	Token  string `json:"token"`
	ChatID int64  `json:"chatID"`
	Env    string `json:"env"`
}

func GetTgService(token string, chatId int64, env string) *TgService {
	return &TgService{
		Token:  token,
		ChatID: chatId,
		Env:    env,
	}
}

// 发送文件
func (tg *TgService) SendDocument(filePath string) error {
	bot, err := tgbotapi.NewBotAPI(tg.Token)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewDocument(tg.ChatID, tgbotapi.FilePath(filePath))
	_, err = bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

// 发送消息
func (tg *TgService) Send(msg string) error {
	bot, err := tgbotapi.NewBotAPI(tg.Token)
	if err != nil {
		return err
	}
	config := tgbotapi.MessageConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           tg.ChatID,
			ReplyToMessageID: 0,
		},
		Text:                  msg,
		DisableWebPagePreview: false,
		ParseMode:             tgbotapi.ModeMarkdown,
	}

	_, err = bot.Send(config)
	if err != nil {
		content, _ := json.Marshal(config)
		log.Printf("TgServiceSend:%s", string(content))
		return err
	}
	return nil
}

// 发送消息
func (tg *TgService) SendMessage(msg string) error {
	bot, err := tgbotapi.NewBotAPI(tg.Token)
	if err != nil {
		return err
	}
	//entry, _ := MakeLogWarn(msg, errAlarm, map[string]interface{}{"env":tg.Env})
	//text, _ := entry.String()
	config := tgbotapi.NewMessage(tg.ChatID, msg)
	_, err = bot.Send(config)
	if err != nil {
		return err
	}
	return nil
}
func (tg *TgService) SendLog(msg string, errAlarm error) error {
	//entry, _ := MakeLogWarn(msg, errAlarm, map[string]interface{}{})
	//text, _ := entry.String()
	//text = fmt.Sprintf("*env:*%s;\n*msg:* %s", tg.Env, text)
	entry, _ := MakeLogWarn(msg, errAlarm, map[string]interface{}{"env": tg.Env})
	text, _ := entry.String()
	return tg.SendMessage(text)
}

type logTg struct {
	log  *logrus.Logger
	Once sync.Once
}

var logTgServ = &logTg{}

func GetLogTg() *logTg {
	logTgServ.Once.Do(func() {
		log := logrus.New()
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			DataKey:         "fields",
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyFunc:  "caller",
				logrus.FieldKeyMsg:   "message",
				logrus.FieldKeyTime:  "time",
				logrus.FieldKeyLevel: "level",
			},
		})
		logTgServ.log = log
	})
	return logTgServ
}
func MakeLogWarn(msg string, err interface{}, data map[string]interface{}) (*logrus.Entry, error) {
	s := logrus.Fields{
		"err": err,
	}
	if len(data) > 0 {
		for k, v := range data {
			s[k] = v
		}
	}
	l := GetLogTg().log.WithFields(s)
	l.Level = logrus.WarnLevel
	l.Time = time.Now()
	l.Message = msg

	return l, nil
}
