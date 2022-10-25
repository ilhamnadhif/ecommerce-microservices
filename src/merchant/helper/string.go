package helper

import "strings"

func GenerateSlug(name string) string {
	return strings.Join(strings.Split(strings.ToLower(name), " "), "-")
}
