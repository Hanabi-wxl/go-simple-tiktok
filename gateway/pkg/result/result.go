package result

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type ServiceError struct {
	Id     string `json:"id"`
	Code   int32  `json:"code"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}

type ClientError struct {
	StatusCode int    `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

func NewError(code int, msg string) error {
	jsonBytes, _ := json.Marshal(map[string]interface{}{
		"status_code": code,
		"status_msg":  msg,
	})
	return errors.New(string(jsonBytes))
}

func NewClientError(code int, msg string) ClientError {
	return ClientError{
		StatusCode: code,
		StatusMsg:  msg,
	}
}

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
