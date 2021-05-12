package utils

import (
	"crypto/rand"
	"fmt"
)

import "strings"

// ConvertToUniqueIntSlice will remove duplicates from a given int slice
func ConvertToUniqueIntSlice(intSlice []int) []int {
	keys := make(map[int]bool, len(intSlice))
	list := make([]int, 0, len(intSlice))

	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}

func GenerateUUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

func HTTPS(URL string, notSecure ...bool) string {
	if strings.HasPrefix(strings.ToLower(URL), "http") {
		return URL
	}
	if len(notSecure) > 0 {
		if notSecure[0] {
			return "http://" + URL
		}
	}
	return "https://" + URL
}
