package models

import (
	"crypto/md5"
	"fmt"
	"io"
	"time"
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
