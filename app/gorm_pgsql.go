package app

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitPgSqlGorm() {
	configs := make(map[string]*PgSqlORMConfig)

	err := Config("postgresql").Unmarshal(&configs)
	if err != nil {
		Log().WithError(err).Error("gorm_init_error")
	}

	for n, c := range configs {
		log := Log().WithField("name", n).WithField("config", c)
		if err := gormManager.InitPg(n, c); err != nil {
			log.WithError(err).Error("gorm_init_error")
			panic("gorm_init_error name: " + n)
		}
		log.Info("gorm_init_success")
	}
}

func (manager *GORMClientManager) InitPg(name string, config *PgSqlORMConfig) error {
	//生成配置
	dsn := "host=" + config.Host + " user=" + config.Username + " password=" + config.Password + " dbname=" + config.Database + " port=" + fmt.Sprintf("%d", config.Port)

	dbEngine, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,  // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true, // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})

	if err != nil {
		return err
	} else {
		client, _ := dbEngine.DB()
		client.SetMaxOpenConns(config.MaxConn)
		client.SetMaxIdleConns(config.MaxIdleConn)
		//测试连接
		if err := client.Ping(); err != nil {
			return err
		}
	}

	//保存到manager
	manager.Set(name, dbEngine)

	return nil
}

// ----------------------------------------
//  Mysql 配置项
// ----------------------------------------

type PgSqlORMConfig struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Database    string `json:"database"`
	Charset     string `json:"charset"`
	MaxConn     int    `json:"max_conn" mapstructure:"max_conn"`
	MaxIdleConn int    `json:"max_idle_conn" mapstructure:"max_idle_conn"`
}
