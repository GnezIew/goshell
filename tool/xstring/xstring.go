package xstring

import "strings"

// 是否存在subSlice中任意一个字符串前缀
func StringsHasPrefixs(s string, subSlice []string) bool {
	for _, v := range subSlice {
		if strings.HasPrefix(s, v) {
			return true
		}
	}
	return false
}

// 去除多个前缀
func TrimMorePrefix(s string, prefixSlice []string) string {
	for _, v := range prefixSlice {
		s = strings.TrimPrefix(s, v)
	}
	return s
}

func ToLowerFisrt(s string) string {
	first := s[0]
	First := strings.ToLower(string(first))
	return First + s[1:]
}
