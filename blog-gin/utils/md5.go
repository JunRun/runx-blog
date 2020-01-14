/**
 *
 * @Description: 加密工具类
 * @Version: 1.0.0
 * @Date: 2020-01-10 14:20
 */
package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func MD5Encryption(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func MD5Check(data, encrypted string) bool {
	return strings.EqualFold(MD5Encryption(data), encrypted)
}
