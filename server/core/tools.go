package core

import "strings"

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// base62Encode 将数字转换为Base62格式
func base62Encode(num int64) string {
	if num == 0 {
		return string(base62Chars[0])
	}

	var encoded strings.Builder
	for num > 0 {
		remainder := num % 62
		encoded.WriteString(string(base62Chars[remainder]))
		num /= 62
	}
	return reverse(encoded.String())
}

// base62Decode 将Base62格式的字符串解码为数字
func base62Decode(encoded string) int64 {
	decoded := int64(0)
	for _, char := range encoded {
		decoded = decoded*62 + int64(strings.Index(base62Chars, string(char)))
	}
	return decoded
}

// reverse 反转字符串
func reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
