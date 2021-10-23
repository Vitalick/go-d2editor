package utils

import (
	"strings"
)

//TitleToJSONTitle transform titles to json title format
func TitleToJSONTitle(s string) string {
	return strings.Replace(strings.ToLower(s), " ", "_", -1)
}
