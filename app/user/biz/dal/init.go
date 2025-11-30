package dal

import (
	"github.com/fuhanyang/CYGlowFlow/app/user/biz/dal/mysql"
	"github.com/fuhanyang/CYGlowFlow/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
