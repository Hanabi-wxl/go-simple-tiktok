package consts

import "gateway/pkg/result"

const (
	ETCDAddr            = "127.0.0.1:2379"
	GateWayAddr         = "192.168.1.8:8000"
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
	ServerIP            = "192.168.1.8"
	ServerPort          = "8000"
)

var (
	VideoTypeMap = map[string]struct{}{
		".mp4":  {},
		".avi":  {},
		".wmv":  {},
		".flv":  {},
		".mpeg": {},
		".mov":  {},
	}
	PictureTypeMap = map[string]struct{}{
		".jpg": {},
		".bmp": {},
		".png": {},
		".svg": {},
	}
)

/*
  请求、操作：10xx
  文件：30xx
  权限：40xx
*/
const (
	SuccessCode          = 1000
	ParamErrCode         = 1001
	PostFormVideoErrCode = 3001
	VideoTypeErrCode     = 3002
	SaveFileTempErrCode  = 3003
	VideoCaptureErrCode  = 3004
	AuthorizationErrCode = 4001
)

var (
	ParamErr         = result.NewClientError(ParamErrCode, "参数异常")
	PostFormVideoErr = result.NewClientError(PostFormVideoErrCode, "获取文件异常")
	VideoTypeErrErr  = result.NewClientError(VideoTypeErrCode, "视频格式不支持")
	SaveFileTempErr  = result.NewClientError(SaveFileTempErrCode, "暂存文件失败")
	VideoCaptureErr  = result.NewClientError(VideoCaptureErrCode, "视频截图失败")
	AuthorizationErr = result.NewClientError(VideoCaptureErrCode, "用户认证已过期")
)
