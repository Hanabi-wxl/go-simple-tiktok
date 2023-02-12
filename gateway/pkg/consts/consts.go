package consts

import (
	"gateway/pkg/result"
)

const (
	ETCDAddr            = "127.0.0.1:2379"
	GateWayAddr         = "192.168.1.8:8000" // "127.0.0.1:8000"
	CoreClientName      = "coreService.client"
	ActionClientName    = "actionService.client"
	RelationClientName  = "relationService.client"
	CoreServiceName     = "rpcCoreService"
	ActionServiceName   = "rpcActionService"
	RelationServiceName = "rpcRelationService"
	GateWayServiceName  = "httpService"
	AuthorizationKey    = "token"
	FfmpegPath          = "/bin/ffmpeg.exe"
	StaticFilePath      = "/static/"
	ServerIP            = "192.168.1.8" // "127.0.0.1"
	ServerPort          = "8000"
	SensitiveDictPath   = "./pkg/doc/dict.txt"
)

var (
	MaxFileSize  int64 = 1024 * 1024 * 30 // 30MB
	DefaultTime        = "1676112808789"
	VideoTypeMap       = map[string]struct{}{
		".mp4":  {},
		".avi":  {},
		".wmv":  {},
		".flv":  {},
		".mpeg": {},
		".mov":  {},
	}
)

/*
请求、操作：10xx
文件：30xx
权限：40xx
*/
const (
	ParamErrCode         = 1001
	PostFormVideoErrCode = 3001
	VideoTypeErrCode     = 3002
	SaveFileTempErrCode  = 3003
	VideoCaptureErrCode  = 3004
	FileToLargeErrCode   = 3005
	NoTokenErrCode       = 4000
	AuthorizationErrCode = 4001
)

var (
	ParamErr           = result.NewClientError(ParamErrCode, "参数异常")
	PostFormVideoErr   = result.NewClientError(PostFormVideoErrCode, "获取文件异常")
	VideoTypeErrErr    = result.NewClientError(VideoTypeErrCode, "视频格式不支持")
	SaveFileTempErr    = result.NewClientError(SaveFileTempErrCode, "暂存文件失败")
	VideoCaptureErr    = result.NewClientError(VideoCaptureErrCode, "视频截图失败")
	FileToLargeErr     = result.NewClientError(FileToLargeErrCode, "视频最大支持30MB")
	FileNotFoundErrErr = result.NewClientError(FileToLargeErrCode, "未找到视频文件")
)
