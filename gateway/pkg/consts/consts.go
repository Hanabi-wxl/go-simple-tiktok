package consts

import "gateway/pkg/result"

const (
	ETCDAddr            = "127.0.0.1:2379"
	GateWayAddr         = "127.0.0.1:8000"
	CoreClientName      = "coreService.client"
	ActionClientName    = "actionService.client"
	RelationClientName  = "relationService.client"
	CoreServiceName     = "rpcCoreService"
	ActionServiceName   = "rpcActionService"
	RelationServiceName = "rpcActionService"
	GateWayServiceName  = "httpService"
)

const (
	SuccessCode  = 100
	ParamErrCode = 10001
)

var (
	ParamErr = result.NewError(ParamErrCode, "参数异常")
)

var ZERO *int32
