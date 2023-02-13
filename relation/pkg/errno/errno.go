package errno

import "github.com/pkg/errors"

var (
	DbSelectErr     = errors.New("查询错误")
	DbInsertErr     = errors.New("新增错误")
	UserNotExistErr = errors.New("用户不存在")

	FollowErr    = errors.New("关注失败")
	DelFollowErr = errors.New("取消关注失败")
)
