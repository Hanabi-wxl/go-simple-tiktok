package errno

import "github.com/pkg/errors"

var (
	DbSelectErr = errors.New("查询错误")
	DbUpdateErr = errors.New("更新错误")
	DbInsertErr = errors.New("新增错误")

	VideoNotExist      = errors.New("视频不存在")
	UserNotExistErr    = errors.New("用户不存在")
	CommentNotExistErr = errors.New("评论不存在")
	RepeatStar         = errors.New("重复点赞")
)
