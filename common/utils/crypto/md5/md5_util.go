package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(str string) string {
	res := md5.Sum([]byte(str))
	return hex.EncodeToString(res[:])
}
