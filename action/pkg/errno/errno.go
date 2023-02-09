package errno

import "github.com/pkg/errors"

var (
	DbSelectErr  = errors.New("查询错误")
	DbUpdateErr  = errors.New("更新错误")
	DbInsertErr  = errors.New("新增错误")
	VideoNotExit = errors.New("视频不存在")
)
