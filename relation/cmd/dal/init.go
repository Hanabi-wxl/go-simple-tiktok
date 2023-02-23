package dal

import (
	"relation/cmd/dal/db"
	"relation/cmd/dal/redis"
)

func Init() {
	db.Init()
	redis.Init()
}
