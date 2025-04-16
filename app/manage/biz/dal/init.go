package dal

import (
	"umdp/app/manage/biz/dal/mysql"
	"umdp/app/manage/biz/dal/redisStore"
)

func Init() {
	mysql.Init()

	// init redis store
	redisStore.Init()
}
