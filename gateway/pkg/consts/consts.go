package consts

import (
	"gateway/pkg/result"
)

const (
	//ETCDAddr            = "192.168.64.3:2379"
	//GateWayAddr         = "0.0.0.0:8000"
	//VideoFileUrl         = "http://tiktok.sinre.top/static/video/"
	//FfmpegPath          = "/bin/ffmpeg"

	ETCDAddr     = "127.0.0.1:2379"
	GateWayAddr  = "192.168.1.9:8000"
	VideoFileUrl = "http://192.168.1.9:8000/static/video/"
	FfmpegPath   = "/bin/ffmpeg.exe"

	CoreClientName      = "coreService.client"
	ActionClientName    = "actionService.client"
	RelationClientName  = "relationService.client"
	CoreServiceName     = "rpcCoreService"
	ActionServiceName   = "rpcActionService"
	RelationServiceName = "rpcRelationService"
	GateWayServiceName  = "httpService"
	AuthorizationKey    = "token"
	StaticFilePath      = "/static/video/"
	UploadFilePath      = "./static/video/"
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
	NoContentErrCode     = 4001
)

var (
	ParamErr         = result.NewClientError(ParamErrCode, "参数异常")
	PostFormVideoErr = result.NewClientError(PostFormVideoErrCode, "获取文件异常")
	VideoTypeErrErr  = result.NewClientError(VideoTypeErrCode, "视频格式不支持")
	SaveFileTempErr  = result.NewClientError(SaveFileTempErrCode, "暂存文件失败")
	VideoCaptureErr  = result.NewClientError(VideoCaptureErrCode, "视频截图失败")
	FileToLargeErr   = result.NewClientError(FileToLargeErrCode, "视频最大支持30MB")
	FileNotFoundErr  = result.NewClientError(FileToLargeErrCode, "未找到视频文件")
	NoContentErr     = result.NewClientError(NoContentErrCode, "发送内容为空")
)
