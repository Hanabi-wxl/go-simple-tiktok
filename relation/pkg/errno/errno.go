package errno

import "github.com/pkg/errors"

var (
	DbSelectErr          = errors.New("查询错误")
	DbUpdateErr          = errors.New("更新错误")
	DbDeleteErr          = errors.New("删除错误")
	DbInsertErr          = errors.New("新增错误")
	UserAlreadyExitErr   = errors.New("用户已存在")
	UserNotExitErr       = errors.New("用户不存在")
	PasswordIncorrectErr = errors.New("用户名或密码错误")

	FollowErr    = errors.New("关注失败")
	DelFollowErr = errors.New("取消关注失败")
	ActionErr    = errors.New("操作异常")
)
