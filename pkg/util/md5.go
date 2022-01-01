package util

import (
	"crypto/md5"
	"encoding/hex"
)

//功能类似于把正常的字符串编译成十六进制码
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
