package Init

import (
	"ZgoMast/application/Config"
	"ZgoMast/application/Model"
	. "ZgoMast/application/databases"
)

func InitDB(cfg *Config.DatabaseConfig) {
	Orm := GetOrm(cfg)
	//禁用复数
	Orm.SingularTable(true)
	//数据迁移
	Orm.AutoMigrate(&Model.User{})
}
