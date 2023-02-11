package dal

import (
	"action/cmd/dal/db"
	"action/cmd/dal/redis"
)

func Init() {
	db.Init()
	redis.Init()
}
