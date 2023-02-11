package utils

import (
	"gateway/pkg/consts"
	"github.com/importcjj/sensitive"
	"log"
)

var Filter *sensitive.Filter

func InitSensitiveFilter() {
	Filter = sensitive.New()
	if err := Filter.LoadWordDict(consts.SensitiveDictPath); err != nil {
		log.Println(err.Error())
	}
}
