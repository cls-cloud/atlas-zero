package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Md5(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func AuthMd5(ip, browserName, browserVersion string) string {
	str := fmt.Sprintf("%s.%s.%s", ip, browserName, browserVersion)
	return Md5(str)
}
