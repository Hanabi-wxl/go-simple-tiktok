package utils

import (
	"fmt"
	"gateway/pkg/consts"
	"github.com/google/uuid"
	"os"
	"path/filepath"
)

func GetFileUrl(fileName string) string {
	base := fmt.Sprintf("http://%s:%s/static/%s", consts.ServerIP, consts.ServerPort, fileName)
	return base
}

// NewFileName 根据userId+用户发布的视频数量连接成独一无二的文件名
func NewFileName() string {
	uu := uuid.New()
	return uu.String()
}

// Capture 将视频切一帧保存到本地
// isDebug用于控制是否打印出执行的ffmepg命令
func Capture(name string, isDebug bool) error {
	vcp := NewVideoCapture()
	if isDebug {
		vcp.Debug()
	}
	dir, _ := os.Getwd()
	vcp.InputPath = filepath.Join(dir+consts.StaticFilePath, name+defaultVideoSuffix)
	vcp.OutputPath = filepath.Join(dir+consts.StaticFilePath, name+defaultImageSuffix)
	vcp.FrameCount = 1
	queryString, err := vcp.GetQueryString()
	if err != nil {
		return err
	}
	return vcp.ExecCommand(queryString)
}
