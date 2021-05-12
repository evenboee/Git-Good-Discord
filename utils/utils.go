package utils

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
