package errno

import "github.com/pkg/errors"

var (
	DbSelectErr          = errors.New("查询错误")
	DbInsertErr          = errors.New("新增错误")
	UserAlreadyExistErr  = errors.New("用户已存在")
	UserNotExistErr      = errors.New("用户不存在")
	PasswordIncorrectErr = errors.New("用户名或密码错误")
)
