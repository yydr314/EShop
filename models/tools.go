package models

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func UnixToTime(timeUnix int) string {
	t := time.Unix(int64(timeUnix), 0)
	return t.Format(time.RFC3339)
}

// 上傳文件
func UploadFile(ctx *gin.Context, picName string) (string, error) {
	//	上傳文件
	file, err := ctx.FormFile(picName)
	if err != nil {
		return "", err
	}

	extName := path.Ext(file.Filename)

	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}

	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("文件格式不合法")
	}
	day := strconv.Itoa(time.Now().Day())
	dir := "./static/upload/" + day

	err = os.MkdirAll(dir, 0666)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fileName := strconv.FormatInt(time.Now().Unix(), 10) + extName

	dst := path.Join(dir, fileName)
	ctx.SaveUploadedFile(file, dst)
	return dst, nil
}
