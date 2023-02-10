package pack

import (
	"relation/cmd/service"
)

func BuildFollowResp(resp *service.DouyinRelationActionResponse) {
	var (
		code int32 = 0
		msg        = "success"
	)
	resp.StatusCode = &code
	resp.StatusMsg = &msg
}

func BuildFollowListResp(resp *service.DouyinRelationFollowListResponse, userList []*service.User) {

	var (
		code int32 = 0
		msg        = "success"
	)
	resp.StatusCode = &code
	resp.StatusMsg = &msg
	resp.UserList = userList
}
