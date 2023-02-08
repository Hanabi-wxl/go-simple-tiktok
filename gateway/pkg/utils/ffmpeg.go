package utils

//#include <stdlib.h>
//int startCmd(const char* cmd){
//	  return system(cmd);
//}
import "C"
import (
	"fmt"
	"gateway/pkg/consts"
	"github.com/pkg/errors"
	"log"
	"os"
	"unsafe"
)

type VideoCapture struct {
	InputPath  string
	OutputPath string
	StartTime  string
	KeepTime   string
	Filter     string
	FrameCount int64
	debug      bool
}

func NewVideoCapture() *VideoCapture {
	return &videoCapture
}

var videoCapture VideoCapture

// ffmpeg的参数
const (
	inputVideoPathOption = "-i"
	startTimeOption      = "-ss"
	keepTimeOption       = "-t"
	videoFilterOption    = "-vf"
	formatToImageOption  = "-f"
	autoReWriteOption    = "-y"
	framesOption         = "-frames:v"
)

var (
	defaultVideoSuffix = ".mp4"
	defaultImageSuffix = ".jpg"
)

func ChangeVideoDefaultSuffix(suffix string) {
	defaultVideoSuffix = suffix
}

func ChangeImageDefaultSuffix(suffix string) {
	defaultImageSuffix = suffix
}

func GetDefaultImageSuffix() string {
	return defaultImageSuffix
}

func paramJoin(s1, s2 string) string {
	return fmt.Sprintf(" %s %s ", s1, s2)
}

func (v *VideoCapture) Debug() {
	v.debug = true
}

func (v *VideoCapture) GetQueryString() (ret string, err error) {
	if v.InputPath == "" || v.OutputPath == "" {
		err = errors.New("输入输出路径未指定")
		return
	}
	path, _ := os.Getwd()
	ret = path + consts.FfmpegPath
	ret += paramJoin(inputVideoPathOption, v.InputPath)
	ret += paramJoin(formatToImageOption, "image2")
	if v.Filter != "" {
		ret += paramJoin(videoFilterOption, v.Filter)
	}
	if v.StartTime != "" {
		ret += paramJoin(startTimeOption, v.StartTime)
	}
	if v.KeepTime != "" {
		ret += paramJoin(keepTimeOption, v.KeepTime)
	}
	if v.FrameCount != 0 {
		ret += paramJoin(framesOption, fmt.Sprintf("%d", v.FrameCount))
	}
	ret += paramJoin(autoReWriteOption, v.OutputPath)
	return
}

func (v *VideoCapture) ExecCommand(cmd string) error {
	if v.debug {
		log.Println(cmd)
	}
	cCmd := C.CString(cmd)
	defer C.free(unsafe.Pointer(cCmd))
	status := C.startCmd(cCmd)
	if status != 0 {
		return errors.New("视频切截图失败")
	}
	return nil
}
