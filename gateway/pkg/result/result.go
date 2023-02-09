package result

import (
	"encoding/json"
	"github.com/pkg/errors"
)

// ServiceError
// @Description: 服务端异常类
type ServiceError struct {
	Id     string `json:"id"`
	Code   int32  `json:"code"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}

// ClientError
// @Description: 客户端异常类
type ClientError struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

// NewError 自定义异常 弃用
func NewError(code int, msg string) error {
	jsonBytes, _ := json.Marshal(map[string]interface{}{
		"status_code": code,
		"status_msg":  msg,
	})
	return errors.New(string(jsonBytes))
}

// NewClientError 自定义客户端异常
func NewClientError(code int, msg string) ClientError {
	return ClientError{
		StatusCode: code,
		StatusMsg:  msg,
	}
}

// ConvertServiceErr
// @Description: 用于将服务端返回的异常信息字符串转换为服务端异常对象
// @auth sinre 2023-02-09 21:06:20
// @param errs 异常信息字符串
// @return ServiceError 服务端异常对象
func ConvertServiceErr(errs string) ServiceError {
	var l, r int
	for i := range errs {
		if errs[i] == '{' {
			l = i
			break
		}
	}
	for i := l + 1; i < len(errs); i++ {
		if errs[i] == '}' {
			r = i
			break
		}
	}

	ss := errs[l : r+1]
	var se ServiceError
	_ = json.Unmarshal([]byte(ss), &se)
	return se
}
