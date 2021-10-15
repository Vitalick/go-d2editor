package utils

import "strings"

//TitleToJSONTitle ...
func TitleToJSONTitle(s string) string {
	return strings.Replace(strings.ToLower(s), " ", "_", -1)
}
