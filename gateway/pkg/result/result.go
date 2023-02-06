package result

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type GlobalError struct {
	Id     string `json:"id"`
	Code   int32  `json:"code"`
	Detail string `json:"detail"`
	Status string `json:"status"`
}

type CommonResult struct {
	Code   int         `json:"code"`
	Detail string      `json:"detail"`
	Data   interface{} `json:"data"`
}

func Ok(detail string) *CommonResult {
	return &CommonResult{
		Code:   200,
		Detail: detail,
	}
}

func NewResult(code int, detail string) *CommonResult {
	return &CommonResult{
		Code:   code,
		Detail: detail,
	}
}

func NewError(code int, detail string) error {
	jsonBytes, _ := json.Marshal(map[string]interface{}{
		"Code":   code,
		"Detail": detail,
	})
	return errors.New(string(jsonBytes))
}

func Data(data interface{}) *CommonResult {
	return &CommonResult{
		Code: 200,
		Data: data,
	}
}
