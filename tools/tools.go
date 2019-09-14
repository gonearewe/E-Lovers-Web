package tools

import (
	"crypto/md5"
	"fmt"
)

func Md5(source string) string {
	checksum := md5.Sum([]byte(source))
	return fmt.Sprintf("%x", checksum)
}
