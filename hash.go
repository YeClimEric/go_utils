package go_utils

import (
	"crypto/md5"
	"fmt"
)
/*
* md5加密
*/
func Md5(source string)string{
	buf := []byte(source)
	has := md5.Sum(buf)
	md5Str := fmt.Sprintf("%x", has)
	return md5Str
}