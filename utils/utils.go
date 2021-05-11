package utils

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