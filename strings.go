package tools

import (
	"crypto/md5"
	"fmt"
	"io"
)

//	converts a string to an MD5 hash
func StrToMD5(str string) string {
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
