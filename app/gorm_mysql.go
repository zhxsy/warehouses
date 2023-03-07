package app

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

func InitGorm() {
	configs := make(map[string]*ORMConfig)

	err := Config("mysql").Unmarshal(&configs)
	if err != nil {
		Log().WithError(err).Error("gorm_init_error")
	}

	for n, c := range configs {
		log := Log().WithField("name", n).WithField("config", c)
		if err := gormManager.Init(n, c); err != nil {
			log.WithError(err).Error("gorm_init_error")
			panic("gorm_init_error name: " + n)
		}
		log.Info("gorm_init_success")
	}
}

func (manager *GORMClientManager) Init(name string, config *ORMConfig) error {
	//生成配置
	c := mysql.NewConfig()
	c.Loc = time.Local
	c.User = config.Username
	c.Passwd = config.Password
	c.Net = "tcp"
	c.Addr = fmt.Sprintf("%s:%d", config.Host, config.Port)
	c.DBName = config.Database
	c.Collation = "utf8mb4_general_ci"
	c.ParseTime = true

	//获取连接driver
	conn, err := mysql.NewConnector(c)
	if err != nil {
		return err
	}

	//获取连接
	client := sql.OpenDB(conn)

	//设置连接池
	client.SetMaxOpenConns(config.MaxConn)
	client.SetMaxIdleConns(config.MaxIdleConn)

	//测试连接
	if err := client.Ping(); err != nil {
		return err
	}

	//创建gorm
	dbEngine, err := gorm.Open(gmysql.New(gmysql.Config{
		Conn: client,
	}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		return err
	}

	//保存到manager
	manager.Set(name, dbEngine)

	return nil
}

// ----------------------------------------
//  Mysql 配置项
// ----------------------------------------

type ORMConfig struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Database    string `json:"database"`
	Charset     string `json:"charset"`
	MaxConn     int    `json:"max_conn" mapstructure:"max_conn"`
	MaxIdleConn int    `json:"max_idle_conn" mapstructure:"max_idle_conn"`
}
