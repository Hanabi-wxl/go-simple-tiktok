package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gateway/biz/service"
	"gateway/pkg/consts"
	"gateway/pkg/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// SendServiceErr 服务端异常响应
func SendServiceErr(c *gin.Context, err error) {
	serviceError := result.ConvertServiceErr(err.Error())
	resMap := map[string]interface{}{
		"status_code": http.StatusBadRequest,
		"status_msg":  serviceError.Detail,
	}
	c.JSON(http.StatusBadRequest, resMap)
}

// SendClientErr 客户端异常响应
func SendClientErr(c *gin.Context, err result.ClientError) {
	resMap := map[string]interface{}{
		"status_code": err.StatusCode,
		"status_msg":  err.StatusMsg,
	}
	c.JSON(http.StatusBadRequest, resMap)
}

// SendValidateErr 数据校验异常响应
func SendValidateErr(c *gin.Context, err error) {
	var reason string
	// 判断validate错误类型
	switch err.(type) {
	case service.DouyinUserLoginRequestValidationError:
		reason = err.(service.DouyinUserLoginRequestValidationError).Reason()
	case service.DouyinUserRegisterRequestValidationError:
		reason = err.(service.DouyinUserRegisterRequestValidationError).Reason()
	}
	resMap := map[string]interface{}{
		"status_code": consts.ParamErrCode,
		"status_msg":  reason,
	}
	c.JSON(http.StatusBadRequest, resMap)
}

// SendMap map响应 弃用
func SendMap(c *gin.Context, response interface{}, strs ...string) {
	var resMap map[string]interface{}
	if marshalContent, err := json.Marshal(response); err != nil {
		fmt.Println(err)
	} else {
		d := json.NewDecoder(bytes.NewReader(marshalContent))
		d.UseNumber() // 设置将float64转为一个number
		if err := d.Decode(&resMap); err != nil {
			fmt.Println(err)
		} else {
			for k, v := range resMap {
				resMap[k] = v
			}
		}
	}

	for i := 0; i < len(strs); i++ {
		resMap[strs[i]] = nil
	}
	c.JSON(http.StatusOK, response)
}
