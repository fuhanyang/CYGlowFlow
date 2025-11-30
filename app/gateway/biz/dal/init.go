package dal

import (
	"github.com/fuhanyang/CYGlowFlow/app/gateway/biz/dal/mysql"
	"github.com/fuhanyang/CYGlowFlow/app/gateway/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
