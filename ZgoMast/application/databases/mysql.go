package databases

import (
	"ZgoMast/application/Config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
	"time"
)

var Orm *gorm.DB

func GetOrm(cfg *Config.DatabaseConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.Charset)
	var err error
	Orm, err = gorm.Open(cfg.Driver, dsn)
	if err != nil {
		zap.S().Errorf("database connection fail: %v", err.Error())
		panic(err.Error())
	}
	// 最大链接数
	Orm.DB().SetMaxOpenConns(cfg.MaximumConn)
	// 最大空闲连接数
	Orm.DB().SetMaxIdleConns(cfg.MaximumFreeConn)
	Orm.DB().SetConnMaxLifetime(time.Duration(cfg.TimeOut))
	if Orm.Error != nil {
		zap.S().Errorf("database error: %v", Orm.Error)
		panic(Orm.Error)
	}
	zap.S().Infof("mysql连接成功: %v", dsn)
	return Orm
}
