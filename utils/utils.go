package utils

import "strings"

func TitleToJsonTitle(s string) string {
	return strings.Replace(strings.ToLower(s), " ", "_", -1)
}
