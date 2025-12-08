package dal

import (
	"github.com/fuhanyang/CYGlowFlow/app/workflow/biz/dal/mysql"
	"github.com/fuhanyang/CYGlowFlow/app/workflow/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
