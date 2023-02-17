package dal

import (
	"core/cmd/dal/db"
	"core/cmd/dal/redis"
)

func Init() {
	db.Init()
	redis.Init()
}
