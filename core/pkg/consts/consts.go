package consts

const (
	//ETCDAddr         = "192.168.64.3:2379"
	//CoreServiceAddr  = "0.0.0.0:8083"
	//MysqlDSN         = "root:@tcp(localhost:3306)/simple_tiktok?charset=utf8&parseTime=True&loc=Local"
	//AvatarFileUrl    = "http://tiktok.sinre.top/static/avatar"
	//BackgroundImgUrl = "http://tiktok.sinre.top/static/avatar/back.jpg"

	ETCDAddr         = "127.0.0.1:2379"
	CoreServiceName  = "rpcCoreService"
	CoreServiceAddr  = "127.0.0.1:8083"
	MysqlDSN         = "root:@tcp(localhost:3306)/simple_tiktok?charset=utf8&parseTime=True&loc=Local"
	VideoLimit       = 5
	PassWordCost     = 12
	AvatarFileUrl    = "http://192.168.1.8:8000/static/avatar/"
	BackgroundImgUrl = "http://192.168.1.8:8000/static/avatar/back.jpg"
)

var DefaultCode int32 = 0
var DefaultMsg = "success"
