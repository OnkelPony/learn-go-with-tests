package wc

import "strings"

func ReplaceAndUpper(text, old, new string, count int) string {
	result := strings.Replace(text, old, new, count)
	result = strings.ToUpper(result)
	return result
}
