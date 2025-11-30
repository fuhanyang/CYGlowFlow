package mysql

import (
	"github.com/fuhanyang/CYGlowFlow/app/gateway/biz/dal/mysql/query"
	"github.com/fuhanyang/CYGlowFlow/app/gateway/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
	Dao *query.Query
)

func Init() {
	DB, err = gorm.Open(mysql.Open(conf.GetConf().MySQL.DSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	// 初始化 gorm-gen 的 Query 对象
	query.SetDefault(DB)
	Dao = query.Q
}
