package basic

import "strings"

// IsEmpty 判断字符串是否为空字符串。
func IsEmpty(s string) bool {
	return s == ""
}

// IsBlank 判断是否为空白字符串。
// - 如果为多个空白字符，也算是空白字符串哦！
func IsBlank(s string) bool {
	return len(strings.TrimSpace(s)) < 1
}
