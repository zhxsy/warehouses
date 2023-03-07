package helper

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strings"
)

func GetCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}

// 文件或是文件夹是否存在
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// 是否是文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 是否是文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// 保证slice中的元素唯一
func UniqStrSlice(s []string) []string {
	tmpMap := make(map[string]string)
	tmpSlice := make([]string, 0)
	for _, val := range s {
		if _, ok := tmpMap[val]; ok {
			continue
		} else {
			tmpMap[val] = ""
			tmpSlice = append(tmpSlice, val)
		}
	}
	return tmpSlice
}

// 获取变量的类型
func GetType(v interface{}) string {
	return reflect.TypeOf(v).String()
}

func TypeOf(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

//获取字符串在列表中位置
func InStringSliceIdx(item string, s []string) int {
	for i, v := range s {
		if v == item {
			return i
		}
	}
	return -1
}

func GetRemoteInfo(ctx *gin.Context) (agent, remote string) {
	if remote = ctx.Request.Header.Get("X-Forwarded-For"); remote == "" {
		if remote = ctx.Request.Header.Get("X-Real-IP"); remote == "" {
			remote, _, _ = net.SplitHostPort(ctx.Request.RemoteAddr)
			if remote == "::1" {
				remote = "127.0.0.1"
			}
		}
	}

	addr := strings.Split(remote, ",")
	if len(addr) >= 2 {
		remote = addr[0]
	}
	agent = strings.ToLower(ctx.Request.UserAgent())
	return
}

/**
Input 必需。规定要填充的字符串。
PadLength 必需。规定新字符串的长度。如果该值小于原始字符串的长度，则不进行任何操作。
PadString 必选。规定供填充使用的字符串。
PadType 必选。0=填充字符串的左侧。1=填充字符串的右侧。2=填充字符串的两侧。
StrPadPadType 0 = "left",1 = "right", 2 = both
*/
func StrPad(Input string, PadLength int, PadString string, PadType int) string {

	var leftPad, rightPad = 0, 0

	numPadChars := PadLength - len(Input)

	if numPadChars <= 0 {
		return Input
	}

	var buffer bytes.Buffer

	buffer.WriteString(Input)

	switch PadType {

	case 0:

		leftPad = numPadChars

		rightPad = 0

	case 1:

		leftPad = 0

		rightPad = numPadChars

	case 2:

		rightPad = numPadChars / 2

		leftPad = numPadChars - rightPad

	}

	var leftBuffer bytes.Buffer

	/* 左填充：循环添加字符*/

	for i := 0; i < leftPad; i++ {

		leftBuffer.WriteString(PadString)

		if leftBuffer.Len() > leftPad {

			leftBuffer.Truncate(leftPad)

			break
		}

	}

	/* 右填充：循环添加字符串*/
	for i := 0; i < rightPad; i++ {

		buffer.WriteString(PadString)

		if buffer.Len() > PadLength {

			buffer.Truncate(PadLength)

			break
		}
	}

	leftBuffer.WriteString(buffer.String())

	return leftBuffer.String()

}
