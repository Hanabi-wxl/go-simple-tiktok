package consts

const (
	ETCDAddr        = "127.0.0.1:2379"
	CoreServiceName = "rpcCoreService"
	CoreServiceAddr = "127.0.0.1:8083"
	MysqlDSN        = "root:@tcp(localhost:3306)/simple_tiktok?charset=utf8&parseTime=True&loc=Local"
	VideoLimit      = 2
	PassWordCost    = 12
)

var DefaultCode int32 = 0
var DefaultMsg = "success"
